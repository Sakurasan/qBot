package bilibili

import (
	"fmt"
	"qBot/tests"
	"testing"
)

func TestUpCoinVideo(t *testing.T) {
	type args struct {
		vmid string
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "bilibili Coin video",
			args: args{vmid: tests.Bili.UpMid},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpCoinVideo(tt.args.vmid)
			if err != nil {
				t.Errorf("UpCoinVideo() error = %v", err)
				return
			}
			for _, v := range got.Data {
				fmt.Printf("%.15s - %10s \t %s %s \n", v.Owner.Name, v.Title, fmt.Sprintf("https://www.bilibili.com/video/av%d", v.Aid), fmt.Sprintf("https://www.bilibili.com/video/%s", v.Bvid))
			}
		})
	}
}
