package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

type AppCore[T dbtype] struct {
	Router   *mux.Router
	Database T
}

func (app *AppCore[T]) ConfigureRouter() {
	app.Router = mux.NewRouter()
}

func (app *AppCore[T]) Run(host string) {
	http.ListenAndServe(host, app.Router)
}
