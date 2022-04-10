package handler

import "testing"

func TestRandomURL(t *testing.T) {

	tests := []struct {
		name string
		size int
		want string
	}{
		{
			name: "passing with",
			size: 2,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomURL(tt.size); got == tt.want {
				t.Errorf("RandomURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
