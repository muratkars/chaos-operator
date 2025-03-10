// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationParams) DeepCopyInto(out *ApplicationParams) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationParams.
func (in *ApplicationParams) DeepCopy() *ApplicationParams {
	if in == nil {
		return nil
	}
	out := new(ApplicationParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosEngine) DeepCopyInto(out *ChaosEngine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosEngine.
func (in *ChaosEngine) DeepCopy() *ChaosEngine {
	if in == nil {
		return nil
	}
	out := new(ChaosEngine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosEngine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosEngineList) DeepCopyInto(out *ChaosEngineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChaosEngine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosEngineList.
func (in *ChaosEngineList) DeepCopy() *ChaosEngineList {
	if in == nil {
		return nil
	}
	out := new(ChaosEngineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosEngineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosEngineSpec) DeepCopyInto(out *ChaosEngineSpec) {
	*out = *in
	out.Appinfo = in.Appinfo
	if in.Experiments != nil {
		in, out := &in.Experiments, &out.Experiments
		*out = make([]ExperimentList, len(*in))
		copy(*out, *in)
	}
	out.Schedule = in.Schedule
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosEngineSpec.
func (in *ChaosEngineSpec) DeepCopy() *ChaosEngineSpec {
	if in == nil {
		return nil
	}
	out := new(ChaosEngineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosEngineStatus) DeepCopyInto(out *ChaosEngineStatus) {
	*out = *in
	if in.Experiments != nil {
		in, out := &in.Experiments, &out.Experiments
		*out = make([]ExperimentStatuses, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosEngineStatus.
func (in *ChaosEngineStatus) DeepCopy() *ChaosEngineStatus {
	if in == nil {
		return nil
	}
	out := new(ChaosEngineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosExperiment) DeepCopyInto(out *ChaosExperiment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosExperiment.
func (in *ChaosExperiment) DeepCopy() *ChaosExperiment {
	if in == nil {
		return nil
	}
	out := new(ChaosExperiment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosExperiment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosExperimentList) DeepCopyInto(out *ChaosExperimentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChaosExperiment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosExperimentList.
func (in *ChaosExperimentList) DeepCopy() *ChaosExperimentList {
	if in == nil {
		return nil
	}
	out := new(ChaosExperimentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosExperimentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosExperimentSpec) DeepCopyInto(out *ChaosExperimentSpec) {
	*out = *in
	in.Definition.DeepCopyInto(&out.Definition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosExperimentSpec.
func (in *ChaosExperimentSpec) DeepCopy() *ChaosExperimentSpec {
	if in == nil {
		return nil
	}
	out := new(ChaosExperimentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosExperimentStatus) DeepCopyInto(out *ChaosExperimentStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosExperimentStatus.
func (in *ChaosExperimentStatus) DeepCopy() *ChaosExperimentStatus {
	if in == nil {
		return nil
	}
	out := new(ChaosExperimentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosResult) DeepCopyInto(out *ChaosResult) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosResult.
func (in *ChaosResult) DeepCopy() *ChaosResult {
	if in == nil {
		return nil
	}
	out := new(ChaosResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosResult) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosResultList) DeepCopyInto(out *ChaosResultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChaosResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosResultList.
func (in *ChaosResultList) DeepCopy() *ChaosResultList {
	if in == nil {
		return nil
	}
	out := new(ChaosResultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosResultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosResultSpec) DeepCopyInto(out *ChaosResultSpec) {
	*out = *in
	out.ExperimentStatus = in.ExperimentStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosResultSpec.
func (in *ChaosResultSpec) DeepCopy() *ChaosResultSpec {
	if in == nil {
		return nil
	}
	out := new(ChaosResultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosResultStatus) DeepCopyInto(out *ChaosResultStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosResultStatus.
func (in *ChaosResultStatus) DeepCopy() *ChaosResultStatus {
	if in == nil {
		return nil
	}
	out := new(ChaosResultStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosSchedule) DeepCopyInto(out *ChaosSchedule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosSchedule.
func (in *ChaosSchedule) DeepCopy() *ChaosSchedule {
	if in == nil {
		return nil
	}
	out := new(ChaosSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ENVPair) DeepCopyInto(out *ENVPair) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ENVPair.
func (in *ENVPair) DeepCopy() *ENVPair {
	if in == nil {
		return nil
	}
	out := new(ENVPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentAttributes) DeepCopyInto(out *ExperimentAttributes) {
	*out = *in
	out.Components = in.Components
	out.Schedule = in.Schedule
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentAttributes.
func (in *ExperimentAttributes) DeepCopy() *ExperimentAttributes {
	if in == nil {
		return nil
	}
	out := new(ExperimentAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentDef) DeepCopyInto(out *ExperimentDef) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ENVList != nil {
		in, out := &in.ENVList, &out.ENVList
		*out = make([]ENVPair, len(*in))
		copy(*out, *in)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentDef.
func (in *ExperimentDef) DeepCopy() *ExperimentDef {
	if in == nil {
		return nil
	}
	out := new(ExperimentDef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentList) DeepCopyInto(out *ExperimentList) {
	*out = *in
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentList.
func (in *ExperimentList) DeepCopy() *ExperimentList {
	if in == nil {
		return nil
	}
	out := new(ExperimentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentSchedule) DeepCopyInto(out *ExperimentSchedule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentSchedule.
func (in *ExperimentSchedule) DeepCopy() *ExperimentSchedule {
	if in == nil {
		return nil
	}
	out := new(ExperimentSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentStatuses) DeepCopyInto(out *ExperimentStatuses) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentStatuses.
func (in *ExperimentStatuses) DeepCopy() *ExperimentStatuses {
	if in == nil {
		return nil
	}
	out := new(ExperimentStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectUnderTest) DeepCopyInto(out *ObjectUnderTest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectUnderTest.
func (in *ObjectUnderTest) DeepCopy() *ObjectUnderTest {
	if in == nil {
		return nil
	}
	out := new(ObjectUnderTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestStatus) DeepCopyInto(out *TestStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestStatus.
func (in *TestStatus) DeepCopy() *TestStatus {
	if in == nil {
		return nil
	}
	out := new(TestStatus)
	in.DeepCopyInto(out)
	return out
}
