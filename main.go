package main

import (
	"net/http"
	"net/http/httputil"

	"./app"
	"./page"
)

func main() {
	// initialization
	app.Init()

	// reverse proxy to laravel admin panel app
	http.HandleFunc("/admin/", adminHandler)
	// main handler
	http.HandleFunc("/", appHandler)

	//server bind
	http.ListenAndServe(app.GetEnv("MAIN_HOST"), nil)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	director := func(req *http.Request) {
		req = r
		req.URL.Scheme = app.GetEnv("ADMIN_SCHEME")
		req.URL.Host = app.GetEnv("ADMIN_HOST")
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	app.SetDefaultHeaders(w)
	page := new(page.Content)

	if r.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		page.Add(app.HomeController(w, r))
	} else {
		w.WriteHeader(http.StatusNotFound)
		page.Add(app.NotFoundController(w, r))
	}

	app.GzipWrite(page.Get(), w)
}
