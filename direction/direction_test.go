package direction

import "testing"

func TestGetDirection(t *testing.T) {
	type args struct {
		sX int64
		sY int64
		dX int64
		dY int64
	}
	tests := []struct {
		name string
		args args
		want Direction
	}{
		{
			name: "Top",
			args: args{
				sX: 0,
				sY: 0,
				dX: 0,
				dY: -1,
			},
			want: Top,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDirection(tt.args.sX, tt.args.sY, tt.args.dX, tt.args.dY); got != tt.want {
				t.Errorf("GetDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
