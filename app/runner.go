package app

import (
	"context"
	"fmt"
	"time"
)

type Runner struct {
	ID      string
	Run     chan bool
	running bool
}

func (me *Runner) Watch(ctx context.Context) {
	fmt.Println("watching...")
	for {
		select {
		case r := <-me.Run:
			fmt.Println(r, me.ID)
			me.running = r
		case <-ctx.Done():
			fmt.Println("context Done", ctx.Err())
			return
		default:
			if me.running {
				fmt.Println("I'm Running ! ", me.ID)
				time.Sleep(time.Second * 3)
			}
		}
	}
}

func (me *Runner) Start() {
	fmt.Println("starting", me.ID)
	me.Run <- true
}

func (me *Runner) Stop() {
	fmt.Println("stopping", me.ID)
	me.Run <- false
}
