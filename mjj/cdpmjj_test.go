package mjj

import (
	"fmt"
	"testing"
)

func Test_cdpmjj(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "cdpmjj",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cdpmjj()
		})
	}
}

func Test_cdpmjjmobile(t *testing.T) {
	tests := []struct {
		name string
		// want    map[string]string
		// wantErr bool
	}{
		// TODO: Add test cases.
		{name: "cdpmjjmobile"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cdpmjjmobile()
			if err != nil {
				t.Errorf("cdpmjjmobile() error = %v", err)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("cdpmjjmobile() = %v, want %v", got, tt.want)
			// }
			fmt.Printf("%+v", got)
		})
	}
}
