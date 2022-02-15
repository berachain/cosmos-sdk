// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: cosmos/group/v1beta1/tx.proto

package groupv1beta1

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateGroup creates a new group with an admin account address, a list of members and some optional metadata.
	CreateGroup(ctx context.Context, in *MsgCreateGroup, opts ...grpc.CallOption) (*MsgCreateGroupResponse, error)
	// UpdateGroupMembers updates the group members with given group id and admin address.
	UpdateGroupMembers(ctx context.Context, in *MsgUpdateGroupMembers, opts ...grpc.CallOption) (*MsgUpdateGroupMembersResponse, error)
	// UpdateGroupAdmin updates the group admin with given group id and previous admin address.
	UpdateGroupAdmin(ctx context.Context, in *MsgUpdateGroupAdmin, opts ...grpc.CallOption) (*MsgUpdateGroupAdminResponse, error)
	// UpdateGroupMetadata updates the group metadata with given group id and admin address.
	UpdateGroupMetadata(ctx context.Context, in *MsgUpdateGroupMetadata, opts ...grpc.CallOption) (*MsgUpdateGroupMetadataResponse, error)
	// CreateGroupPolicy creates a new group policy using given DecisionPolicy.
	CreateGroupPolicy(ctx context.Context, in *MsgCreateGroupPolicy, opts ...grpc.CallOption) (*MsgCreateGroupPolicyResponse, error)
	// UpdateGroupPolicyAdmin updates a group policy admin.
	UpdateGroupPolicyAdmin(ctx context.Context, in *MsgUpdateGroupPolicyAdmin, opts ...grpc.CallOption) (*MsgUpdateGroupPolicyAdminResponse, error)
	// UpdateGroupPolicyDecisionPolicy allows a group policy's decision policy to be updated.
	UpdateGroupPolicyDecisionPolicy(ctx context.Context, in *MsgUpdateGroupPolicyDecisionPolicy, opts ...grpc.CallOption) (*MsgUpdateGroupPolicyDecisionPolicyResponse, error)
	// UpdateGroupPolicyMetadata updates a group policy metadata.
	UpdateGroupPolicyMetadata(ctx context.Context, in *MsgUpdateGroupPolicyMetadata, opts ...grpc.CallOption) (*MsgUpdateGroupPolicyMetadataResponse, error)
	// SubmitProposal submits a new proposal.
	SubmitProposal(ctx context.Context, in *MsgSubmitProposal, opts ...grpc.CallOption) (*MsgSubmitProposalResponse, error)
	// WithdrawProposal aborts a proposal.
	WithdrawProposal(ctx context.Context, in *MsgWithdrawProposal, opts ...grpc.CallOption) (*MsgWithdrawProposalResponse, error)
	// Vote allows a voter to vote on a proposal.
	Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error)
	// VoteWeighted defines a method to add a weighted vote on a specific proposal.
	VoteWeighted(ctx context.Context, in *MsgVoteWeighted, opts ...grpc.CallOption) (*MsgVoteWeightedResponse, error)
	// Exec executes a proposal.
	Exec(ctx context.Context, in *MsgExec, opts ...grpc.CallOption) (*MsgExecResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateGroup(ctx context.Context, in *MsgCreateGroup, opts ...grpc.CallOption) (*MsgCreateGroupResponse, error) {
	out := new(MsgCreateGroupResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGroupMembers(ctx context.Context, in *MsgUpdateGroupMembers, opts ...grpc.CallOption) (*MsgUpdateGroupMembersResponse, error) {
	out := new(MsgUpdateGroupMembersResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/UpdateGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGroupAdmin(ctx context.Context, in *MsgUpdateGroupAdmin, opts ...grpc.CallOption) (*MsgUpdateGroupAdminResponse, error) {
	out := new(MsgUpdateGroupAdminResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/UpdateGroupAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGroupMetadata(ctx context.Context, in *MsgUpdateGroupMetadata, opts ...grpc.CallOption) (*MsgUpdateGroupMetadataResponse, error) {
	out := new(MsgUpdateGroupMetadataResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/UpdateGroupMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateGroupPolicy(ctx context.Context, in *MsgCreateGroupPolicy, opts ...grpc.CallOption) (*MsgCreateGroupPolicyResponse, error) {
	out := new(MsgCreateGroupPolicyResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/CreateGroupPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGroupPolicyAdmin(ctx context.Context, in *MsgUpdateGroupPolicyAdmin, opts ...grpc.CallOption) (*MsgUpdateGroupPolicyAdminResponse, error) {
	out := new(MsgUpdateGroupPolicyAdminResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/UpdateGroupPolicyAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGroupPolicyDecisionPolicy(ctx context.Context, in *MsgUpdateGroupPolicyDecisionPolicy, opts ...grpc.CallOption) (*MsgUpdateGroupPolicyDecisionPolicyResponse, error) {
	out := new(MsgUpdateGroupPolicyDecisionPolicyResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/UpdateGroupPolicyDecisionPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGroupPolicyMetadata(ctx context.Context, in *MsgUpdateGroupPolicyMetadata, opts ...grpc.CallOption) (*MsgUpdateGroupPolicyMetadataResponse, error) {
	out := new(MsgUpdateGroupPolicyMetadataResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/UpdateGroupPolicyMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitProposal(ctx context.Context, in *MsgSubmitProposal, opts ...grpc.CallOption) (*MsgSubmitProposalResponse, error) {
	out := new(MsgSubmitProposalResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/SubmitProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawProposal(ctx context.Context, in *MsgWithdrawProposal, opts ...grpc.CallOption) (*MsgWithdrawProposalResponse, error) {
	out := new(MsgWithdrawProposalResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/WithdrawProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error) {
	out := new(MsgVoteResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) VoteWeighted(ctx context.Context, in *MsgVoteWeighted, opts ...grpc.CallOption) (*MsgVoteWeightedResponse, error) {
	out := new(MsgVoteWeightedResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/VoteWeighted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Exec(ctx context.Context, in *MsgExec, opts ...grpc.CallOption) (*MsgExecResponse, error) {
	out := new(MsgExecResponse)
	err := c.cc.Invoke(ctx, "/cosmos.group.v1beta1.Msg/Exec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateGroup creates a new group with an admin account address, a list of members and some optional metadata.
	CreateGroup(context.Context, *MsgCreateGroup) (*MsgCreateGroupResponse, error)
	// UpdateGroupMembers updates the group members with given group id and admin address.
	UpdateGroupMembers(context.Context, *MsgUpdateGroupMembers) (*MsgUpdateGroupMembersResponse, error)
	// UpdateGroupAdmin updates the group admin with given group id and previous admin address.
	UpdateGroupAdmin(context.Context, *MsgUpdateGroupAdmin) (*MsgUpdateGroupAdminResponse, error)
	// UpdateGroupMetadata updates the group metadata with given group id and admin address.
	UpdateGroupMetadata(context.Context, *MsgUpdateGroupMetadata) (*MsgUpdateGroupMetadataResponse, error)
	// CreateGroupPolicy creates a new group policy using given DecisionPolicy.
	CreateGroupPolicy(context.Context, *MsgCreateGroupPolicy) (*MsgCreateGroupPolicyResponse, error)
	// UpdateGroupPolicyAdmin updates a group policy admin.
	UpdateGroupPolicyAdmin(context.Context, *MsgUpdateGroupPolicyAdmin) (*MsgUpdateGroupPolicyAdminResponse, error)
	// UpdateGroupPolicyDecisionPolicy allows a group policy's decision policy to be updated.
	UpdateGroupPolicyDecisionPolicy(context.Context, *MsgUpdateGroupPolicyDecisionPolicy) (*MsgUpdateGroupPolicyDecisionPolicyResponse, error)
	// UpdateGroupPolicyMetadata updates a group policy metadata.
	UpdateGroupPolicyMetadata(context.Context, *MsgUpdateGroupPolicyMetadata) (*MsgUpdateGroupPolicyMetadataResponse, error)
	// SubmitProposal submits a new proposal.
	SubmitProposal(context.Context, *MsgSubmitProposal) (*MsgSubmitProposalResponse, error)
	// WithdrawProposal aborts a proposal.
	WithdrawProposal(context.Context, *MsgWithdrawProposal) (*MsgWithdrawProposalResponse, error)
	// Vote allows a voter to vote on a proposal.
	Vote(context.Context, *MsgVote) (*MsgVoteResponse, error)
	// VoteWeighted defines a method to add a weighted vote on a specific proposal.
	VoteWeighted(context.Context, *MsgVoteWeighted) (*MsgVoteWeightedResponse, error)
	// Exec executes a proposal.
	Exec(context.Context, *MsgExec) (*MsgExecResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateGroup(context.Context, *MsgCreateGroup) (*MsgCreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedMsgServer) UpdateGroupMembers(context.Context, *MsgUpdateGroupMembers) (*MsgUpdateGroupMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupMembers not implemented")
}
func (UnimplementedMsgServer) UpdateGroupAdmin(context.Context, *MsgUpdateGroupAdmin) (*MsgUpdateGroupAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupAdmin not implemented")
}
func (UnimplementedMsgServer) UpdateGroupMetadata(context.Context, *MsgUpdateGroupMetadata) (*MsgUpdateGroupMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupMetadata not implemented")
}
func (UnimplementedMsgServer) CreateGroupPolicy(context.Context, *MsgCreateGroupPolicy) (*MsgCreateGroupPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupPolicy not implemented")
}
func (UnimplementedMsgServer) UpdateGroupPolicyAdmin(context.Context, *MsgUpdateGroupPolicyAdmin) (*MsgUpdateGroupPolicyAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupPolicyAdmin not implemented")
}
func (UnimplementedMsgServer) UpdateGroupPolicyDecisionPolicy(context.Context, *MsgUpdateGroupPolicyDecisionPolicy) (*MsgUpdateGroupPolicyDecisionPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupPolicyDecisionPolicy not implemented")
}
func (UnimplementedMsgServer) UpdateGroupPolicyMetadata(context.Context, *MsgUpdateGroupPolicyMetadata) (*MsgUpdateGroupPolicyMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupPolicyMetadata not implemented")
}
func (UnimplementedMsgServer) SubmitProposal(context.Context, *MsgSubmitProposal) (*MsgSubmitProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitProposal not implemented")
}
func (UnimplementedMsgServer) WithdrawProposal(context.Context, *MsgWithdrawProposal) (*MsgWithdrawProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawProposal not implemented")
}
func (UnimplementedMsgServer) Vote(context.Context, *MsgVote) (*MsgVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (UnimplementedMsgServer) VoteWeighted(context.Context, *MsgVoteWeighted) (*MsgVoteWeightedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteWeighted not implemented")
}
func (UnimplementedMsgServer) Exec(context.Context, *MsgExec) (*MsgExecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateGroup(ctx, req.(*MsgCreateGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGroupMembers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/UpdateGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGroupMembers(ctx, req.(*MsgUpdateGroupMembers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGroupAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGroupAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGroupAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/UpdateGroupAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGroupAdmin(ctx, req.(*MsgUpdateGroupAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGroupMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGroupMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGroupMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/UpdateGroupMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGroupMetadata(ctx, req.(*MsgUpdateGroupMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateGroupPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateGroupPolicy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateGroupPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/CreateGroupPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateGroupPolicy(ctx, req.(*MsgCreateGroupPolicy))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGroupPolicyAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGroupPolicyAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGroupPolicyAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/UpdateGroupPolicyAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGroupPolicyAdmin(ctx, req.(*MsgUpdateGroupPolicyAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGroupPolicyDecisionPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGroupPolicyDecisionPolicy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGroupPolicyDecisionPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/UpdateGroupPolicyDecisionPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGroupPolicyDecisionPolicy(ctx, req.(*MsgUpdateGroupPolicyDecisionPolicy))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGroupPolicyMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGroupPolicyMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGroupPolicyMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/UpdateGroupPolicyMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGroupPolicyMetadata(ctx, req.(*MsgUpdateGroupPolicyMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/SubmitProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitProposal(ctx, req.(*MsgSubmitProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/WithdrawProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawProposal(ctx, req.(*MsgWithdrawProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Vote(ctx, req.(*MsgVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_VoteWeighted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVoteWeighted)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VoteWeighted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/VoteWeighted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VoteWeighted(ctx, req.(*MsgVoteWeighted))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.group.v1beta1.Msg/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Exec(ctx, req.(*MsgExec))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.group.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _Msg_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroupMembers",
			Handler:    _Msg_UpdateGroupMembers_Handler,
		},
		{
			MethodName: "UpdateGroupAdmin",
			Handler:    _Msg_UpdateGroupAdmin_Handler,
		},
		{
			MethodName: "UpdateGroupMetadata",
			Handler:    _Msg_UpdateGroupMetadata_Handler,
		},
		{
			MethodName: "CreateGroupPolicy",
			Handler:    _Msg_CreateGroupPolicy_Handler,
		},
		{
			MethodName: "UpdateGroupPolicyAdmin",
			Handler:    _Msg_UpdateGroupPolicyAdmin_Handler,
		},
		{
			MethodName: "UpdateGroupPolicyDecisionPolicy",
			Handler:    _Msg_UpdateGroupPolicyDecisionPolicy_Handler,
		},
		{
			MethodName: "UpdateGroupPolicyMetadata",
			Handler:    _Msg_UpdateGroupPolicyMetadata_Handler,
		},
		{
			MethodName: "SubmitProposal",
			Handler:    _Msg_SubmitProposal_Handler,
		},
		{
			MethodName: "WithdrawProposal",
			Handler:    _Msg_WithdrawProposal_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Msg_Vote_Handler,
		},
		{
			MethodName: "VoteWeighted",
			Handler:    _Msg_VoteWeighted_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _Msg_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/group/v1beta1/tx.proto",
}
