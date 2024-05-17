package term

import (
	"fmt"
	"syscall"
	"unsafe"
)

// Some terminal ioctl stuff; we do it here with the stdlib syscall package
// because I'd like to avoid depending on x/term and x/sys, as x/sys is quite
// large (~8.5M). This is used only for the "keyscan" demo, and not any critical
// functionality.
//
// The syscall package hasn't been maintained for quite a while, but this basic
// stuff should (hopefully) work on most systems. If not: well, not a big deal
// really. Maybe we can add a build flag to prefer x/term(?)

type termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Line   uint8
	Cc     [19]uint8
	Ispeed uint32
	Ospeed uint32
}

func getTermios() (termios, error) {
	var t termios
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, 1, syscall.TCGETS, uintptr(unsafe.Pointer(&t)))
	if err > 0 {
		return t, fmt.Errorf("%s", err)
	}
	return t, nil
}

func MakeRaw() (func(), error) {
	termios, err := getTermios()
	if err != nil {
		return nil, err
	}

	old := termios
	termios.Iflag &^= syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP |
		syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON
	termios.Lflag &^= syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
	termios.Oflag &^= syscall.OPOST
	termios.Cflag &^= syscall.CSIZE | syscall.PARENB
	termios.Cflag |= syscall.CS8
	termios.Cc[syscall.VMIN] = 1
	termios.Cc[syscall.VTIME] = 0

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, 1, syscall.TCSETS, uintptr(unsafe.Pointer(&termios)))
	if errno > 0 {
		return nil, fmt.Errorf("%s", errno)
	}

	return func() {
		syscall.Syscall(syscall.SYS_IOCTL, 1, syscall.TCSETS, uintptr(unsafe.Pointer(&old)))
	}, nil
}

func IsTerminal() bool {
	_, err := getTermios()
	return err == nil
}

func Size() (int, int) {
	var size struct {
		Height, Width  uint16
		Xpixel, Ypixel uint16
	}
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, 0, syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&size)))
	if err > 0 || size.Width <= 0 || size.Height <= 0 {
		return -1, -1
	}
	return int(size.Width), int(size.Height)

}
