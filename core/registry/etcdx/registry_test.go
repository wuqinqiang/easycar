package etcdx

import (
	"context"
	"reflect"
	"testing"

	"github.com/wuqinqiang/easycar/core/registry"
)

func TestRegistry_Register(t *testing.T) {

	type args struct {
		ctx      context.Context
		instance *registry.Instance
	}

	r, err := New(Conf{
		Hosts: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		t.Errorf("new etcd client=%v", err)
	}
	defer r.client.Close() //nolint:errcheck

	type want struct {
		service *registry.Instance
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "normal",
			args: args{
				ctx: context.Background(),
				instance: &registry.Instance{
					Id:      "1",
					Name:    "testService",
					Version: "v1",
					Nodes:   []string{"grpc://127.0.0.1:2233"},
				},
			},
			want: want{service: &registry.Instance{
				Id:      "1",
				Name:    "testService",
				Version: "v1",
				Nodes:   []string{"grpc://127.0.0.1:2233"},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.Register(tt.args.ctx, tt.args.instance); err != nil {
				t.Errorf("Register() error = %v, wantErr %v", err, nil)
				return
			}

			watcher, err := NewWatcher(context.Background(), r.client, "/"+tt.args.instance.Name)
			if err != nil {
				t.Errorf("NewWatcher err=%v", err)
				return
			}
			instances, err := watcher.GetInstances()
			if err != nil {
				t.Errorf("watcher getInstances err=%v", err)
				return
			}
			if len(instances) == 0 {
				t.Errorf("instances is empty")
				return
			}

			if !reflect.DeepEqual(instances[0], tt.want.service) {
				t.Errorf("get service = %v, wantErr %v", instances[0], tt.want.service)
			}

			if err := r.DeRegister(context.Background(), tt.args.instance); err != nil {
				t.Errorf("DeRegister err=%v", err)
			}
		})
	}
}
