package qqmusic

import "testing"

func Test_get_album_count(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "t1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get_album_count(); got != -1 {
				t.Logf("get_album_count() = %v", got)
			} else {
				t.Errorf("get_album_count() = %v", got)
			}
		})
	}
}
