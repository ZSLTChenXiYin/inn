package service

import (
	"context"
	"testing"
	captcha "captcha/kitex_gen/captcha"
)

func TestEmailRefresh_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEmailRefreshService(ctx)
	// init req and assert value

	req := &captcha.EmailRefreshRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
