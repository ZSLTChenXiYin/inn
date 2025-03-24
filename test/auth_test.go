package test

import (
	"auth/kitex_gen/auth"
	"auth/kitex_gen/auth/authservice"
	"context"
	"testing"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func TestAuth(t *testing.T) {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		t.Fatal(err)
	}

	cli, err := authservice.NewClient("auth", client.WithResolver(r))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	/*
		token_resp, err := cli.Token(ctx, &auth.TokenRequest{
			Id: []string{
				"123",
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(token_resp)
	*/

	// access_token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMyIsInRpbWUiOjE3NDI4NDIyODF9.6j73cvkkfp3fH42fxKa9aLwAyC6c7j5YZqVUI158-Ns"
	// refresh_token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMyIsInRpbWUiOjE3NDM0MzI2ODF9.LkNErX0FGNQ_GmLeEmMfCd43c65RGhYGQLELzWHzkSg"
	// new_access_token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMyIsInRpbWUiOjE3NDI4NDIzNDd9.wqgpOKcRvlhGNwrlWwMggK8GobbWXwY_YxFXc2WKlpQ"

	validate_resp, err := cli.Validate(ctx, &auth.ValidateRequest{
		AccessToken: []string{
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMyIsInRpbWUiOjE3NDI4NDIzNDd9.wqgpOKcRvlhGNwrlWwMggK8GobbWXwY_YxFXc2WKlpQ",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(validate_resp)

	/*
		refresh_resp, err := cli.Refresh(ctx, &auth.RefreshRequest{
			RefreshToken: []string{
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMyIsInRpbWUiOjE3NDM0MzI2ODF9.LkNErX0FGNQ_GmLeEmMfCd43c65RGhYGQLELzWHzkSg"},
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(refresh_resp)
	*/

	/*
		revoke_resp, err := cli.Revoke(ctx, &auth.RevokeRequest{
			AccessToken: []string{
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMyIsInRpbWUiOjE3NDI4NDIzNDd9.wqgpOKcRvlhGNwrlWwMggK8GobbWXwY_YxFXc2WKlpQ",
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(revoke_resp)
	*/
}
