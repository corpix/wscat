package cli

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/cryptounicorns/queues"
	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/message"
	"github.com/cryptounicorns/queues/producer"
	"github.com/cryptounicorns/queues/queue/websocket"
	"github.com/cryptounicorns/queues/result"
	"github.com/urfave/cli"
)

var (
	// RootCommands is a list of subcommands for the application.
	RootCommands = []cli.Command{}

	// RootFlags is a list of flags for the application.
	RootFlags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "add this flag to enable debug mode",
		},
	}
)

// RootAction is executing when program called without any subcommand.
func RootAction(c *cli.Context) error {
	var (
		wg  = &sync.WaitGroup{}
		q   queues.Queue
		cr  consumer.Consumer
		pr  producer.Producer
		err error
	)

	if len(c.Args()) == 0 {
		log.Fatal("You should specify an endpoint address in a format <scheme>://<hostname>[:port][/path]")
	}

	q = websocket.New(
		websocket.Config{Addr: c.Args().Get(0)},
		log,
	)
	defer q.Close()

	cr, err = q.Consumer()
	if err != nil {
		log.Fatal(err)
	}
	defer cr.Close()

	pr, err = q.Producer()
	if err != nil {
		log.Fatal(err)
	}
	defer pr.Close()

	go func() {
		defer wg.Done()

		var (
			stream <-chan result.Result
			err    error
		)

		stream, err = cr.Consume()
		if err != nil {
			panic(err)
		}

		for r := range stream {
			switch {
			case r.Err != nil:
				panic(r.Err)
			default:
				log.Debugf("Consumed: %s", r.Value)
				fmt.Printf("%s\n", r.Value)
			}
		}
	}()
	wg.Add(1)

	go func() {
		defer wg.Done()

		var (
			s   = bufio.NewScanner(os.Stdin)
			m   message.Message
			err error
		)

		for s.Scan() {
			m = message.Message(s.Text())

			log.Debugf("Producing: %s", m)
			err = pr.Produce(m)
			if err != nil {
				panic(err)
			}
		}

		err = s.Err()
		if err != nil {
			panic(err)
		}
	}()
	wg.Add(1)

	wg.Wait()

	return nil
}
