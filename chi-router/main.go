package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/v3/layout/{abc}/match", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome " + r.RequestURI))
	})
	r.Get("/v3/layout/{styleId}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		styleId := chi.URLParamFromCtx(ctx, "styleId")
		fmt.Println(r.RequestURI)
		w.Write([]byte("welcome " + styleId))
	})
	r.Get("/v3/layout/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		w.Write([]byte("welcome " + r.RequestURI))
	})
	http.ListenAndServe(":3000", r)
}
