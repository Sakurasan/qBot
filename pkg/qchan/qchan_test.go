package qchan

import (
	"qBot/tests"
	"testing"
)

func TestSend(t *testing.T) {
	type args struct {
		msg string
	}
	ts := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{msg: "刺客567"},
		},
		{
			name: "2",
			args: args{msg: "梅花13"},
		},
	}
	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			Send(tt.args.msg)
			SendGroup(tt.args.msg, tests.Qchan.GroupID)
		})
	}
}
