package app

import (
	"net/http"
	"strconv"

	"../frontend"
)

const (
	// srcRoot filepath to project root considering we are in bin
	projectRoot = "../"
	// srcFolder filepath to src
	srcFolder = "src/"
)

var (
	// env - content from parsed .env file
	env Enviroment
)

// Init application
func Init() {
	initEnv() // w8 till it will be done first time
	go liveEnv()
}

// SetDefaultHeaders sets default server's http headers
func SetDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Keep-Alive", "timeout=15")
	w.Header().Set("Content-Encoding", "gzip")
	SetHTTPCache(w, 60) // one minute by default
}

// SetHTTPCache set http cache-control
func SetHTTPCache(w http.ResponseWriter, seconds int) {
	w.Header().Set("Cache-Control", "max-age="+strconv.Itoa(seconds)+",must-revalidate")
}

// HomeController main page controller
func HomeController(w http.ResponseWriter, r *http.Request) string {
	abc := frontend.ExampleExecute()
	return "home" + strconv.Itoa(int(abc))
}

// NotFoundController 404
func NotFoundController(w http.ResponseWriter, r *http.Request) string {
	return "not found"
}
