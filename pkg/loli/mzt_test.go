package loli

import (
	"fmt"
	"testing"
)

func TestMzt(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "mzt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Mzt()
			if err != nil {
				t.Errorf("Mzt() error = %v", err)
				return
			}
			fmt.Println(got)
		})
	}
}
