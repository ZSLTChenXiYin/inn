package test

import (
	"context"
	"file/kitex_gen/file"
	"file/kitex_gen/file/fileservice"
	"testing"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func TestFile(t *testing.T) {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		t.Fatal(err)
	}

	cli, err := fileservice.NewClient("file", client.WithResolver(r))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	/*
		upload_resp, err := cli.Upload(ctx, &file.UploadRequest{
			Lists: []*file.UploadList{
				{
					FileName:     "百鬼COSPLAY_妖刀姬_完整版.zip",
					FilePath:     "D:/Picture/百鬼COSPLAY_妖刀姬_完整版.zip",
					OwnerId:      "123456",
					AccessLevel:  2,
					FilePassword: "123456",
				},
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(upload_resp)
	*/

	/*
		download_resp, err := cli.Download(ctx, &file.DownloadRequest{
			Lists: []*file.DownloadList{
				{
					FileId:       "7acfe5a0-6165-418b-a00b-a850d808b252",
					UserId:       "1234567",
					FilePassword: "123456",
				},
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(download_resp)
	*/

	/*
		update_resp, err := cli.Update(ctx, &file.UpdateRequest{
			Files: []*file.UpdateFile{
				{
					FileId:       "7acfe5a0-6165-418b-a00b-a850d808b252",
					UserId:       "123456",
					FilePassword: "123456",
					NewFileName:  "百鬼COSPLAY_妖刀姬_完整版_new.zip",
				},
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(update_resp)
	*/

	/*
		detail_resp, err := cli.Detail(ctx, &file.DetailRequest{
			Files: []*file.DetailFile{
				{
					FileId: "7acfe5a0-6165-418b-a00b-a850d808b252",
					UserId: "1234567",
				},
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(detail_resp)
	*/

	/*
		check_resp, err := cli.Check(ctx, &file.CheckRequest{
			OwnerId: "123456",
			UserId:  "1234567",
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(check_resp)
	*/

	delete_resp, err := cli.Delete(ctx, &file.DeleteRequest{
		Files: []*file.DeleteFile{
			{
				FileId: "7acfe5a0-6165-418b-a00b-a850d808b252",
				UserId: "123456",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(delete_resp)
}
