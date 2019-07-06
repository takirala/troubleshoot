// +build !ignore_autogenerated

// autogenerated by controller-gen object, do not modify manually

package v1beta1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Analyzer) DeepCopyInto(out *Analyzer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Analyzer.
func (in *Analyzer) DeepCopy() *Analyzer {
	if in == nil {
		return nil
	}
	out := new(Analyzer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Analyzer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerJob) DeepCopyInto(out *AnalyzerJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerJob.
func (in *AnalyzerJob) DeepCopy() *AnalyzerJob {
	if in == nil {
		return nil
	}
	out := new(AnalyzerJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyzerJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerJobList) DeepCopyInto(out *AnalyzerJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnalyzerJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerJobList.
func (in *AnalyzerJobList) DeepCopy() *AnalyzerJobList {
	if in == nil {
		return nil
	}
	out := new(AnalyzerJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyzerJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerJobSpec) DeepCopyInto(out *AnalyzerJobSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerJobSpec.
func (in *AnalyzerJobSpec) DeepCopy() *AnalyzerJobSpec {
	if in == nil {
		return nil
	}
	out := new(AnalyzerJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerJobStatus) DeepCopyInto(out *AnalyzerJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerJobStatus.
func (in *AnalyzerJobStatus) DeepCopy() *AnalyzerJobStatus {
	if in == nil {
		return nil
	}
	out := new(AnalyzerJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerList) DeepCopyInto(out *AnalyzerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Analyzer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerList.
func (in *AnalyzerList) DeepCopy() *AnalyzerList {
	if in == nil {
		return nil
	}
	out := new(AnalyzerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyzerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerSpec) DeepCopyInto(out *AnalyzerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerSpec.
func (in *AnalyzerSpec) DeepCopy() *AnalyzerSpec {
	if in == nil {
		return nil
	}
	out := new(AnalyzerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerStatus) DeepCopyInto(out *AnalyzerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerStatus.
func (in *AnalyzerStatus) DeepCopy() *AnalyzerStatus {
	if in == nil {
		return nil
	}
	out := new(AnalyzerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Collect) DeepCopyInto(out *Collect) {
	*out = *in
	if in.KubernetesAPIVersions != nil {
		in, out := &in.KubernetesAPIVersions, &out.KubernetesAPIVersions
		*out = new(KubernetesAPIVersionsOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.KubernetesClusterInfo != nil {
		in, out := &in.KubernetesClusterInfo, &out.KubernetesClusterInfo
		*out = new(KubernetesClusterInfoOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.KubernetesContainerCp != nil {
		in, out := &in.KubernetesContainerCp, &out.KubernetesContainerCp
		*out = new(KubernetesContainerCpOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.KubernetesLogs != nil {
		in, out := &in.KubernetesLogs, &out.KubernetesLogs
		*out = new(KubernetesLogsOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.KubernetesResourceList != nil {
		in, out := &in.KubernetesResourceList, &out.KubernetesResourceList
		*out = new(KubernetesResourceListOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(KubernetesVersionOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Collect.
func (in *Collect) DeepCopy() *Collect {
	if in == nil {
		return nil
	}
	out := new(Collect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Collector) DeepCopyInto(out *Collector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Collector.
func (in *Collector) DeepCopy() *Collector {
	if in == nil {
		return nil
	}
	out := new(Collector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Collector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorJob) DeepCopyInto(out *CollectorJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorJob.
func (in *CollectorJob) DeepCopy() *CollectorJob {
	if in == nil {
		return nil
	}
	out := new(CollectorJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CollectorJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorJobList) DeepCopyInto(out *CollectorJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CollectorJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorJobList.
func (in *CollectorJobList) DeepCopy() *CollectorJobList {
	if in == nil {
		return nil
	}
	out := new(CollectorJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CollectorJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorJobSpec) DeepCopyInto(out *CollectorJobSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorJobSpec.
func (in *CollectorJobSpec) DeepCopy() *CollectorJobSpec {
	if in == nil {
		return nil
	}
	out := new(CollectorJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorJobStatus) DeepCopyInto(out *CollectorJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorJobStatus.
func (in *CollectorJobStatus) DeepCopy() *CollectorJobStatus {
	if in == nil {
		return nil
	}
	out := new(CollectorJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorList) DeepCopyInto(out *CollectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Collector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorList.
func (in *CollectorList) DeepCopy() *CollectorList {
	if in == nil {
		return nil
	}
	out := new(CollectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CollectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorSpec) DeepCopyInto(out *CollectorSpec) {
	*out = *in
	if in.Collectors != nil {
		in, out := &in.Collectors, &out.Collectors
		*out = make([]*Collect, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Collect)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorSpec.
func (in *CollectorSpec) DeepCopy() *CollectorSpec {
	if in == nil {
		return nil
	}
	out := new(CollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorStatus) DeepCopyInto(out *CollectorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorStatus.
func (in *CollectorStatus) DeepCopy() *CollectorStatus {
	if in == nil {
		return nil
	}
	out := new(CollectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesAPIVersionsOptions) DeepCopyInto(out *KubernetesAPIVersionsOptions) {
	*out = *in
	in.SpecShared.DeepCopyInto(&out.SpecShared)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesAPIVersionsOptions.
func (in *KubernetesAPIVersionsOptions) DeepCopy() *KubernetesAPIVersionsOptions {
	if in == nil {
		return nil
	}
	out := new(KubernetesAPIVersionsOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesClusterInfoOptions) DeepCopyInto(out *KubernetesClusterInfoOptions) {
	*out = *in
	in.SpecShared.DeepCopyInto(&out.SpecShared)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesClusterInfoOptions.
func (in *KubernetesClusterInfoOptions) DeepCopy() *KubernetesClusterInfoOptions {
	if in == nil {
		return nil
	}
	out := new(KubernetesClusterInfoOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesContainerCpOptions) DeepCopyInto(out *KubernetesContainerCpOptions) {
	*out = *in
	in.SpecShared.DeepCopyInto(&out.SpecShared)
	if in.PodListOptions != nil {
		in, out := &in.PodListOptions, &out.PodListOptions
		*out = new(metav1.ListOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesContainerCpOptions.
func (in *KubernetesContainerCpOptions) DeepCopy() *KubernetesContainerCpOptions {
	if in == nil {
		return nil
	}
	out := new(KubernetesContainerCpOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesLogsOptions) DeepCopyInto(out *KubernetesLogsOptions) {
	*out = *in
	in.SpecShared.DeepCopyInto(&out.SpecShared)
	if in.PodLogOptions != nil {
		in, out := &in.PodLogOptions, &out.PodLogOptions
		*out = new(v1.PodLogOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.ListOptions != nil {
		in, out := &in.ListOptions, &out.ListOptions
		*out = new(metav1.ListOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesLogsOptions.
func (in *KubernetesLogsOptions) DeepCopy() *KubernetesLogsOptions {
	if in == nil {
		return nil
	}
	out := new(KubernetesLogsOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesResourceListOptions) DeepCopyInto(out *KubernetesResourceListOptions) {
	*out = *in
	in.SpecShared.DeepCopyInto(&out.SpecShared)
	if in.ListOptions != nil {
		in, out := &in.ListOptions, &out.ListOptions
		*out = new(metav1.ListOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesResourceListOptions.
func (in *KubernetesResourceListOptions) DeepCopy() *KubernetesResourceListOptions {
	if in == nil {
		return nil
	}
	out := new(KubernetesResourceListOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesVersionOptions) DeepCopyInto(out *KubernetesVersionOptions) {
	*out = *in
	in.SpecShared.DeepCopyInto(&out.SpecShared)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesVersionOptions.
func (in *KubernetesVersionOptions) DeepCopy() *KubernetesVersionOptions {
	if in == nil {
		return nil
	}
	out := new(KubernetesVersionOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Meta) DeepCopyInto(out *Meta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Meta.
func (in *Meta) DeepCopy() *Meta {
	if in == nil {
		return nil
	}
	out := new(Meta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Preflight) DeepCopyInto(out *Preflight) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Preflight.
func (in *Preflight) DeepCopy() *Preflight {
	if in == nil {
		return nil
	}
	out := new(Preflight)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Preflight) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightJob) DeepCopyInto(out *PreflightJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightJob.
func (in *PreflightJob) DeepCopy() *PreflightJob {
	if in == nil {
		return nil
	}
	out := new(PreflightJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreflightJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightJobList) DeepCopyInto(out *PreflightJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PreflightJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightJobList.
func (in *PreflightJobList) DeepCopy() *PreflightJobList {
	if in == nil {
		return nil
	}
	out := new(PreflightJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreflightJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightJobSpec) DeepCopyInto(out *PreflightJobSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightJobSpec.
func (in *PreflightJobSpec) DeepCopy() *PreflightJobSpec {
	if in == nil {
		return nil
	}
	out := new(PreflightJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightJobStatus) DeepCopyInto(out *PreflightJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightJobStatus.
func (in *PreflightJobStatus) DeepCopy() *PreflightJobStatus {
	if in == nil {
		return nil
	}
	out := new(PreflightJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightList) DeepCopyInto(out *PreflightList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Preflight, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightList.
func (in *PreflightList) DeepCopy() *PreflightList {
	if in == nil {
		return nil
	}
	out := new(PreflightList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreflightList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightSpec) DeepCopyInto(out *PreflightSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightSpec.
func (in *PreflightSpec) DeepCopy() *PreflightSpec {
	if in == nil {
		return nil
	}
	out := new(PreflightSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightStatus) DeepCopyInto(out *PreflightStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightStatus.
func (in *PreflightStatus) DeepCopy() *PreflightStatus {
	if in == nil {
		return nil
	}
	out := new(PreflightStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scrub) DeepCopyInto(out *Scrub) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scrub.
func (in *Scrub) DeepCopy() *Scrub {
	if in == nil {
		return nil
	}
	out := new(Scrub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecShared) DeepCopyInto(out *SpecShared) {
	*out = *in
	if in.Scrub != nil {
		in, out := &in.Scrub, &out.Scrub
		*out = new(Scrub)
		**out = **in
	}
	if in.Meta != nil {
		in, out := &in.Meta, &out.Meta
		*out = new(Meta)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecShared.
func (in *SpecShared) DeepCopy() *SpecShared {
	if in == nil {
		return nil
	}
	out := new(SpecShared)
	in.DeepCopyInto(out)
	return out
}
