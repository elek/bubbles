package ui

import (
	"reflect"
	"testing"
)

func TestResolve(t *testing.T) {
	type args struct {
		req []SizeDefinition
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "fixed and flexi",
			args: args{
				req: []SizeDefinition{
					fixed{
						Value: 10,
					},
					remaining{},
				},
				max: 200,
			},
			want: []int{10, 190},
		},
		{
			name: "flexi and fixed",
			args: args{
				req: []SizeDefinition{
					remaining{},
					fixed{
						Value: 10,
					},
				},
				max: 200,
			},
			want: []int{190, 10},
		},
		{
			name: "flexi fixed flexi",
			args: args{
				req: []SizeDefinition{
					remaining{},
					fixed{
						Value: 10,
					},
					remaining{},
				},
				max: 200,
			},
			want: []int{95, 10, 95},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Resolve(tt.args.req, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
