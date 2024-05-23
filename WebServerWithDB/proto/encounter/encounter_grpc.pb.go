// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: encounter.proto

package encounter

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Encounter_CreateSocialEncounter_FullMethodName = "/Encounter/CreateSocialEncounter"
)

// EncounterClient is the client API for Encounter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EncounterClient interface {
	CreateSocialEncounter(ctx context.Context, in *WholeSocialEncounterMongoDto, opts ...grpc.CallOption) (*WholeSocialEncounterMongoDto, error)
}

type encounterClient struct {
	cc grpc.ClientConnInterface
}

func NewEncounterClient(cc grpc.ClientConnInterface) EncounterClient {
	return &encounterClient{cc}
}

func (c *encounterClient) CreateSocialEncounter(ctx context.Context, in *WholeSocialEncounterMongoDto, opts ...grpc.CallOption) (*WholeSocialEncounterMongoDto, error) {
	out := new(WholeSocialEncounterMongoDto)
	err := c.cc.Invoke(ctx, Encounter_CreateSocialEncounter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncounterServer is the server API for Encounter service.
// All implementations must embed UnimplementedEncounterServer
// for forward compatibility
type EncounterServer interface {
	CreateSocialEncounter(context.Context, *WholeSocialEncounterMongoDto) (*WholeSocialEncounterMongoDto, error)
	mustEmbedUnimplementedEncounterServer()
}

// UnimplementedEncounterServer must be embedded to have forward compatible implementations.
type UnimplementedEncounterServer struct {
}

func (UnimplementedEncounterServer) CreateSocialEncounter(context.Context, *WholeSocialEncounterMongoDto) (*WholeSocialEncounterMongoDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSocialEncounter not implemented")
}
func (UnimplementedEncounterServer) mustEmbedUnimplementedEncounterServer() {}

// UnsafeEncounterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EncounterServer will
// result in compilation errors.
type UnsafeEncounterServer interface {
	mustEmbedUnimplementedEncounterServer()
}

func RegisterEncounterServer(s grpc.ServiceRegistrar, srv EncounterServer) {
	s.RegisterService(&Encounter_ServiceDesc, srv)
}

func _Encounter_CreateSocialEncounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WholeSocialEncounterMongoDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServer).CreateSocialEncounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Encounter_CreateSocialEncounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServer).CreateSocialEncounter(ctx, req.(*WholeSocialEncounterMongoDto))
	}
	return interceptor(ctx, in, info, handler)
}

// Encounter_ServiceDesc is the grpc.ServiceDesc for Encounter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Encounter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Encounter",
	HandlerType: (*EncounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSocialEncounter",
			Handler:    _Encounter_CreateSocialEncounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "encounter.proto",
}
