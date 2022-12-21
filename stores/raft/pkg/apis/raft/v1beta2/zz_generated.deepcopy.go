//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta2

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSinkConfig) DeepCopyInto(out *FileSinkConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSinkConfig.
func (in *FileSinkConfig) DeepCopy() *FileSinkConfig {
	if in == nil {
		return nil
	}
	out := new(FileSinkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggerConfig) DeepCopyInto(out *LoggerConfig) {
	*out = *in
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = make(map[string]OutputConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggerConfig.
func (in *LoggerConfig) DeepCopy() *LoggerConfig {
	if in == nil {
		return nil
	}
	out := new(LoggerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfig) DeepCopyInto(out *LoggingConfig) {
	*out = *in
	if in.Loggers != nil {
		in, out := &in.Loggers, &out.Loggers
		*out = make(map[string]LoggerConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Sinks != nil {
		in, out := &in.Sinks, &out.Sinks
		*out = make(map[string]SinkConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfig.
func (in *LoggingConfig) DeepCopy() *LoggingConfig {
	if in == nil {
		return nil
	}
	out := new(LoggingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputConfig) DeepCopyInto(out *OutputConfig) {
	*out = *in
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(string)
		**out = **in
	}
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputConfig.
func (in *OutputConfig) DeepCopy() *OutputConfig {
	if in == nil {
		return nil
	}
	out := new(OutputConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftCluster) DeepCopyInto(out *RaftCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftCluster.
func (in *RaftCluster) DeepCopy() *RaftCluster {
	if in == nil {
		return nil
	}
	out := new(RaftCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftClusterConfig) DeepCopyInto(out *RaftClusterConfig) {
	*out = *in
	out.RTT = in.RTT
	in.Server.DeepCopyInto(&out.Server)
	in.Logging.DeepCopyInto(&out.Logging)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftClusterConfig.
func (in *RaftClusterConfig) DeepCopy() *RaftClusterConfig {
	if in == nil {
		return nil
	}
	out := new(RaftClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftClusterList) DeepCopyInto(out *RaftClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RaftCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftClusterList.
func (in *RaftClusterList) DeepCopy() *RaftClusterList {
	if in == nil {
		return nil
	}
	out := new(RaftClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftClusterPartitionStatus) DeepCopyInto(out *RaftClusterPartitionStatus) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftClusterPartitionStatus.
func (in *RaftClusterPartitionStatus) DeepCopy() *RaftClusterPartitionStatus {
	if in == nil {
		return nil
	}
	out := new(RaftClusterPartitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftClusterSpec) DeepCopyInto(out *RaftClusterSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeClaimTemplate != nil {
		in, out := &in.VolumeClaimTemplate, &out.VolumeClaimTemplate
		*out = new(v1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftClusterSpec.
func (in *RaftClusterSpec) DeepCopy() *RaftClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RaftClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftClusterStatus) DeepCopyInto(out *RaftClusterStatus) {
	*out = *in
	if in.PartitionStatuses != nil {
		in, out := &in.PartitionStatuses, &out.PartitionStatuses
		*out = make([]RaftClusterPartitionStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftClusterStatus.
func (in *RaftClusterStatus) DeepCopy() *RaftClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RaftClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftConfig) DeepCopyInto(out *RaftConfig) {
	*out = *in
	if in.ElectionRTT != nil {
		in, out := &in.ElectionRTT, &out.ElectionRTT
		*out = new(uint64)
		**out = **in
	}
	if in.HeartbeatRTT != nil {
		in, out := &in.HeartbeatRTT, &out.HeartbeatRTT
		*out = new(uint64)
		**out = **in
	}
	if in.SnapshotEntries != nil {
		in, out := &in.SnapshotEntries, &out.SnapshotEntries
		*out = new(int64)
		**out = **in
	}
	if in.CompactionOverhead != nil {
		in, out := &in.CompactionOverhead, &out.CompactionOverhead
		*out = new(int64)
		**out = **in
	}
	if in.MaxInMemLogSize != nil {
		in, out := &in.MaxInMemLogSize, &out.MaxInMemLogSize
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftConfig.
func (in *RaftConfig) DeepCopy() *RaftConfig {
	if in == nil {
		return nil
	}
	out := new(RaftConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftMember) DeepCopyInto(out *RaftMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftMember.
func (in *RaftMember) DeepCopy() *RaftMember {
	if in == nil {
		return nil
	}
	out := new(RaftMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftMemberConfig) DeepCopyInto(out *RaftMemberConfig) {
	*out = *in
	in.RaftConfig.DeepCopyInto(&out.RaftConfig)
	if in.Peers != nil {
		in, out := &in.Peers, &out.Peers
		*out = make([]RaftMemberReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftMemberConfig.
func (in *RaftMemberConfig) DeepCopy() *RaftMemberConfig {
	if in == nil {
		return nil
	}
	out := new(RaftMemberConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftMemberList) DeepCopyInto(out *RaftMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RaftMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftMemberList.
func (in *RaftMemberList) DeepCopy() *RaftMemberList {
	if in == nil {
		return nil
	}
	out := new(RaftMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftMemberReference) DeepCopyInto(out *RaftMemberReference) {
	*out = *in
	out.Pod = in.Pod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftMemberReference.
func (in *RaftMemberReference) DeepCopy() *RaftMemberReference {
	if in == nil {
		return nil
	}
	out := new(RaftMemberReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftMemberSpec) DeepCopyInto(out *RaftMemberSpec) {
	*out = *in
	out.Cluster = in.Cluster
	out.Pod = in.Pod
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftMemberSpec.
func (in *RaftMemberSpec) DeepCopy() *RaftMemberSpec {
	if in == nil {
		return nil
	}
	out := new(RaftMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftMemberStatus) DeepCopyInto(out *RaftMemberStatus) {
	*out = *in
	if in.PodRef != nil {
		in, out := &in.PodRef, &out.PodRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(int32)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(RaftMemberRole)
		**out = **in
	}
	if in.Leader != nil {
		in, out := &in.Leader, &out.Leader
		*out = new(MemberID)
		**out = **in
	}
	if in.Term != nil {
		in, out := &in.Term, &out.Term
		*out = new(uint64)
		**out = **in
	}
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = (*in).DeepCopy()
	}
	if in.LastSnapshotIndex != nil {
		in, out := &in.LastSnapshotIndex, &out.LastSnapshotIndex
		*out = new(uint64)
		**out = **in
	}
	if in.LastSnapshotTime != nil {
		in, out := &in.LastSnapshotTime, &out.LastSnapshotTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftMemberStatus.
func (in *RaftMemberStatus) DeepCopy() *RaftMemberStatus {
	if in == nil {
		return nil
	}
	out := new(RaftMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftPartition) DeepCopyInto(out *RaftPartition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftPartition.
func (in *RaftPartition) DeepCopy() *RaftPartition {
	if in == nil {
		return nil
	}
	out := new(RaftPartition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftPartition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftPartitionList) DeepCopyInto(out *RaftPartitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RaftPartition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftPartitionList.
func (in *RaftPartitionList) DeepCopy() *RaftPartitionList {
	if in == nil {
		return nil
	}
	out := new(RaftPartitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftPartitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftPartitionMemberStatus) DeepCopyInto(out *RaftPartitionMemberStatus) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftPartitionMemberStatus.
func (in *RaftPartitionMemberStatus) DeepCopy() *RaftPartitionMemberStatus {
	if in == nil {
		return nil
	}
	out := new(RaftPartitionMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftPartitionSpec) DeepCopyInto(out *RaftPartitionSpec) {
	*out = *in
	in.RaftConfig.DeepCopyInto(&out.RaftConfig)
	out.Cluster = in.Cluster
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftPartitionSpec.
func (in *RaftPartitionSpec) DeepCopy() *RaftPartitionSpec {
	if in == nil {
		return nil
	}
	out := new(RaftPartitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftPartitionStatus) DeepCopyInto(out *RaftPartitionStatus) {
	*out = *in
	if in.Term != nil {
		in, out := &in.Term, &out.Term
		*out = new(uint64)
		**out = **in
	}
	if in.Leader != nil {
		in, out := &in.Leader, &out.Leader
		*out = new(MemberID)
		**out = **in
	}
	if in.Followers != nil {
		in, out := &in.Followers, &out.Followers
		*out = make([]MemberID, len(*in))
		copy(*out, *in)
	}
	if in.MemberStatuses != nil {
		in, out := &in.MemberStatuses, &out.MemberStatuses
		*out = make([]RaftPartitionMemberStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftPartitionStatus.
func (in *RaftPartitionStatus) DeepCopy() *RaftPartitionStatus {
	if in == nil {
		return nil
	}
	out := new(RaftPartitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftServerConfig) DeepCopyInto(out *RaftServerConfig) {
	*out = *in
	if in.ReadBufferSize != nil {
		in, out := &in.ReadBufferSize, &out.ReadBufferSize
		*out = new(int)
		**out = **in
	}
	if in.WriteBufferSize != nil {
		in, out := &in.WriteBufferSize, &out.WriteBufferSize
		*out = new(int)
		**out = **in
	}
	if in.MaxRecvMsgSize != nil {
		in, out := &in.MaxRecvMsgSize, &out.MaxRecvMsgSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.MaxSendMsgSize != nil {
		in, out := &in.MaxSendMsgSize, &out.MaxSendMsgSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.NumStreamWorkers != nil {
		in, out := &in.NumStreamWorkers, &out.NumStreamWorkers
		*out = new(uint32)
		**out = **in
	}
	if in.MaxConcurrentStreams != nil {
		in, out := &in.MaxConcurrentStreams, &out.MaxConcurrentStreams
		*out = new(uint32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftServerConfig.
func (in *RaftServerConfig) DeepCopy() *RaftServerConfig {
	if in == nil {
		return nil
	}
	out := new(RaftServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftStore) DeepCopyInto(out *RaftStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftStore.
func (in *RaftStore) DeepCopy() *RaftStore {
	if in == nil {
		return nil
	}
	out := new(RaftStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftStoreList) DeepCopyInto(out *RaftStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RaftStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftStoreList.
func (in *RaftStoreList) DeepCopy() *RaftStoreList {
	if in == nil {
		return nil
	}
	out := new(RaftStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftStoreSpec) DeepCopyInto(out *RaftStoreSpec) {
	*out = *in
	in.RaftConfig.DeepCopyInto(&out.RaftConfig)
	out.Cluster = in.Cluster
	if in.ReplicationFactor != nil {
		in, out := &in.ReplicationFactor, &out.ReplicationFactor
		*out = new(uint32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftStoreSpec.
func (in *RaftStoreSpec) DeepCopy() *RaftStoreSpec {
	if in == nil {
		return nil
	}
	out := new(RaftStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftStoreStatus) DeepCopyInto(out *RaftStoreStatus) {
	*out = *in
	if in.ReplicationFactor != nil {
		in, out := &in.ReplicationFactor, &out.ReplicationFactor
		*out = new(uint32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftStoreStatus.
func (in *RaftStoreStatus) DeepCopy() *RaftStoreStatus {
	if in == nil {
		return nil
	}
	out := new(RaftStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkConfig) DeepCopyInto(out *SinkConfig) {
	*out = *in
	if in.Encoding != nil {
		in, out := &in.Encoding, &out.Encoding
		*out = new(string)
		**out = **in
	}
	if in.Stdout != nil {
		in, out := &in.Stdout, &out.Stdout
		*out = new(StdoutSinkConfig)
		**out = **in
	}
	if in.Stderr != nil {
		in, out := &in.Stderr, &out.Stderr
		*out = new(StderrSinkConfig)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(FileSinkConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkConfig.
func (in *SinkConfig) DeepCopy() *SinkConfig {
	if in == nil {
		return nil
	}
	out := new(SinkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StderrSinkConfig) DeepCopyInto(out *StderrSinkConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StderrSinkConfig.
func (in *StderrSinkConfig) DeepCopy() *StderrSinkConfig {
	if in == nil {
		return nil
	}
	out := new(StderrSinkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StdoutSinkConfig) DeepCopyInto(out *StdoutSinkConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StdoutSinkConfig.
func (in *StdoutSinkConfig) DeepCopy() *StdoutSinkConfig {
	if in == nil {
		return nil
	}
	out := new(StdoutSinkConfig)
	in.DeepCopyInto(out)
	return out
}