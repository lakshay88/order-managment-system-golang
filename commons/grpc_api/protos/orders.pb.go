// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: orders.proto

package grpc_api

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

// Message definition for an Order
type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CustomerID  string  `protobuf:"bytes,2,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Status      string  `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Items       []*Item `protobuf:"bytes,4,rep,name=Items,proto3" json:"Items,omitempty"`
	TotalAmount float64 `protobuf:"fixed64,5,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_orders_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[0]
	if x != nil {
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
	return file_orders_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Order) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

// Message definition for an Item
type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Quantity int32   `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Price    float64 `protobuf:"fixed64,4,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_orders_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Item) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

// Request message for getting a specific order
type GetOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID    string `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	CustomerID string `protobuf:"bytes,2,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_orders_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *GetOrderRequest) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

// Request message for creating an order
type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerID  string  `protobuf:"bytes,1,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Items       []*Item `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	TotalAmount float64 `protobuf:"fixed64,3,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_orders_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderRequest) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *CreateOrderRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreateOrderRequest) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

// Response message for returning all orders
type GetAllOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetAllOrdersResponse) Reset() {
	*x = GetAllOrdersResponse{}
	mi := &file_orders_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrdersResponse) ProtoMessage() {}

func (x *GetAllOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrdersResponse.ProtoReflect.Descriptor instead.
func (*GetAllOrdersResponse) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_orders_proto protoreflect.FileDescriptor

var file_orders_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e,
	0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1b, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x5c, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x4b, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x22, 0x73, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1b, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x36, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x32, 0xb9, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x10,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x06, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x61, 0x6b, 0x73, 0x68, 0x61, 0x79, 0x38, 0x38, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_orders_proto_rawDescOnce sync.Once
	file_orders_proto_rawDescData = file_orders_proto_rawDesc
)

func file_orders_proto_rawDescGZIP() []byte {
	file_orders_proto_rawDescOnce.Do(func() {
		file_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_proto_rawDescData)
	})
	return file_orders_proto_rawDescData
}

var file_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_orders_proto_goTypes = []any{
	(*Order)(nil),                // 0: Order
	(*Item)(nil),                 // 1: Item
	(*GetOrderRequest)(nil),      // 2: GetOrderRequest
	(*CreateOrderRequest)(nil),   // 3: CreateOrderRequest
	(*GetAllOrdersResponse)(nil), // 4: GetAllOrdersResponse
}
var file_orders_proto_depIdxs = []int32{
	1, // 0: Order.Items:type_name -> Item
	1, // 1: CreateOrderRequest.Items:type_name -> Item
	0, // 2: GetAllOrdersResponse.orders:type_name -> Order
	2, // 3: OrderService.GetAllOrders:input_type -> GetOrderRequest
	2, // 4: OrderService.GetOrders:input_type -> GetOrderRequest
	3, // 5: OrderService.CreateOrder:input_type -> CreateOrderRequest
	0, // 6: OrderService.UpdateOrder:input_type -> Order
	4, // 7: OrderService.GetAllOrders:output_type -> GetAllOrdersResponse
	0, // 8: OrderService.GetOrders:output_type -> Order
	0, // 9: OrderService.CreateOrder:output_type -> Order
	0, // 10: OrderService.UpdateOrder:output_type -> Order
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_orders_proto_init() }
func file_orders_proto_init() {
	if File_orders_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_proto_goTypes,
		DependencyIndexes: file_orders_proto_depIdxs,
		MessageInfos:      file_orders_proto_msgTypes,
	}.Build()
	File_orders_proto = out.File
	file_orders_proto_rawDesc = nil
	file_orders_proto_goTypes = nil
	file_orders_proto_depIdxs = nil
}