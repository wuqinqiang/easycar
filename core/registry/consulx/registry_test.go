package consulx

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/hashicorp/consul/api"

	"github.com/wuqinqiang/easycar/core/registry"
)

func GetDefault(t *testing.T) *api.Client {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		t.Errorf("New consul client err:%v", err)
	}
	return client
}

func TestRegistry_Register(t *testing.T) {
	type fields struct {
		client *api.Client
	}
	type args struct {
		ctx       context.Context
		instances []*registry.Instance
	}

	type want struct {
		instance *registry.Instance
		err      error
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "normal",
			fields: fields{
				client: GetDefault(t),
			},
			args: args{
				ctx: context.Background(),
				instances: []*registry.Instance{
					{
						Id:      "1",
						Name:    "normalService",
						Version: "v1",
						Nodes:   []string{"grpc://127.0.0.1:2233"},
					},
				},
			},
			want: want{
				instance: &registry.Instance{
					Id:      "1",
					Name:    "normalService",
					Version: "v1",
					Nodes:   []string{"grpc://127.0.0.1:2233"},
				},
				err: nil,
			},
		},

		{
			name: "replace",
			fields: fields{
				client: GetDefault(t),
			},
			args: args{
				ctx: context.Background(),
				instances: []*registry.Instance{
					{
						Id:      "1",
						Name:    "normalService",
						Version: "v1",
						Nodes:   []string{"grpc://127.0.0.1:2233"},
					},
					{
						Id:      "1",
						Name:    "normalService",
						Version: "v2",
						Nodes:   []string{"grpc://127.0.0.1:3355"},
					},
				},
			},
			want: want{
				instance: &registry.Instance{
					Id:      "1",
					Name:    "normalService",
					Version: "v2",
					Nodes:   []string{"grpc://127.0.0.1:3355"},
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New(tt.fields.client)

			for _, instance := range tt.args.instances {
				if err := r.Register(tt.args.ctx, instance); err != tt.want.err {
					t.Errorf("Register() error = %v, wantErr %v", err, tt.want.err)
					return
				}
			}

			time.Sleep(2 * time.Second)

			watcher := newWatcher(context.Background(), r.client, tt.args.instances[0].Name)
			services, err := watcher.GetInstances()
			if err != nil {
				t.Errorf("getInstances err=%v", err)
				return
			}
			if len(services) == 0 {
				t.Errorf("service is nil")
				return
			}

			if !reflect.DeepEqual(services[0], tt.want.instance) {
				t.Errorf("server is = %v, want %v", services[0], tt.want.instance)
			}

			if err := r.DeRegister(context.Background(), tt.args.instances[0]); err != nil {
				t.Errorf("Register() error= %v", err)
			}
		})
	}
}
