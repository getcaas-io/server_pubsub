package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Identity struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Identity) Reset()                    { *m = Identity{} }
func (m *Identity) String() string            { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()               {}
func (*Identity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Identity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Message struct {
	Data       []byte            `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Attributes map[string]string `protobuf:"bytes,2,rep,name=attributes" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Id         string            `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PublishRequest struct {
	Key      string     `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Messages []*Message `protobuf:"bytes,2,rep,name=messages" json:"messages,omitempty"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (m *PublishRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PublishRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PublishRequest) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type PublishResponse struct {
	MessageIds []string `protobuf:"bytes,1,rep,name=message_ids,json=messageIds" json:"message_ids,omitempty"`
}

func (m *PublishResponse) Reset()                    { *m = PublishResponse{} }
func (m *PublishResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()               {}
func (*PublishResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PublishResponse) GetMessageIds() []string {
	if m != nil {
		return m.MessageIds
	}
	return nil
}

type SubscribeRequest struct {
	Identity     *Identity     `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	Subscription *Subscription `protobuf:"bytes,2,opt,name=subscription" json:"subscription,omitempty"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SubscribeRequest) GetIdentity() *Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *SubscribeRequest) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

type Subscription struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Subscription) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Identity)(nil), "Identity")
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*PublishRequest)(nil), "PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "PublishResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "SubscribeRequest")
	proto.RegisterType((*Subscription)(nil), "Subscription")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Subscriber service

type SubscriberClient interface {
	Authenticate(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Identity, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*Subscription, error)
	Unsubscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*Subscription, error)
	Pull(ctx context.Context, in *Identity, opts ...grpc.CallOption) (Subscriber_PullClient, error)
}

type subscriberClient struct {
	cc *grpc.ClientConn
}

func NewSubscriberClient(cc *grpc.ClientConn) SubscriberClient {
	return &subscriberClient{cc}
}

func (c *subscriberClient) Authenticate(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := grpc.Invoke(ctx, "/Subscriber/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriberClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*Subscription, error) {
	out := new(Subscription)
	err := grpc.Invoke(ctx, "/Subscriber/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriberClient) Unsubscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*Subscription, error) {
	out := new(Subscription)
	err := grpc.Invoke(ctx, "/Subscriber/Unsubscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriberClient) Pull(ctx context.Context, in *Identity, opts ...grpc.CallOption) (Subscriber_PullClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Subscriber_serviceDesc.Streams[0], c.cc, "/Subscriber/Pull", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriberPullClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Subscriber_PullClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type subscriberPullClient struct {
	grpc.ClientStream
}

func (x *subscriberPullClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Subscriber service

type SubscriberServer interface {
	Authenticate(context.Context, *Identity) (*Identity, error)
	Subscribe(context.Context, *SubscribeRequest) (*Subscription, error)
	Unsubscribe(context.Context, *SubscribeRequest) (*Subscription, error)
	Pull(*Identity, Subscriber_PullServer) error
}

func RegisterSubscriberServer(s *grpc.Server, srv SubscriberServer) {
	s.RegisterService(&_Subscriber_serviceDesc, srv)
}

func _Subscriber_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Subscriber/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberServer).Authenticate(ctx, req.(*Identity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscriber_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Subscriber/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscriber_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Subscriber/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberServer).Unsubscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscriber_Pull_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Identity)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriberServer).Pull(m, &subscriberPullServer{stream})
}

type Subscriber_PullServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type subscriberPullServer struct {
	grpc.ServerStream
}

func (x *subscriberPullServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _Subscriber_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Subscriber",
	HandlerType: (*SubscriberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _Subscriber_Authenticate_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _Subscriber_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _Subscriber_Unsubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pull",
			Handler:       _Subscriber_Pull_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "PubSub.proto",
}

// Client API for Publisher service

type PublisherClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
}

type publisherClient struct {
	cc *grpc.ClientConn
}

func NewPublisherClient(cc *grpc.ClientConn) PublisherClient {
	return &publisherClient{cc}
}

func (c *publisherClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := grpc.Invoke(ctx, "/Publisher/Publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Publisher service

type PublisherServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
}

func RegisterPublisherServer(s *grpc.Server, srv PublisherServer) {
	s.RegisterService(&_Publisher_serviceDesc, srv)
}

func _Publisher_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Publisher/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Publisher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Publisher",
	HandlerType: (*PublisherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Publisher_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "PubSub.proto",
}

func init() { proto.RegisterFile("PubSub.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0xfd, 0x87, 0x16, 0x7b, 0xec, 0x02, 0x5d, 0xf5, 0x60, 0xf9, 0x50, 0xac, 0x55, 0x5b,
	0x71, 0xda, 0x16, 0xf7, 0x82, 0x5a, 0xf5, 0xc0, 0xa1, 0x52, 0x38, 0x44, 0x42, 0x8b, 0x72, 0x8e,
	0x6c, 0xbc, 0x0a, 0xab, 0x18, 0x9b, 0x78, 0x77, 0x23, 0xf1, 0x42, 0x79, 0x87, 0xbc, 0x5d, 0xc4,
	0x62, 0x1b, 0x43, 0x72, 0xc9, 0x6d, 0x76, 0xbe, 0xf9, 0x3c, 0xbf, 0x99, 0x31, 0xf8, 0x4b, 0x95,
	0xae, 0x54, 0x4a, 0x76, 0x55, 0x29, 0x4b, 0xfc, 0x15, 0x9c, 0x45, 0xc6, 0x0a, 0xc9, 0xe5, 0x1e,
	0x21, 0xe8, 0x15, 0xc9, 0x96, 0x05, 0x66, 0x64, 0x4e, 0x5c, 0xaa, 0x63, 0xfc, 0x64, 0x42, 0xff,
	0x9a, 0x09, 0x91, 0xdc, 0xb1, 0x83, 0x9e, 0x25, 0x32, 0xd1, 0xba, 0x4f, 0x75, 0x8c, 0x66, 0x00,
	0x89, 0x94, 0x15, 0x4f, 0x95, 0x64, 0x22, 0xb0, 0x22, 0x7b, 0xe2, 0xc5, 0x01, 0xa9, 0x1d, 0x64,
	0xde, 0x4a, 0xff, 0x0b, 0x59, 0xed, 0x69, 0xa7, 0x16, 0x0d, 0xc0, 0xe2, 0x59, 0x60, 0xeb, 0x5e,
	0x16, 0xcf, 0xc2, 0x7f, 0x30, 0xbc, 0x28, 0x47, 0x23, 0xb0, 0xef, 0xd9, 0xbe, 0xe6, 0x39, 0x84,
	0xe8, 0x0b, 0x7c, 0x78, 0x4c, 0x72, 0xc5, 0x02, 0x4b, 0xe7, 0x8e, 0x8f, 0x3f, 0xd6, 0xcc, 0xc4,
	0x57, 0x30, 0x58, 0xaa, 0x34, 0xe7, 0x62, 0x43, 0xd9, 0x83, 0x62, 0x42, 0xbe, 0xe1, 0xfe, 0x06,
	0xce, 0xf6, 0x48, 0xd6, 0xa0, 0x3a, 0x0d, 0x2a, 0x6d, 0x15, 0x1c, 0xc3, 0xb0, 0xfd, 0x92, 0xd8,
	0x95, 0x85, 0x60, 0x68, 0x0c, 0x5e, 0x2d, 0xdf, 0xf2, 0x4c, 0x04, 0x66, 0x64, 0x4f, 0x5c, 0x0a,
	0x75, 0x6a, 0x91, 0x09, 0x9c, 0xc3, 0x68, 0xa5, 0x52, 0xb1, 0xae, 0x78, 0xca, 0x9a, 0xfe, 0xdf,
	0xc1, 0xe1, 0xf5, 0x6a, 0x35, 0x84, 0x17, 0xbb, 0xa4, 0xd9, 0x35, 0x6d, 0x25, 0x34, 0x05, 0x5f,
	0x1c, 0xad, 0x3b, 0xc9, 0xcb, 0x42, 0x4f, 0xe6, 0xc5, 0x9f, 0xc8, 0xaa, 0x93, 0xa4, 0x67, 0x25,
	0x38, 0x02, 0xbf, 0xab, 0xbe, 0x9e, 0x34, 0x7e, 0x36, 0x01, 0x5a, 0xa0, 0x0a, 0xfd, 0x00, 0x7f,
	0xae, 0xe4, 0xe6, 0xd0, 0x71, 0x9d, 0x48, 0x86, 0x4e, 0x20, 0xe1, 0x29, 0xc4, 0x06, 0xfa, 0x09,
	0x6e, 0xeb, 0x42, 0x9f, 0xc9, 0xe5, 0x48, 0xe1, 0x39, 0x15, 0x36, 0xd0, 0x14, 0xbc, 0x9b, 0x42,
	0xbc, 0xcb, 0x32, 0x86, 0xde, 0x52, 0xe5, 0x79, 0x97, 0xa1, 0xbd, 0x02, 0x36, 0x7e, 0x99, 0xf1,
	0x5f, 0x70, 0xeb, 0xfd, 0xb3, 0x0a, 0x11, 0xe8, 0xd7, 0x0f, 0x34, 0x24, 0xe7, 0x07, 0x0e, 0x47,
	0xe4, 0xe2, 0x4e, 0xd8, 0x48, 0x3f, 0xea, 0xdf, 0xfa, 0xf7, 0x4b, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xee, 0x48, 0x08, 0x9d, 0xe6, 0x02, 0x00, 0x00,
}
