package service

import (
	"context"
	"testing"
	user "user/kitex_gen/user"
)

func TestCheck_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckService(ctx)
	// init req and assert value

	req := &user.CheckRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
