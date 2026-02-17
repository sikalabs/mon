package server

import (
	"log"
	"time"

	"github.com/sikalabs/mon/pkg/run"
)

type ServerOpts struct {
	WaitTime time.Duration
}

func Server(opts ServerOpts) {
	for {
		log.Println("=== mon run ===\n")
		run.RunOrDie()
		time.Sleep(opts.WaitTime)
	}
}
