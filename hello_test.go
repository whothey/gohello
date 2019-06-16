package hello

import (
	"testing"
)

func TestWorld(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "returns hello world", want: "Hello world!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := World(); got != tt.want {
				t.Errorf("World() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsg(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "returns a custom msg", args: args{"my message"}, want: "Hello my message"},
		{name: "returns no msg if empty string", args: args{""}, want: "Hello "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Msg(tt.args.msg); got != tt.want {
				t.Errorf("Msg() = %v, want %v", got, tt.want)
			}
		})
	}
}
