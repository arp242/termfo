package termfo

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func BenchmarkNew(b *testing.B) {
	t := os.Getenv("TERM")
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		New(t)
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		term    string
		intSize int
	}{
		{"xterm", 2},
		{"xterm-256color", 4},
	}

	for _, tt := range tests {
		t.Run(tt.term, func(t *testing.T) {
			ti, err := New(tt.term)
			if err != nil {
				t.Fatal(err)
			}

			if ti.Name != tt.term {
				t.Errorf("wrong name: %q", ti.Name)
			}
			if ti.IntSize != tt.intSize {
				t.Errorf("intsize: %d", ti.IntSize)
			}

			// Just a basic sanity check.
			if len(ti.Bools) < 10 || len(ti.Numbers) < 5 || len(ti.Strings) < 200 || len(ti.Extended) < 50 {
				t.Errorf("bools: %d  nums: %d  strs: %d  ext: %d", len(ti.Bools), len(ti.Numbers), len(ti.Strings), len(ti.Extended))
			}
		})
	}
}

func TestKeys(t *testing.T) {
	run := func(ti *Terminfo, seq string) []string {
		var collect []string
		ch := ti.FindKeys(strings.NewReader(seq))
		for e := <-ch; ; e = <-ch {
			if e.Err != nil {
				if e.Err == io.EOF {
					break
				}
				t.Fatal(e.Err)
			}
			collect = append(collect, e.Key.String())
		}
		return collect
	}

	{
		ti, err := New("xterm")
		if err != nil {
			t.Fatal(err)
		}
		f := run(ti, "\x1bOP\x1b[1;2P\x1bOA")
		if fmt.Sprintf("%s", f) != "[<F1> <S-F1> <Up>]" {
			t.Error(f)
		}
	}

	{
		ti, err := New("vi200")
		if err != nil {
			t.Fatal(err)
		}
		f := run(ti, "\x1b?q\x1bA")
		if fmt.Sprintf("%s", f) != "[<F1> <Up>]" {
			t.Error(f)
		}
	}
}
