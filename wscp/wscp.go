package main

import (
	"runtime"

	"github.com/corpix/wscp/cli"
)

func init() { runtime.GOMAXPROCS(runtime.NumCPU()) }
func main() { cli.Execute() }
