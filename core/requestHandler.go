package core

import "net/http"

func (app *AppCore[T]) AddRoute(path string, method MethodType, handler RequestHandlerFunction[T]) {
	switch method {
	case GET:
		app.Get(path, app.handleRequest(handler))
	case POST:
		app.Get(path, app.handleRequest(handler))
	case PUT:
		app.Get(path, app.handleRequest(handler))
	case DELETE:
		app.Get(path, app.handleRequest(handler))
	}
}

type RequestHandlerFunction[T dbtype] func(database T, w http.ResponseWriter, r *http.Request)

func (app *AppCore[T]) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (app *AppCore[T]) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (app *AppCore[T]) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (app *AppCore[T]) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("DELETE")
}

func (app *AppCore[T]) handleRequest(handler RequestHandlerFunction[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.Database, w, r)
	}
}
