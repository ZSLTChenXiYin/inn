package test

import (
	"context"
	"testing"
	"user/kitex_gen/user"
	"user/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func TestUser(t *testing.T) {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		t.Fatal(err)
	}

	cli, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	register_resp, err := cli.Register(ctx, &user.RegisterRequest{
		Registers: []*user.Register{
			{
				Email:    "434223@qq.com",
				UserName: "434223",
				Password: "434223",
			},
		},
	})
	if err == nil {
		t.Log("Register:", register_resp)
	} else {
		t.Error(err)
	}

	login_resp, err := cli.Login(ctx, &user.LoginRequest{
		Logins: []*user.Login{
			{
				Email:    "434223@qq.com",
				Password: "434223",
			},
		},
	})
	if err == nil {
		t.Log("Login:", login_resp)
	} else {
		t.Error(err)
	}

	check_resp, err := cli.Check(ctx, &user.CheckRequest{
		UserName: "434223",
	})
	if err == nil {
		t.Log("Check:", check_resp)
	} else {
		t.Error(err)
	}

	update_resp, err := cli.Update(ctx, &user.UpdateRequest{
		Updates: []*user.Update{
			{
				UserId:       login_resp.UserId[0],
				NewUserName:  "1234567895",
				NewSignature: "test",
				NewPicture:   "123.jpg",
				NewPassword:  "1234567895",
			},
		},
	})
	if err == nil {
		t.Log("Update:", update_resp)
	} else {
		t.Error(err)
	}

	detail_resp, err := cli.Detail(ctx, &user.DetailRequest{
		UserId: []string{
			login_resp.UserId[0],
		},
	})
	if err == nil {
		t.Log("Detail:", detail_resp)
	} else {
		t.Error(err)
	}

	logout_resp, err := cli.Logout(ctx, &user.LogoutRequest{
		UserId: []string{
			login_resp.UserId[0],
		},
	})
	if err == nil {
		t.Log("Logout:", logout_resp)
	} else {
		t.Error(err)
	}
}
