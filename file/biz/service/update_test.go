package service

import (
	"context"
	"testing"
	file "file/kitex_gen/file"
)

func TestUpdate_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateService(ctx)
	// init req and assert value

	req := &file.UpdateRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
