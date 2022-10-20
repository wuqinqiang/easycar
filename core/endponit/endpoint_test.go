package endponit

import "testing"

func TestGetHostByEndpoint(t *testing.T) {
	type args struct {
		endpoints []string
		scheme    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "nil endpoints",
			args:    args{scheme: "nil"},
			want:    "",
			wantErr: false,
		},
		{
			name: "http scheme",
			args: args{
				endpoints: []string{"http://127.0.0.1:8080/xxxx", "http://127.0.0.1:8089/yyy"},
				scheme:    "http",
			},
			want:    "127.0.0.1:8080",
			wantErr: false,
		},
		{
			name: "grpc scheme",
			args: args{
				endpoints: []string{"http://127.0.0.1:8080/xxxx", "grpc://127.0.0.1:8089/yyy"},
				scheme:    "grpc",
			},
			want:    "127.0.0.1:8089",
			wantErr: false,
		},
		{
			name: "endpoints err",
			args: args{
				endpoints: []string{":8080"},
				scheme:    "grpc",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHostByEndpoint(tt.args.endpoints, tt.args.scheme)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostByEndpoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetHostByEndpoint() got = %v, want %v", got, tt.want)
			}
		})
	}
}
