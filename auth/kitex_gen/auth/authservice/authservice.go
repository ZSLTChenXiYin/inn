// Code generated by Kitex v0.9.1. DO NOT EDIT.

package authservice

import (
	auth "auth/kitex_gen/auth"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Token": kitex.NewMethodInfo(
		tokenHandler,
		newTokenArgs,
		newTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Validate": kitex.NewMethodInfo(
		validateHandler,
		newValidateArgs,
		newValidateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Refresh": kitex.NewMethodInfo(
		refreshHandler,
		newRefreshArgs,
		newRefreshResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Revoke": kitex.NewMethodInfo(
		revokeHandler,
		newRevokeArgs,
		newRevokeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	authServiceServiceInfo                = NewServiceInfo()
	authServiceServiceInfoForClient       = NewServiceInfoForClient()
	authServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return authServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return authServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return authServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "AuthService"
	handlerType := (*auth.AuthService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "auth",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func tokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.TokenRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).Token(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *TokenArgs:
		success, err := handler.(auth.AuthService).Token(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*TokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newTokenArgs() interface{} {
	return &TokenArgs{}
}

func newTokenResult() interface{} {
	return &TokenResult{}
}

type TokenArgs struct {
	Req *auth.TokenRequest
}

func (p *TokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.TokenRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *TokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *TokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *TokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *TokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.TokenRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var TokenArgs_Req_DEFAULT *auth.TokenRequest

func (p *TokenArgs) GetReq() *auth.TokenRequest {
	if !p.IsSetReq() {
		return TokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *TokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *TokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type TokenResult struct {
	Success *auth.TokenResponse
}

var TokenResult_Success_DEFAULT *auth.TokenResponse

func (p *TokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.TokenResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *TokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *TokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *TokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *TokenResult) Unmarshal(in []byte) error {
	msg := new(auth.TokenResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *TokenResult) GetSuccess() *auth.TokenResponse {
	if !p.IsSetSuccess() {
		return TokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *TokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.TokenResponse)
}

func (p *TokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TokenResult) GetResult() interface{} {
	return p.Success
}

func validateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.ValidateRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).Validate(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ValidateArgs:
		success, err := handler.(auth.AuthService).Validate(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ValidateResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newValidateArgs() interface{} {
	return &ValidateArgs{}
}

func newValidateResult() interface{} {
	return &ValidateResult{}
}

type ValidateArgs struct {
	Req *auth.ValidateRequest
}

func (p *ValidateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.ValidateRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ValidateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ValidateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ValidateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ValidateArgs) Unmarshal(in []byte) error {
	msg := new(auth.ValidateRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ValidateArgs_Req_DEFAULT *auth.ValidateRequest

func (p *ValidateArgs) GetReq() *auth.ValidateRequest {
	if !p.IsSetReq() {
		return ValidateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ValidateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ValidateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ValidateResult struct {
	Success *auth.ValidateResponse
}

var ValidateResult_Success_DEFAULT *auth.ValidateResponse

func (p *ValidateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.ValidateResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ValidateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ValidateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ValidateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ValidateResult) Unmarshal(in []byte) error {
	msg := new(auth.ValidateResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ValidateResult) GetSuccess() *auth.ValidateResponse {
	if !p.IsSetSuccess() {
		return ValidateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ValidateResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.ValidateResponse)
}

func (p *ValidateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ValidateResult) GetResult() interface{} {
	return p.Success
}

func refreshHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.RefreshRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).Refresh(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RefreshArgs:
		success, err := handler.(auth.AuthService).Refresh(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RefreshResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRefreshArgs() interface{} {
	return &RefreshArgs{}
}

func newRefreshResult() interface{} {
	return &RefreshResult{}
}

type RefreshArgs struct {
	Req *auth.RefreshRequest
}

func (p *RefreshArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.RefreshRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RefreshArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RefreshArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RefreshArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RefreshArgs) Unmarshal(in []byte) error {
	msg := new(auth.RefreshRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RefreshArgs_Req_DEFAULT *auth.RefreshRequest

func (p *RefreshArgs) GetReq() *auth.RefreshRequest {
	if !p.IsSetReq() {
		return RefreshArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RefreshArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RefreshArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RefreshResult struct {
	Success *auth.RefreshResponse
}

var RefreshResult_Success_DEFAULT *auth.RefreshResponse

func (p *RefreshResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.RefreshResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RefreshResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RefreshResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RefreshResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RefreshResult) Unmarshal(in []byte) error {
	msg := new(auth.RefreshResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RefreshResult) GetSuccess() *auth.RefreshResponse {
	if !p.IsSetSuccess() {
		return RefreshResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RefreshResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.RefreshResponse)
}

func (p *RefreshResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RefreshResult) GetResult() interface{} {
	return p.Success
}

func revokeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.RevokeRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).Revoke(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RevokeArgs:
		success, err := handler.(auth.AuthService).Revoke(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RevokeResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRevokeArgs() interface{} {
	return &RevokeArgs{}
}

func newRevokeResult() interface{} {
	return &RevokeResult{}
}

type RevokeArgs struct {
	Req *auth.RevokeRequest
}

func (p *RevokeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.RevokeRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RevokeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RevokeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RevokeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RevokeArgs) Unmarshal(in []byte) error {
	msg := new(auth.RevokeRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RevokeArgs_Req_DEFAULT *auth.RevokeRequest

func (p *RevokeArgs) GetReq() *auth.RevokeRequest {
	if !p.IsSetReq() {
		return RevokeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RevokeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RevokeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RevokeResult struct {
	Success *auth.RevokeResponse
}

var RevokeResult_Success_DEFAULT *auth.RevokeResponse

func (p *RevokeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.RevokeResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RevokeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RevokeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RevokeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RevokeResult) Unmarshal(in []byte) error {
	msg := new(auth.RevokeResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RevokeResult) GetSuccess() *auth.RevokeResponse {
	if !p.IsSetSuccess() {
		return RevokeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RevokeResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.RevokeResponse)
}

func (p *RevokeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RevokeResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Token(ctx context.Context, Req *auth.TokenRequest) (r *auth.TokenResponse, err error) {
	var _args TokenArgs
	_args.Req = Req
	var _result TokenResult
	if err = p.c.Call(ctx, "Token", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Validate(ctx context.Context, Req *auth.ValidateRequest) (r *auth.ValidateResponse, err error) {
	var _args ValidateArgs
	_args.Req = Req
	var _result ValidateResult
	if err = p.c.Call(ctx, "Validate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Refresh(ctx context.Context, Req *auth.RefreshRequest) (r *auth.RefreshResponse, err error) {
	var _args RefreshArgs
	_args.Req = Req
	var _result RefreshResult
	if err = p.c.Call(ctx, "Refresh", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Revoke(ctx context.Context, Req *auth.RevokeRequest) (r *auth.RevokeResponse, err error) {
	var _args RevokeArgs
	_args.Req = Req
	var _result RevokeResult
	if err = p.c.Call(ctx, "Revoke", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
