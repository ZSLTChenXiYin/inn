package service

import (
	"context"
	"testing"
	auth "auth/kitex_gen/auth"
)

func TestRevoke_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRevokeService(ctx)
	// init req and assert value

	req := &auth.RevokeRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
