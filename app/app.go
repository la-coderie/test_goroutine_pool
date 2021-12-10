package app

import (
	"context"
	"fmt"

	"github.com/brisouamaury/goroutine-pool-test/types"
)

type App struct {
	ManagementChan types.ManagementChan
	Runners        map[string]Runner
}

func (me *App) Create(ctx context.Context) {
	for i := range me.Runners {
		r := me.Runners[i]
		go r.Watch(ctx)
	}

	for i := range me.ManagementChan {
		r := me.Runners[i.ID]

		fmt.Println(i, me.ManagementChan)
		if i.Action == "start" {
			r.Start()
		}

		if i.Action == "stop" {
			r.Stop()
		}

	}

	<-ctx.Done()
}

func (me *App) Start(id string) {
	me.ManagementChan <- types.ManagementEntry{ID: id, Action: "start"}
}

func (me *App) Stop(id string) {
	me.ManagementChan <- types.ManagementEntry{ID: id, Action: "stop"}
}
