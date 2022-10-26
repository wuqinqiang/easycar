package core

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

var _ Server = (*MockServer)(nil)

type MockServer struct {
	err error
}

func NewMockServer(err error) *MockServer {
	return &MockServer{err: err}
}

func (m *MockServer) Run(ctx context.Context) error {
	return m.err
}

func (m *MockServer) Stop(ctx context.Context) error {
	return nil
}

func TestRun(t *testing.T) {
	type args struct {
		opts []Option
		ctx  context.Context
	}

	err := errors.New("mock server start err")

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "context deadline exceeded",
			args: args{
				opts: []Option{WithServers(NewMockServer(nil))},
				ctx:  ctx,
			},
			want: context.DeadlineExceeded,
		},
		{
			name: "faild",
			args: args{
				opts: []Option{WithServers(NewMockServer(err))},
				ctx:  context.Background(),
			},
			want: err,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := New(tt.args.opts...)
			if got := core.Run(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
