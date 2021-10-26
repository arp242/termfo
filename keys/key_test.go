package keys

import (
	"testing"
)

func TestKey(t *testing.T) {
	t.Skip()
	tests := []struct {
		k    Key
		want string
	}{
		{'a', "<a>"},
		{'a' | Shift, "<S-a>"},
		{'a' | Ctrl | Shift, "<S-C-a>"},
		{'a' | Ctrl | Shift | Alt, "<S-C-A-a>"},
		{Tab, "<Tab>"},
		{Tab | Ctrl, "<C-Tab>"},
		{Up, "<Up>"},
		{Up | Ctrl, "<C-Up>"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			h := tt.k.String()
			if h != tt.want {
				t.Errorf("\nwant: %s\nhave: %s", tt.want, h)
			}
		})
	}
}
