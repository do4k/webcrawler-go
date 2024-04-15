package strings

import (
	"testing"
)

func TestStartsWithAny(t *testing.T) {
	tests := []struct {
		name  string
		array []string
		s     string
		want  bool
	}{
		{
			name:  "String starts with one of the prefixes",
			array: []string{"pre", "suf", "mid"},
			s:     "prefix",
			want:  true,
		},
		{
			name:  "String does not start with any of the prefixes",
			array: []string{"pre", "suf", "mid"},
			s:     "postfix",
			want:  false,
		},
		{
			name:  "Empty string",
			array: []string{"pre", "suf", "mid"},
			s:     "",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartsWithAny(tt.array, tt.s); got != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
