package client

import (
	"testing"
)

func defaultGroups(opts ...GroupOpt) []*Group {
	uri := "127.0.0.1"

	return []*Group{
		NewTccGroup(uri, uri, uri, opts...),
		NewSagaGroup(uri, uri, opts...),
	}
}

func getGroups() []*Group {
	uri := "127.0.0.1"

	return []*Group{
		NewTccGroup(uri, uri, uri).SetLevel(2),
		NewSagaGroup(uri, uri).SetLevel(3),
	}
}

func TestManger_AddGroup(t *testing.T) {
	type args struct {
		skip   bool
		groups []*Group
	}
	tests := []struct {
		name string
		args args
		want []*Group
	}{
		{
			name: "add incr false",
			args: args{
				skip:   false,
				groups: defaultGroups(),
			},
			want: defaultGroups(),
		},
		{
			name: "add incr true",
			args: args{
				skip:   true,
				groups: defaultGroups(),
			},
			want: getGroups(),
		},
		{
			name: "level fixed",
			args: args{
				skip:   true,
				groups: defaultGroups(LevelFixed()),
			},
			want: defaultGroups(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewManger()
			for i := range tt.args.groups {
				if tt.args.skip {
					got.AddNextWaitGroups(tt.args.groups[i])
					continue
				}
				got.addGroups(tt.args.groups[i])
			}
			if len(got.groups) != len(tt.want) {
				t.Errorf("got groups len=%v,want %v", len(got.groups), len(tt.want))
			}
			for index, group := range got.groups {
				gotBranches := group.branches
				wantBranches := tt.want[index].branches
				if len(gotBranches) != len(wantBranches) {
					t.Errorf("got branchers len=%v,want %v", len(group.branches), len(wantBranches))
				}
				for index, gotBranch := range gotBranches {
					if gotBranch.level != wantBranches[index].level {
						t.Errorf("got branch level =%v,want=%v", gotBranch.level,
							wantBranches[index].level)
					}
				}

			}
		})
	}
}
