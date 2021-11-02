package bitbucketserver

import (
	"context"
	bitbucketv1 "github.com/gfleury/go-bitbucket-v1"
	"os"
	"testing"
)

func TestClient_GetFile(t *testing.T) {
	type args struct {
		projectKey string
		repoName   string
		path       string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "get existing file - file-max.txt",
			args: args{
				projectKey: "CTO",
				repoName:   "squash-merge-poc",
				path:       "test2",
			},
			want:    "hi6",
			wantErr: false,
		},
		{
			name: "get non existing file",
			args: args{
				projectKey: "CTO",
				repoName:   "squash-merge-poc",
				path:       "non-existing-file-path",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contextWithBitbucketToken(t, context.Background())
			c := NewClient(ctx)
			got, err := c.GetFile(tt.args.projectKey, tt.args.repoName, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func contextWithBitbucketToken(t *testing.T, ctx context.Context) context.Context {
	token := os.Getenv("BITBUCKET_ACCESS_TOKEN")
	if token == "" {

		t.Errorf(`

░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▄▄███░░░░░
░░▄▄░░░░░░░░░░░░░░░░░░░░░░░░░███████░░░░
░░███▄░░░░░░░░░░░░░░░░░░░░░▄█████▀░█░░░░
░░▀█████▄▄▄▄▀▀▀▀▀▀▀▀░▄▄▄▄▄███▀▀░▀███░░░░
░░░░███▀▀░░░░░░░░░░░░░░▀▀▀███░░░░██▀░░░░
░░░░██░░░░░░▄░░░░░░░░░░░░░░░▀▀▄▄███░░░░░
░░░░▄█▄▄████▀█░█▄██▄▄░░░░░░░░░████▀░░░░░
░░░▄████████░░░██████▄▄▄▄░░░░░████░░░░░░
░░░███░█░▀██░░░▀███░█░░███▄▄░░░░▀█░░░░░░
░░░████▄███▄▄░░░███▄▄▄█████▀░░░░░██░░░░░
░░▄████▀▀░▀██▀░░░▀█████████░░░░░░██░░░░░
░░▀███░░░▄▄▀▀▀▄▄░░░░▀██████░░░░░░░█░░░░░
░░░███░░█░░░░░░░▀░░░░▀███▀░░░░░░░░█░░░░░
░░░████▄▀░░░░░░░░▀░░░████▄░░░░░░░░░█░░░░
░░░██████▄░░░░░░░░░▀▀████▀░░░░░░░░░█░░░░
░░▄█████████▀▀▀▀░░░░░░░░░░░░░░░░░░░▀█░░░
░░███████████▄▄▄▄░░░░░░░░░░░░░░░░░░░█▄░░
░░████████▀▀▀▀▀▀░░░░░░░░░░░░░░░░░░░░░█▄░
░░████████▄▄░░░░░░░░░░░░░░░░░░░░░░░░░░█░
░▄███████▄▄░░░░░░░░░░░░░░░░░░░░░░░░░░░░█
░▀▀▀▀▀▀▀▀▀█▀▀▀░░░░░░░░░░░░░░░░░░░░░░░░░█
 BITBUCKET_ACCESS_TOKEN env var not set

`)
		t.FailNow()
	}
	ctx = context.WithValue(ctx, bitbucketv1.ContextAccessToken, token)
	return ctx
}
