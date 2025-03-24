package service

import (
	"context"
	"testing"
	auth "auth/kitex_gen/auth"
)

func TestToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewTokenService(ctx)
	// init req and assert value

	req := &auth.TokenRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
