package tts

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSpeak(t *testing.T) {
	type args struct {
		t Config
	}
	tests := []struct {
		name    string
		args    args
		want    *Speech
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "tts",
			args: args{t: Config{Language: "ja", Speak: "支付宝"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Speak(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Speak() error = %v", err)
				return
			}
			file, err := os.OpenFile("tts.mp3", os.O_CREATE|os.O_RDWR, os.ModePerm)
			defer file.Close()
			if err != nil {
				log.Println(err)
			}
			file.Write(got.Bytes())
			var a bool
			fmt.Println(a)

		})
	}
}
