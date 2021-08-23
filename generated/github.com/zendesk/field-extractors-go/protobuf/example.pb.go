// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: example_proto/example.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Kind int32

const (
	Kind_UNKOWN Kind = 0
	Kind_KINDA  Kind = 1
	Kind_KINDB  Kind = 2
)

// Enum value maps for Kind.
var (
	Kind_name = map[int32]string{
		0: "UNKOWN",
		1: "KINDA",
		2: "KINDB",
	}
	Kind_value = map[string]int32{
		"UNKOWN": 0,
		"KINDA":  1,
		"KINDB":  2,
	}
)

func (x Kind) Enum() *Kind {
	p := new(Kind)
	*p = x
	return p
}

func (x Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_example_proto_example_proto_enumTypes[0].Descriptor()
}

func (Kind) Type() protoreflect.EnumType {
	return &file_example_proto_example_proto_enumTypes[0]
}

func (x Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Kind.Descriptor instead.
func (Kind) EnumDescriptor() ([]byte, []int) {
	return file_example_proto_example_proto_rawDescGZIP(), []int{0}
}

type LineItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *LineItem) Reset() {
	*x = LineItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItem) ProtoMessage() {}

func (x *LineItem) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItem.ProtoReflect.Descriptor instead.
func (*LineItem) Descriptor() ([]byte, []int) {
	return file_example_proto_example_proto_rawDescGZIP(), []int{0}
}

func (x *LineItem) GetId() *wrapperspb.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *LineItem) GetQuantity() *wrapperspb.Int64Value {
	if x != nil {
		return x.Quantity
	}
	return nil
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_example_proto_example_proto_rawDescGZIP(), []int{1}
}

func (x *Cart) GetId() *wrapperspb.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

type OrderUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cart *Cart                  `protobuf:"bytes,2,opt,name=cart,proto3" json:"cart,omitempty"`
}

func (x *OrderUpdated) Reset() {
	*x = OrderUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpdated) ProtoMessage() {}

func (x *OrderUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUpdated.ProtoReflect.Descriptor instead.
func (*OrderUpdated) Descriptor() ([]byte, []int) {
	return file_example_proto_example_proto_rawDescGZIP(), []int{2}
}

func (x *OrderUpdated) GetId() *wrapperspb.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *OrderUpdated) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type OrderCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cart *Cart                  `protobuf:"bytes,2,opt,name=cart,proto3" json:"cart,omitempty"`
}

func (x *OrderCreated) Reset() {
	*x = OrderCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreated) ProtoMessage() {}

func (x *OrderCreated) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreated.ProtoReflect.Descriptor instead.
func (*OrderCreated) Descriptor() ([]byte, []int) {
	return file_example_proto_example_proto_rawDescGZIP(), []int{3}
}

func (x *OrderCreated) GetId() *wrapperspb.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *OrderCreated) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *wrapperspb.Int64Value  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token      *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Properties *structpb.Struct        `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
	Lineitems  []*LineItem             `protobuf:"bytes,4,rep,name=lineitems,proto3" json:"lineitems,omitempty"`
	// Types that are assignable to Event:
	//	*Order_Orderupdated
	//	*Order_Ordercreated
	Event  isOrder_Event `protobuf_oneof:"event"`
	List   []string      `protobuf:"bytes,7,rep,name=list,proto3" json:"list,omitempty"`
	Cart   *Cart         `protobuf:"bytes,8,opt,name=cart,proto3" json:"cart,omitempty"`
	Opaque string        `protobuf:"bytes,9,opt,name=opaque,proto3" json:"opaque,omitempty"`
	Kind   Kind          `protobuf:"varint,10,opt,name=kind,proto3,enum=zendesk.protobuf.example.Kind" json:"kind,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_example_proto_example_proto_rawDescGZIP(), []int{4}
}

func (x *Order) GetId() *wrapperspb.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Order) GetToken() *wrapperspb.StringValue {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *Order) GetProperties() *structpb.Struct {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Order) GetLineitems() []*LineItem {
	if x != nil {
		return x.Lineitems
	}
	return nil
}

func (m *Order) GetEvent() isOrder_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *Order) GetOrderupdated() *OrderUpdated {
	if x, ok := x.GetEvent().(*Order_Orderupdated); ok {
		return x.Orderupdated
	}
	return nil
}

func (x *Order) GetOrdercreated() *OrderCreated {
	if x, ok := x.GetEvent().(*Order_Ordercreated); ok {
		return x.Ordercreated
	}
	return nil
}

func (x *Order) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *Order) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

func (x *Order) GetOpaque() string {
	if x != nil {
		return x.Opaque
	}
	return ""
}

func (x *Order) GetKind() Kind {
	if x != nil {
		return x.Kind
	}
	return Kind_UNKOWN
}

type isOrder_Event interface {
	isOrder_Event()
}

type Order_Orderupdated struct {
	Orderupdated *OrderUpdated `protobuf:"bytes,5,opt,name=orderupdated,proto3,oneof"`
}

type Order_Ordercreated struct {
	Ordercreated *OrderCreated `protobuf:"bytes,6,opt,name=ordercreated,proto3,oneof"`
}

