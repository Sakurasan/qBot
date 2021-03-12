package tb

import (
	"testing"
)

func TestTao(t *testing.T) {
	tests := []struct {
		name string
		want map[string]string
	}{
		// TODO: Add test cases.
		{
			name: "TB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Tao()
			if err != nil {
				t.Errorf("Tao() error = %v", err)
				return
			}
			t.Log(got)
		})
	}
}
