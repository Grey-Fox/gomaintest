// +build testrunmain

package main

import (
	"syscall"
	"testing"
	"time"
)

func TestRunMain(t *testing.T) {
	go func() {
		time.Sleep(2 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()
	main()
}
