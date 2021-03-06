package mjj

import (
	"fmt"
	"qBot/pkg/config"
	"strings"
	"testing"
)

func Test_MjjCdp(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "MjjCdp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MjjCdp()
		})
	}
}

func Test_MjjCdpMobile(t *testing.T) {
	config.Init()
	tests := []struct {
		name string
		// want    map[string]string
		// wantErr bool
	}{
		// TODO: Add test cases.
		{name: "MjjCdpMobile"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MjjCdpMobile()
			if err != nil {
				t.Errorf("cdpmjjmobile() error = %v", err)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("cdpmjjmobile() = %v, want %v", got, tt.want)
			// }
			for _, v := range got {
				fmt.Println(strings.ReplaceAll(strings.ReplaceAll(v[1], " ", ""), "\n", ""), v[0])
			}
		})
	}
}

func TestLocalmobile(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "Localmobile"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Localmobile()
		})
	}
}
