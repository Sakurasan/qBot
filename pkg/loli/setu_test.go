package loli

import (
	"fmt"
	"testing"
)

func TestSetuReq(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "loli",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if list, err := SetuReq(); err != nil {
				t.Errorf("SetuReq() error = %v", err)
			} else {
				for _, v := range list {
					fmt.Println(v)
				}
			}
		})
	}
}
