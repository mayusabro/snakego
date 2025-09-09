package engine

import (
	"fmt"
	"runtime"
)

func removeElement[T any](s []T, i int) []T {
	if i >= len(s) || i < 0 {
		return s
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func printMemStats(r *Renderer) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Fprintf(r.buf, "\r\nAlloc = %v KiB\r\n", bToMb(m.Alloc))
	fmt.Fprintf(r.buf, "NumGC = %v\r\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024
}
