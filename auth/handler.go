package main

import (
	"auth/biz/service"
	auth "auth/kitex_gen/auth"
	"context"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// Token implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Token(ctx context.Context, req *auth.TokenRequest) (resp *auth.TokenResponse, err error) {
	resp, err = service.NewTokenService(ctx).Run(req)

	return resp, err
}

// Validate implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Validate(ctx context.Context, req *auth.ValidateRequest) (resp *auth.ValidateResponse, err error) {
	resp, err = service.NewValidateService(ctx).Run(req)

	return resp, err
}

// Refresh implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Refresh(ctx context.Context, req *auth.RefreshRequest) (resp *auth.RefreshResponse, err error) {
	resp, err = service.NewRefreshService(ctx).Run(req)

	return resp, err
}

// Revoke implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Revoke(ctx context.Context, req *auth.RevokeRequest) (resp *auth.RevokeResponse, err error) {
	resp, err = service.NewRevokeService(ctx).Run(req)

	return resp, err
}
