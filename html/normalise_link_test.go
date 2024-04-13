package html

import (
	"testing"
)

func TestNormaliseLink(t *testing.T) {
	tests := []struct {
		name        string
		link        string
		baseAddress string
		want        string
	}{
		{
			name:        "Test with empty link",
			link:        "",
			baseAddress: "http://example.com",
			want:        "",
		},
		{
			name:        "Test relative links",
			link:        "/test",
			baseAddress: "http://example.com",
			want:        "http://example.com/test",
		},
		{
			name:        "Test with link containing query params",
			link:        "http://example.com/test?param=value",
			baseAddress: "",
			want:        "http://example.com/test",
		},
		{
			name:        "Test with link containing anchor",
			link:        "http://example.com/test#anchor",
			baseAddress: "",
			want:        "http://example.com/test",
		},
		{
			name:        "Test with link ending with slash",
			link:        "http://example.com/test/",
			baseAddress: "",
			want:        "http://example.com/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormaliseLink(tt.link, tt.baseAddress); got != tt.want {
				t.Errorf("NormaliseLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
