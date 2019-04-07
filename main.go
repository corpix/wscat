package main

import (
	"runtime"

	"github.com/corpix/wscat/cli"
)

func init() { runtime.GOMAXPROCS(runtime.NumCPU()) }
func main() { cli.Execute() }
