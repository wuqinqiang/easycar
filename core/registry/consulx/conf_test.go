package consulx

import (
	"reflect"
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestConf_IsEmpty(t *testing.T) {
	type fields struct {
		Config  api.Config
		Default bool
	}

	type want struct {
		config api.Config
		empty  bool
	}

	defaultConf := *api.DefaultConfig()

	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "empty",
			fields: fields{
				Default: false,
			},
			want: want{
				empty: true,
			},
		},

		{
			name: "default conf",
			fields: fields{
				Config:  defaultConf,
				Default: false,
			},
			want: want{
				empty:  false,
				config: defaultConf,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conf{
				Config:  tt.fields.Config,
				Default: tt.fields.Default,
			}
			b := c.Empty()
			if b != tt.want.empty {
				t.Errorf("Empty() = %v, want %v", b, tt.want.empty)
			}

			if !b {
				if !reflect.DeepEqual(c.Conf(), &tt.want.config) {
					t.Errorf("conf = %v,want =%v", tt.fields.Config, tt.want.config)
				}
			}

		})
	}
}
