package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/brisouamaury/goroutine-pool-test/app"
	"github.com/brisouamaury/goroutine-pool-test/types"
	"github.com/gorilla/mux"
)

func main() {
	app := app.App{
		ManagementChan: make(types.ManagementChan, 10),
		Runners: map[string]app.Runner{
			"A": {ID: "A", Run: make(chan bool, 10)},
			"B": {ID: "B", Run: make(chan bool, 10)},
		},
	}

	go app.Create(context.Background())

	r := mux.NewRouter()
	r.HandleFunc("/start/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		app.Start(vars["id"])
	})

	r.HandleFunc("/stop/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		app.Stop(vars["id"])
	})

	srv := http.Server{
		Handler: r,
		Addr:    ":8888",
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("closing")
}
