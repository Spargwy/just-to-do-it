package service

import "testing"

func Test_buildWhereConditionFromParams(t *testing.T) {
	type args struct {
		filterParams map[string][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				filterParams: map[string][]string{
					"title":    {"cook diner"},
					"archived": {"false"},
				},
			},
			want: "title = 'cook diner' and archived = 'false'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildWhereConditionFromParams(tt.args.filterParams); got != tt.want {
				t.Errorf("buildWhereConditionFromParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
