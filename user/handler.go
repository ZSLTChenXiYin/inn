package main

import (
	"context"
	user "user/kitex_gen/user"
	"user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Check implements the UserServiceImpl interface.
func (s *UserServiceImpl) Check(ctx context.Context, req *user.CheckRequest) (resp *user.CheckResponse, err error) {
	resp, err = service.NewCheckService(ctx).Run(req)

	return resp, err
}

// Logout implements the UserServiceImpl interface.
func (s *UserServiceImpl) Logout(ctx context.Context, req *user.LogoutRequest) (resp *user.LogoutResponse, err error) {
	resp, err = service.NewLogoutService(ctx).Run(req)

	return resp, err
}

// Detail implements the UserServiceImpl interface.
func (s *UserServiceImpl) Detail(ctx context.Context, req *user.DetailRequest) (resp *user.DetailResponse, err error) {
	resp, err = service.NewDetailService(ctx).Run(req)

	return resp, err
}

// Update implements the UserServiceImpl interface.
func (s *UserServiceImpl) Update(ctx context.Context, req *user.UpdateRequest) (resp *user.UpdateResponse, err error) {
	resp, err = service.NewUpdateService(ctx).Run(req)

	return resp, err
}
