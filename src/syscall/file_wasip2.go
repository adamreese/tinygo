//go:build wasip2

// This file assumes there is a libc available that runs on a real operating
// system.

package syscall

const pathMax = 1024

func Getwd() (string, error) {
	var buf [pathMax]byte
	s := getcwd(&buf[0], uint(len(buf)))
	if s == nil {
		return "", getErrno()
	}
	n := clen(buf[:])
	return string(buf[:n]), nil
}
