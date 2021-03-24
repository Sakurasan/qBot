package appso

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestXianMian(t *testing.T) {
	tests := []struct {
		name string
		want []It
	}{
		// TODO: Add test cases.
		{
			name: "RSS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := XianMian()
			if err != nil {
				t.Errorf("XianMian() error = %v", err)
				return
			}
			for _, v := range got {
				desp := strings.ReplaceAll(v.Description, "<br>\n", "")
				desp = strings.ReplaceAll(desp, " ", "")

				desp = strings.ReplaceAll(desp, regexp.MustCompile("<imgsrc(.*?)>").FindString(desp), "")
				fmt.Println(v.Title, desp, v.Guid.Text+"\n")
			}
		})
	}
}
