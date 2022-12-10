package client

import (
	"reflect"
	"testing"

	"github.com/wuqinqiang/easycar/core/consts"
)

func TestNewBranch(t *testing.T) {
	type args struct {
		uri    string
		action consts.BranchAction
	}
	tests := []struct {
		name string
		args args
		want *Branch
	}{
		{
			name: "new http branch",
			args: args{
				uri:    "http://127.0.0.1",
				action: consts.Try,
			},
			want: &Branch{
				uri:      "http://127.0.0.1",
				data:     nil,
				header:   nil,
				action:   consts.Try,
				level:    1,
				timeout:  3, //default
				protocol: "http",
			},
		},
		{
			name: "new grpc branch",
			args: args{
				uri:    "127.0.0.1",
				action: consts.Try,
			},
			want: &Branch{
				uri:      "grpc://127.0.0.1",
				data:     nil,
				header:   nil,
				action:   consts.Try,
				level:    1,
				timeout:  3, //default
				protocol: "grpc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBranch(tt.args.uri, tt.args.action); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_protocol(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 Protocol
	}{
		{
			name: "http",
			args: args{
				uri: "http://127.0.0.1",
			},
			want:  "http://127.0.0.1",
			want1: HTTP,
		},
		{
			name: "grpc normal",
			args: args{
				uri: "grpc://127.0.0.1",
			},
			want:  "grpc://127.0.0.1",
			want1: GRPC,
		},
		{
			name: "grpc normal2",
			args: args{
				uri: "127.0.0.1",
			},
			want:  "grpc://127.0.0.1",
			want1: GRPC,
		},
		{
			name: "undefined",
			args: args{
				uri: "other://127.0.0.1",
			},
			want:  "other://127.0.0.1",
			want1: Undefined,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := protocol(tt.args.uri)
			if got != tt.want {
				t.Errorf("protocol() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("protocol() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
