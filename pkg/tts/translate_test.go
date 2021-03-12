package tts

import (
	"fmt"
	"regexp"
	"testing"
)

func TestTrans(t *testing.T) {
	type args struct {
		t *transType
	}
	transt := NewTransT()
	transt.Source = "zh-CN"
	transt.Target = "ja"
	transt.Text = "你的名字"
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "trans",
			args: args{t: transt},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(Trans(tt.args.t))
		})
	}
}

func TestRegexp(t *testing.T) {
	var s = `[[["あなたの名前","你的名字",null,null,11]
	]
	,null,"zh-CN",null,null,null,null,[]
	]`
	// regexp.MustCompile(`"(.*)"`)
	transreg := regexp.MustCompile(`"(.*?)"`).FindStringSubmatch(s)
	fmt.Println(transreg)
}