func (*Order_Orderupdated) isOrder_Event() {}

func (*Order_Ordercreated) isOrder_Event() {}

var File_example_proto_example_proto protoreflect.FileDescriptor

var file_example_proto_example_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x7a,
	0x65, 0x6e, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x33, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6f, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x63, 0x61, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x7a, 0x65, 0x6e, 0x64, 0x65, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x6f, 0x0a,
	0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2b, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x63, 0x61,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x7a, 0x65, 0x6e, 0x64, 0x65,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x9c,
	0x04, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x40, 0x0a, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x7a, 0x65, 0x6e, 0x64, 0x65, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x4c, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x7a, 0x65, 0x6e,
	0x64, 0x65, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x4c, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x7a, 0x65, 0x6e, 0x64, 0x65,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x7a, 0x65, 0x6e, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x12, 0x32, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x7a, 0x65, 0x6e, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2a, 0x28, 0x0a,
	0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x4b, 0x49, 0x4e, 0x44, 0x41, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x4b, 0x49, 0x4e, 0x44, 0x42, 0x10, 0x02, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x65, 0x6e, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x2d, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2d, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_example_proto_example_proto_rawDescOnce sync.Once
	file_example_proto_example_proto_rawDescData = file_example_proto_example_proto_rawDesc
)

func file_example_proto_example_proto_rawDescGZIP() []byte {
	file_example_proto_example_proto_rawDescOnce.Do(func() {
		file_example_proto_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_example_proto_rawDescData)
	})
	return file_example_proto_example_proto_rawDescData
}

var file_example_proto_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_example_proto_example_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_example_proto_example_proto_goTypes = []interface{}{
	(Kind)(0),                      // 0: zendesk.protobuf.example.Kind
	(*LineItem)(nil),               // 1: zendesk.protobuf.example.LineItem
	(*Cart)(nil),                   // 2: zendesk.protobuf.example.Cart
	(*OrderUpdated)(nil),           // 3: zendesk.protobuf.example.OrderUpdated
	(*OrderCreated)(nil),           // 4: zendesk.protobuf.example.OrderCreated
	(*Order)(nil),                  // 5: zendesk.protobuf.example.Order
	(*wrapperspb.Int64Value)(nil),  // 6: google.protobuf.Int64Value
	(*wrapperspb.StringValue)(nil), // 7: google.protobuf.StringValue
	(*structpb.Struct)(nil),        // 8: google.protobuf.Struct
}
var file_example_proto_example_proto_depIdxs = []int32{
	6,  // 0: zendesk.protobuf.example.LineItem.id:type_name -> google.protobuf.Int64Value
	6,  // 1: zendesk.protobuf.example.LineItem.quantity:type_name -> google.protobuf.Int64Value
	6,  // 2: zendesk.protobuf.example.Cart.id:type_name -> google.protobuf.Int64Value
	6,  // 3: zendesk.protobuf.example.OrderUpdated.id:type_name -> google.protobuf.Int64Value
	2,  // 4: zendesk.protobuf.example.OrderUpdated.cart:type_name -> zendesk.protobuf.example.Cart
	6,  // 5: zendesk.protobuf.example.OrderCreated.id:type_name -> google.protobuf.Int64Value
	2,  // 6: zendesk.protobuf.example.OrderCreated.cart:type_name -> zendesk.protobuf.example.Cart
	6,  // 7: zendesk.protobuf.example.Order.id:type_name -> google.protobuf.Int64Value
	7,  // 8: zendesk.protobuf.example.Order.token:type_name -> google.protobuf.StringValue
	8,  // 9: zendesk.protobuf.example.Order.properties:type_name -> google.protobuf.Struct
	1,  // 10: zendesk.protobuf.example.Order.lineitems:type_name -> zendesk.protobuf.example.LineItem
	3,  // 11: zendesk.protobuf.example.Order.orderupdated:type_name -> zendesk.protobuf.example.OrderUpdated
	4,  // 12: zendesk.protobuf.example.Order.ordercreated:type_name -> zendesk.protobuf.example.OrderCreated
	2,  // 13: zendesk.protobuf.example.Order.cart:type_name -> zendesk.protobuf.example.Cart
	0,  // 14: zendesk.protobuf.example.Order.kind:type_name -> zendesk.protobuf.example.Kind
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_example_proto_example_proto_init() }
func file_example_proto_example_proto_init() {
	if File_example_proto_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineItem); i {
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
		file_example_proto_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
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
		file_example_proto_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderUpdated); i {
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
		file_example_proto_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreated); i {
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
		file_example_proto_example_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
	file_example_proto_example_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Order_Orderupdated)(nil),
		(*Order_Ordercreated)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_proto_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_proto_example_proto_goTypes,
		DependencyIndexes: file_example_proto_example_proto_depIdxs,
		EnumInfos:         file_example_proto_example_proto_enumTypes,
		MessageInfos:      file_example_proto_example_proto_msgTypes,
	}.Build()
	File_example_proto_example_proto = out.File
	file_example_proto_example_proto_rawDesc = nil
	file_example_proto_example_proto_goTypes = nil
	file_example_proto_example_proto_depIdxs = nil
}
