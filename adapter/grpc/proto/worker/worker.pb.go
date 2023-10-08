// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: worker.proto

// 包名

package worker

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Algorithm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type    int64  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Part    int64  `protobuf:"varint,3,opt,name=part,proto3" json:"part,omitempty"`
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Path    string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Schema  string `protobuf:"bytes,6,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *Algorithm) Reset() {
	*x = Algorithm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Algorithm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Algorithm) ProtoMessage() {}

func (x *Algorithm) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Algorithm.ProtoReflect.Descriptor instead.
func (*Algorithm) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

func (x *Algorithm) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Algorithm) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Algorithm) GetPart() int64 {
	if x != nil {
		return x.Part
	}
	return 0
}

func (x *Algorithm) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Algorithm) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Algorithm) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

type StartTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TaskId        string     `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	WorkerId      string     `protobuf:"bytes,3,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	Algorithm     *Algorithm `protobuf:"bytes,4,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Status        int64      `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	ExecutorState int64      `protobuf:"varint,6,opt,name=executor_state,json=executorState,proto3" json:"executor_state,omitempty"`
	Desc          string     `protobuf:"bytes,7,opt,name=desc,proto3" json:"desc,omitempty"`
	Input         string     `protobuf:"bytes,8,opt,name=input,proto3" json:"input,omitempty"`
	Output        string     `protobuf:"bytes,9,opt,name=output,proto3" json:"output,omitempty"`
	Ctime         int64      `protobuf:"varint,10,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Utime         int64      `protobuf:"varint,11,opt,name=utime,proto3" json:"utime,omitempty"`
}

func (x *StartTaskReq) Reset() {
	*x = StartTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTaskReq) ProtoMessage() {}

func (x *StartTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTaskReq.ProtoReflect.Descriptor instead.
func (*StartTaskReq) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{1}
}

func (x *StartTaskReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StartTaskReq) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *StartTaskReq) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *StartTaskReq) GetAlgorithm() *Algorithm {
	if x != nil {
		return x.Algorithm
	}
	return nil
}

func (x *StartTaskReq) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *StartTaskReq) GetExecutorState() int64 {
	if x != nil {
		return x.ExecutorState
	}
	return 0
}

func (x *StartTaskReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *StartTaskReq) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *StartTaskReq) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *StartTaskReq) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

func (x *StartTaskReq) GetUtime() int64 {
	if x != nil {
		return x.Utime
	}
	return 0
}

type StartTaskRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartTaskRsp) Reset() {
	*x = StartTaskRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTaskRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTaskRsp) ProtoMessage() {}

func (x *StartTaskRsp) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTaskRsp.ProtoReflect.Descriptor instead.
func (*StartTaskRsp) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{2}
}

type StopTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *StopTaskReq) Reset() {
	*x = StopTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTaskReq) ProtoMessage() {}

func (x *StopTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTaskReq.ProtoReflect.Descriptor instead.
func (*StopTaskReq) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{3}
}

func (x *StopTaskReq) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type StopTaskRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopTaskRsp) Reset() {
	*x = StopTaskRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopTaskRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTaskRsp) ProtoMessage() {}

func (x *StopTaskRsp) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTaskRsp.ProtoReflect.Descriptor instead.
func (*StopTaskRsp) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{4}
}

type GetTaskMetricReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTaskMetricReq) Reset() {
	*x = GetTaskMetricReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskMetricReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskMetricReq) ProtoMessage() {}

func (x *GetTaskMetricReq) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskMetricReq.ProtoReflect.Descriptor instead.
func (*GetTaskMetricReq) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{5}
}

type GetTaskMetricRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuStats    *CpuStats    `protobuf:"bytes,1,opt,name=cpu_stats,json=cpuStats,proto3" json:"cpu_stats,omitempty"`
	MemoryStats *MemoryStats `protobuf:"bytes,2,opt,name=memory_stats,json=memoryStats,proto3" json:"memory_stats,omitempty"`
	GpuStats    *GpuStats    `protobuf:"bytes,3,opt,name=gpu_stats,json=gpuStats,proto3" json:"gpu_stats,omitempty"`
}

func (x *GetTaskMetricRsp) Reset() {
	*x = GetTaskMetricRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskMetricRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskMetricRsp) ProtoMessage() {}

func (x *GetTaskMetricRsp) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskMetricRsp.ProtoReflect.Descriptor instead.
func (*GetTaskMetricRsp) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{6}
}

func (x *GetTaskMetricRsp) GetCpuStats() *CpuStats {
	if x != nil {
		return x.CpuStats
	}
	return nil
}

func (x *GetTaskMetricRsp) GetMemoryStats() *MemoryStats {
	if x != nil {
		return x.MemoryStats
	}
	return nil
}

func (x *GetTaskMetricRsp) GetGpuStats() *GpuStats {
	if x != nil {
		return x.GpuStats
	}
	return nil
}

type CpuStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usage float32 `protobuf:"fixed32,1,opt,name=usage,proto3" json:"usage,omitempty"`
}

func (x *CpuStats) Reset() {
	*x = CpuStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuStats) ProtoMessage() {}

func (x *CpuStats) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuStats.ProtoReflect.Descriptor instead.
func (*CpuStats) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{7}
}

func (x *CpuStats) GetUsage() float32 {
	if x != nil {
		return x.Usage
	}
	return 0
}

type MemoryStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usage float32 `protobuf:"fixed32,1,opt,name=usage,proto3" json:"usage,omitempty"`
	Used  uint64  `protobuf:"varint,2,opt,name=used,proto3" json:"used,omitempty"`
	Free  uint64  `protobuf:"varint,3,opt,name=free,proto3" json:"free,omitempty"`
}

func (x *MemoryStats) Reset() {
	*x = MemoryStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryStats) ProtoMessage() {}

func (x *MemoryStats) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryStats.ProtoReflect.Descriptor instead.
func (*MemoryStats) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{8}
}

func (x *MemoryStats) GetUsage() float32 {
	if x != nil {
		return x.Usage
	}
	return 0
}

func (x *MemoryStats) GetUsed() uint64 {
	if x != nil {
		return x.Used
	}
	return 0
}

func (x *MemoryStats) GetFree() uint64 {
	if x != nil {
		return x.Free
	}
	return 0
}

type GpuStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CudaVersion      string              `protobuf:"bytes,1,opt,name=cuda_version,json=cudaVersion,proto3" json:"cuda_version,omitempty"`
	GpuInstanceStats []*GpuInstanceStats `protobuf:"bytes,2,rep,name=gpu_instance_stats,json=gpuInstanceStats,proto3" json:"gpu_instance_stats,omitempty"`
}

func (x *GpuStats) Reset() {
	*x = GpuStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GpuStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuStats) ProtoMessage() {}

func (x *GpuStats) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuStats.ProtoReflect.Descriptor instead.
func (*GpuStats) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{9}
}

func (x *GpuStats) GetCudaVersion() string {
	if x != nil {
		return x.CudaVersion
	}
	return ""
}

func (x *GpuStats) GetGpuInstanceStats() []*GpuInstanceStats {
	if x != nil {
		return x.GpuInstanceStats
	}
	return nil
}

type GpuInstanceStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductName string  `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	GpuUsage    float32 `protobuf:"fixed32,3,opt,name=gpu_usage,json=gpuUsage,proto3" json:"gpu_usage,omitempty"`
	MemoryUsage float32 `protobuf:"fixed32,4,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	MemoryUsed  uint64  `protobuf:"varint,5,opt,name=memory_used,json=memoryUsed,proto3" json:"memory_used,omitempty"`
	MemoryFree  uint64  `protobuf:"varint,6,opt,name=memory_free,json=memoryFree,proto3" json:"memory_free,omitempty"`
}

func (x *GpuInstanceStats) Reset() {
	*x = GpuInstanceStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GpuInstanceStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuInstanceStats) ProtoMessage() {}

func (x *GpuInstanceStats) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuInstanceStats.ProtoReflect.Descriptor instead.
func (*GpuInstanceStats) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{10}
}

func (x *GpuInstanceStats) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GpuInstanceStats) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *GpuInstanceStats) GetGpuUsage() float32 {
	if x != nil {
		return x.GpuUsage
	}
	return 0
}

func (x *GpuInstanceStats) GetMemoryUsage() float32 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *GpuInstanceStats) GetMemoryUsed() uint64 {
	if x != nil {
		return x.MemoryUsed
	}
	return 0
}

func (x *GpuInstanceStats) GetMemoryFree() uint64 {
	if x != nil {
		return x.MemoryFree
	}
	return 0
}

var File_worker_proto protoreflect.FileDescriptor

var file_worker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x22, 0x8d, 0x01, 0x0a, 0x09, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x72, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0xb2, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a,
	0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x73, 0x70, 0x22, 0x26, 0x0a, 0x0b, 0x53,
	0x74, 0x6f, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x73, 0x70, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x22, 0xa8, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x09, 0x63,
	0x70, 0x75, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x70, 0x75, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x08, 0x63, 0x70, 0x75, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x0c, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x67, 0x70, 0x75, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x47,
	0x70, 0x75, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x08, 0x67, 0x70, 0x75, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x22, 0x20, 0x0a, 0x08, 0x43, 0x70, 0x75, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x4b, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66, 0x72, 0x65, 0x65,
	0x22, 0x75, 0x0a, 0x08, 0x47, 0x70, 0x75, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x75, 0x64, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x64, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x46, 0x0a, 0x12, 0x67, 0x70, 0x75, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x70, 0x75, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x10, 0x67, 0x70, 0x75, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0xc7, 0x01, 0x0a, 0x10, 0x47, 0x70, 0x75, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x67, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x46, 0x72, 0x65,
	0x65, 0x32, 0xbc, 0x01, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a,
	0x14, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x13, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x6f, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x18, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x73, 0x70,
	0x42, 0x41, 0x5a, 0x3f, 0x68, 0x61, 0x6e, 0x67, 0x6c, 0x6f, 0x6b, 0x2d, 0x74, 0x65, 0x63, 0x68,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x3b, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_worker_proto_rawDescOnce sync.Once
	file_worker_proto_rawDescData = file_worker_proto_rawDesc
)

func file_worker_proto_rawDescGZIP() []byte {
	file_worker_proto_rawDescOnce.Do(func() {
		file_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_worker_proto_rawDescData)
	})
	return file_worker_proto_rawDescData
}

var file_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_worker_proto_goTypes = []interface{}{
	(*Algorithm)(nil),        // 0: worker.Algorithm
	(*StartTaskReq)(nil),     // 1: worker.StartTaskReq
	(*StartTaskRsp)(nil),     // 2: worker.StartTaskRsp
	(*StopTaskReq)(nil),      // 3: worker.StopTaskReq
	(*StopTaskRsp)(nil),      // 4: worker.StopTaskRsp
	(*GetTaskMetricReq)(nil), // 5: worker.GetTaskMetricReq
	(*GetTaskMetricRsp)(nil), // 6: worker.GetTaskMetricRsp
	(*CpuStats)(nil),         // 7: worker.CpuStats
	(*MemoryStats)(nil),      // 8: worker.MemoryStats
	(*GpuStats)(nil),         // 9: worker.GpuStats
	(*GpuInstanceStats)(nil), // 10: worker.GpuInstanceStats
}
var file_worker_proto_depIdxs = []int32{
	0,  // 0: worker.StartTaskReq.algorithm:type_name -> worker.Algorithm
	7,  // 1: worker.GetTaskMetricRsp.cpu_stats:type_name -> worker.CpuStats
	8,  // 2: worker.GetTaskMetricRsp.memory_stats:type_name -> worker.MemoryStats
	9,  // 3: worker.GetTaskMetricRsp.gpu_stats:type_name -> worker.GpuStats
	10, // 4: worker.GpuStats.gpu_instance_stats:type_name -> worker.GpuInstanceStats
	1,  // 5: worker.Worker.StartTask:input_type -> worker.StartTaskReq
	3,  // 6: worker.Worker.StopTask:input_type -> worker.StopTaskReq
	5,  // 7: worker.Worker.GetTaskMetric:input_type -> worker.GetTaskMetricReq
	2,  // 8: worker.Worker.StartTask:output_type -> worker.StartTaskRsp
	4,  // 9: worker.Worker.StopTask:output_type -> worker.StopTaskRsp
	6,  // 10: worker.Worker.GetTaskMetric:output_type -> worker.GetTaskMetricRsp
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_worker_proto_init() }
func file_worker_proto_init() {
	if File_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Algorithm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTaskReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTaskRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopTaskReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopTaskRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskMetricReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskMetricRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GpuStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GpuInstanceStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_worker_proto_goTypes,
		DependencyIndexes: file_worker_proto_depIdxs,
		MessageInfos:      file_worker_proto_msgTypes,
	}.Build()
	File_worker_proto = out.File
	file_worker_proto_rawDesc = nil
	file_worker_proto_goTypes = nil
	file_worker_proto_depIdxs = nil
}
