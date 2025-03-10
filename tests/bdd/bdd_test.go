package bdd

import (
	"fmt"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	scheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	chaosClient "github.com/litmuschaos/chaos-operator/pkg/client/clientset/versioned/typed/litmuschaos/v1alpha1"
)

var kubeconfig = "/home/circleci/.kube/config"
var config, _ = clientcmd.BuildConfigFromFlags("", kubeconfig)
var client, _ = kubernetes.NewForConfig(config)
var clientSet, _ = chaosClient.NewForConfig(config)

func TestChaos(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "BDD test")
}

var _ = BeforeSuite(func() {

	err := v1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		fmt.Println(err)
	}

	//Creating crds
	By("creating chaosengine crd")
	err = exec.Command("kubectl", "create", "-f", "../../deploy/chaos_crds.yaml").Run()
	if err != nil {
		fmt.Println(err)
	}

	//Creating rbacs
	err = exec.Command("kubectl", "create", "-f", "../../deploy/rbac.yaml").Run()
	if err != nil {
		fmt.Println(err)
	}

	//Creating Chaos-Operator
	By("creating operator")
	err = exec.Command("kubectl", "create", "-f", "../../deploy/operator.yaml").Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("chaos-operator created successfully")

	//Wait for the creation of chaos-operator
	time.Sleep(30 * time.Second)

	//Check for the status of the chaos-operator
	operator, _ := client.CoreV1().Pods("litmus").List(metav1.ListOptions{LabelSelector: "name=chaos-operator"})
	for _, v := range operator.Items {

		Expect(string(v.Status.Phase)).To(Equal("Running"))
		break
	}
})

//BDD Tests to check secondary resources
var _ = Describe("BDD on chaos-operator", func() {

	// BDD TEST CASE 1
	Context("Check for the custom resources", func() {

		It("Should check for creation of runner pod", func() {

			//creating nginx deployment
			deployment := &appv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nginx",
					Namespace: "litmus",
					Labels: map[string]string{
						"app": "nginx",
					},
					Annotations: map[string]string{
						"litmuschaos.io/chaos": "true",
					},
				},
				Spec: appv1.DeploymentSpec{
					Replicas: func(i int32) *int32 { return &i }(3),
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app": "nginx",
						},
					},
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								"app": "nginx",
							},
						},
						Spec: v1.PodSpec{
							ServiceAccountName: "litmus",
							Containers: []v1.Container{
								{
									Name:  "nginx",
									Image: "nginx:latest",
									Ports: []v1.ContainerPort{
										{

											ContainerPort: 80,
										},
									},
								},
							},
						},
					},
				},
			}

			_, err := client.AppsV1().Deployments("litmus").Create(deployment)
			if err != nil {
				fmt.Println("Deployment is not created and error is ", err)
			}

			//creating chaos-experiment for pod-delete
			By("Creating ChaosExperiments")
			ChaosExperiment := &v1alpha1.ChaosExperiment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod-delete",
					Namespace: "litmus",
					Labels: map[string]string{
						"litmuschaos.io/name": "kubernetes",
					},
				},
				Spec: v1alpha1.ChaosExperimentSpec{
					Definition: v1alpha1.ExperimentDef{

						Args:    []string{"-c", "ansible-playbook ./experiments/chaos/pod_delete/test.yml -i /etc/ansible/hosts -vv; exit 0"},
						Command: []string{"/bin/bash"},

						ENVList: []v1alpha1.ENVPair{
							{
								Name:  "ANSIBLE_STDOUT_CALLBACK",
								Value: "default",
							},
							{
								Name:  "TOTAL_CHAOS_DURATION",
								Value: "15",
							},
							{
								Name:  "CHAOS_INTERVAL",
								Value: "5",
							},
							{
								Name:  "LIB",
								Value: "",
							},
						},
						Image: "",
						Labels: map[string]string{
							"name": "pod-delete",
						},
					},
				},
			}

			_, err = clientSet.ChaosExperiments("litmus").Create(ChaosExperiment)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("ChaosExperiment created successfully...")

			//Creating chaosEngine
			By("Creating ChaosEngine")
			chaosEngine := &v1alpha1.ChaosEngine{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "engine-nginx",
					Namespace: "litmus",
				},
				Spec: v1alpha1.ChaosEngineSpec{
					Appinfo: v1alpha1.ApplicationParams{
						Appns:    "litmus",
						Applabel: "app=nginx",
					},
					ChaosServiceAccount: "litmus",
					Experiments: []v1alpha1.ExperimentList{
						{
							Name: "pod-delete",
						},
					},
				},
			}

			_, err = clientSet.ChaosEngines("litmus").Create(chaosEngine)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Chaosengine created successfully...")

			//Wait till the creation of runner pod and monitor svc
			time.Sleep(100 * time.Second)

			//Fetching engine-nginx-runner pod
			runner, err := client.CoreV1().Pods("litmus").Get("engine-nginx-runner", metav1.GetOptions{})

			//Check for the Availabilty and status of the runner pod
			fmt.Println("name : ", runner.Name)
			Expect(err).To(BeNil())
			Expect(string(runner.Status.Phase)).To(Equal("Running"))
		})
	})

	// BDD TEST CASE 2
	Context("check for the custom resources", func() {

		It("Should check for creation of monitor service", func() {
			_, err := client.CoreV1().Services("litmus").Get("engine-nginx-monitor", metav1.GetOptions{})

			Expect(err).To(BeNil())

		})

	})
})

//Deleting all unused resources
var _ = AfterSuite(func() {

	By("Deleting all CRDs")
	crdDeletion := exec.Command("kubectl", "delete", "-f", "../../deploy/chaos_crds.yaml").Run()
	Expect(crdDeletion).To(BeNil())
})
