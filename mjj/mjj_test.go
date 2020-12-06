package mjj

import (
	"testing"
)

func Test_initlist(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initlist()
		})
	}
}

func Test_browser(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			browser()
		})
	}
}

func Test_browser2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "browser2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			browser2()
		})
	}
}
