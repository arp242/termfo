package termfo

import (
	"os/exec"
	"strconv"
	"testing"
)

func TestParamsTput(t *testing.T) {
	if _, err := exec.LookPath("tput"); err != nil {
		t.Skipf("needs tput in PATH: %s", err)
	}

	tests := []struct {
		term, capName string
		in            string
		params        []int
	}{
		{"xterm", "ech", "\x1b[%p1%dX", []int{5}},
		// TODO: some more tests.
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			cmd := []string{"-T" + tt.term, tt.capName}
			for _, a := range tt.params {
				cmd = append(cmd, strconv.Itoa(a))
			}
			o, err := exec.Command("tput", cmd...).CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}
			want := string(o)
			have := replaceParams(tt.in, tt.params...)
			if have != want {
				t.Errorf("\nin:   %q\nhave: %q\nwant: %q\n%s\n",
					tt.in, have, want, printParams(tt.in, tt.params...))
			}
		})
	}
}

func TestParams(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		// xterm-256color
		{"\x1b[%p1%dX", "\x1b[666X"},                                       // ech
		{"\x1b[%i%p1%dG", "\x1b[667G"},                                     // hpa
		{"\x1b[%p1%d q", "\x1b[666 q"},                                     // Ss
		{"\x1b]12;%p1%s\a", "\x1b]12;666\a"},                               // Cs
		{"\x1b[%i%p1%dd", "\x1b[667d"},                                     // vpa
		{"\x1b]52;%p1%s;%p2%s\a", "\x1b]52;666;777\a"},                     // Ms
		{"\x1b[%i%p1%d;%p2%dr", "\x1b[667;778r"},                           // csr
		{"\x1b[%i%p1%d;%p2%dH", "\x1b[667;778H"},                           // cup
		{"%p1%c\x1b[%p2%{1}%-%db", "ʚ\x1b[776b"},                           // rep
		{"\x1b[?69h\x1b[%i%p1%d;%p2%ds", "\x1b[?69h\x1b[667;778s"},         // smglr
		{"\x1b[%i%d;%dR", "\x1b[1;1R"},                                     // u6
		{"\x1b[?1006;1000%?%p1%{1}%=%th%el%;", "\x1b[?1006;1000l"},         // XM
		{"\x1b[<%i%p3%d;%p1%d;%p2%d;%?%p4%tM%em%;", "\x1b[<888;667;778;M"}, // xm

		{"\x1b]4;%p1%d;rgb:%p2%{255}%*%{1000}%/%2.2X/%p3%{255}%*%{1000}%/%2.2X/%p4%{255}%*%{1000}%/%2.2X\x1b\\",
			"\x1b]4;666;rgb:C6/E2/FE\x1b\\"}, // initc

		{"\x1b[%?%p1%{8}%<%t4%p1%d%e%p1%{16}%<%t10%p1%{8}%-%d%e48;5;%p1%d%;m",
			"\x1b[48;5;666m"}, // setab
		{"\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;m",
			"\x1b[38;5;666m"}, // setaf
		{"%?%p9%t\x1b(0%e\x1b(B%;\x1b[0%?%p6%t;1%;%?%p5%t;2%;%?%p2%t;4%;%?%p1%p3%|%t;7%;%?%p4%t;5%;%?%p7%t;8%;m",
			"\x1b(B\x1b[0;4;7;5m"}, // sgr

		// TODO: what is %[? Is this correct?
		{"\x1b[?%[;0123456789]c", "\x1b[?%[;0123456789]c"}, // u8

		// tmux terminfo
		{"\x1b]52;%p1%s;%p2%s\a", "\x1b]52;666;777\a"}, // Ms
		{"\x1b(%p1%c", "\x1b(ʚ"},                       // S0
		{"\x1b[4:%p1%dm", "\x1b[4:666m"},               // Smulx
		{"\x1b[%i%p1%d;%p2%dr", "\x1b[667;778r"},       // csr
		{"\x1b[%i%p1%d;%p2%dH", "\x1b[667;778H"},       // cup
		{"\x1b[%i%p1%dd", "\x1b[667d"},                 // vpa

		{"\x1b[0%?%p6%t;1%;%?%p2%t;4%;%?%p1%p3%|%t;7%;%?%p4%t;5%;%?%p5%t;2%;%?%p7%t;8%;m%?%p9%t\x0e%e\x0f%;",
			"\x1b[0;4;7;5m\x0f"}, // sgr

		// alacritty
		{"\x1b[%i%p1%d;%p2%dr", "\x1b[667;778r"},       // change_scroll_region
		{"\x1b[%i%p1%d;%p2%dH", "\x1b[667;778H"},       // cursor_address
		{"\x1bP=%p1%ds\x1b\\", "\x1bP=666s\x1b\\"},     // Sync
		{"\x1b]12;%p1%s^G", "\x1b]12;666^G"},           // Cs
		{"\x1b[4\072%p1%dm", "\x1b[4\072666m"},         // Smulx
		{"\x1b]52;%p1%s;%p2%s^G", "\x1b]52;666;777^G"}, // Ms

		{"\x1b[%?%p1%{8}%<%t4%p1%d%e48\0722\072\072%p1%{65536}%/%d\072%p1%{256}%/%{255}%&%d\072%p1%{255}%&%d%;m",
			"\x1b[48:2::0:2:154m"}, // set_a_background

		{"\x1b[%?%p1%{8}%<%t3%p1%d%e38\0722\072\072%p1%{65536}%/%d\072%p1%{256}%/%{255}%&%d\072%p1%{255}%&%d%;m",
			"\x1b[38:2::0:2:154m"}, // set_a_foreground
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			args := []int{666, 777, 888, 999}
			have := replaceParams(tt.in, args...)
			if have != tt.want {
				t.Errorf("\nin:   %q\nhave: %q\nwant: %q\n%s\n",
					tt.in, have, tt.want, printParams(tt.in, args...))
			}
		})
	}
}
