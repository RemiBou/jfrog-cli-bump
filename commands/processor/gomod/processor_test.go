package gomod

import (
	"context"
	"github.com/jfrog/jfrog-cli-bump/commands/processor/model"
	"testing"
)

func Test_processor_Process(t *testing.T) {
	type args struct {
		ctx              context.Context
		versionToUpgrade model.PackageVersion
		inVersionFile    string
		inChecksumFile   string
	}
	tests := []struct {
		name             string
		args             args
		wantVersionFile  string
		wantChecksumFile string
		wantErr          bool
	}{
		{
			name:             "should bump version",
			args:             args{
				versionToUpgrade: model.PackageVersion{ Name: "jfrog.com/test/abc", Version: "v1.2.3" },
				inVersionFile:    "module github.com/jfrog/jfrog-cli-bump\n\ngo 1.17\n\nrequire (\n\tjfrog.com/test/abc v1.0.0\n\tgithub.com/jfrog/jfrog-cli-core/v2 v2.3.0\n\tgithub.com/jfrog/jfrog-client-go v1.4.0\n\tgithub.com/stretchr/testify v1.7.0\n)\n",
			},
			wantVersionFile:  "module github.com/jfrog/jfrog-cli-bump\n\ngo 1.17\n\nrequire (\n\tjfrog.com/test/abc v1.2.3\n\tgithub.com/jfrog/jfrog-cli-core/v2 v2.3.0\n\tgithub.com/jfrog/jfrog-client-go v1.4.0\n\tgithub.com/stretchr/testify v1.7.0\n)\n",
		},
		{
			name:             "should bump multiple versions",
			args:             args{
				versionToUpgrade: model.PackageVersion{ Name: "jfrog.com/test/abc", Version: "v1.2.3" },
				inVersionFile:    "module github.com/jfrog/jfrog-cli-bump\n\ngo 1.17\n\nrequire (\n\tjfrog.com/test/abc/v2 v1.0.0\n\tjfrog.com/test/abc/v4 v2.3.0\n\tgithub.com/jfrog/jfrog-client-go v1.4.0\n\tgithub.com/stretchr/testify v1.7.0\n)\n",
			},
			wantVersionFile:  "module github.com/jfrog/jfrog-cli-bump\n\ngo 1.17\n\nrequire (\n\tjfrog.com/test/abc/v2 v1.2.3\n\tjfrog.com/test/abc/v4 v1.2.3\n\tgithub.com/jfrog/jfrog-client-go v1.4.0\n\tgithub.com/stretchr/testify v1.7.0\n)\n",
		},
		{
			name:             "should not bump any version",
			args:             args{
				versionToUpgrade: model.PackageVersion{ Name: "jfrog.com/test/foo", Version: "v1.2.3" },
				inVersionFile:    "module github.com/jfrog/jfrog-cli-bump\n\ngo 1.17\n\nrequire (\n\tjfrog.com/test/abc/v2 v1.0.0\n\tjfrog.com/test/abc/v4 v2.3.0\n\tgithub.com/jfrog/jfrog-client-go v1.4.0\n\tgithub.com/stretchr/testify v1.7.0\n)\n",
			},
			wantVersionFile:  "module github.com/jfrog/jfrog-cli-bump\n\ngo 1.17\n\nrequire (\n\tjfrog.com/test/abc/v2 v1.0.0\n\tjfrog.com/test/abc/v4 v2.3.0\n\tgithub.com/jfrog/jfrog-client-go v1.4.0\n\tgithub.com/stretchr/testify v1.7.0\n)\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := processor{}
			gotVersionFile, gotChecksumFile, err := pr.Process(context.Background(), tt.args.versionToUpgrade, tt.args.inVersionFile, tt.args.inChecksumFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotVersionFile != tt.wantVersionFile {
				t.Errorf("Process() gotVersionFile = %v, want %v", gotVersionFile, tt.wantVersionFile)
			}
			if gotChecksumFile != tt.wantChecksumFile {
				t.Errorf("Process() gotChecksumFile = %v, want %v", gotChecksumFile, tt.wantChecksumFile)
			}
		})
	}
}
