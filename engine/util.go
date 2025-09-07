package engine

import (
	"fmt"
	"runtime"
)

func printMemStats(r *Renderer) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Fprintf(r.buf, "\r\nAlloc = %v KiB\r\n", bToMb(m.Alloc))
	fmt.Fprintf(r.buf, "TotalAlloc = %v KiB\r\n", bToMb(m.TotalAlloc))
	fmt.Fprintf(r.buf, "Sys = %v KiB\r\n", bToMb(m.Sys))
	fmt.Fprintf(r.buf, "HeapAlloc = %v KiB\r\n", bToMb(m.HeapAlloc))
	fmt.Fprintf(r.buf, "NumGC = %v\r\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024
}
