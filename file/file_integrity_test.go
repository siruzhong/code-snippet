package file

import (
	"fmt"
	"testing"
)

const commitID = "441438abd1ac652551dbe4d408dfcec8a499b8bf"
const platform = "server-linux-x64"

func TestDownloadFile(t *testing.T) {
	type args struct {
		originUrl string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy case",
			args: args{
				originUrl: fmt.Sprintf("https://update.code.visualstudio.com/commit:%s/%s/stable", commitID, platform),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DownloadFile(tt.args.originUrl); (err != nil) != tt.wantErr {
				t.Errorf("DownloadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
