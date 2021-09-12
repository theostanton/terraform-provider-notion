package utils

import "testing"

func TestNormalizeId(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want string
	}{
		{
			name: "some-id-123",
			id:   "some-id-123",
			want: "someid123",
		},
		{
			name: "-some-id-123-",
			id:   "-some-id-123-",
			want: "someid123",
		},
		{
			name: "someid123",
			id:   "someid123",
			want: "someid123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeId(tt.id); got != tt.want {
				t.Errorf("NormalizeId() = %v, want %v", got, tt.want)
			}
		})
	}
}
