package controllers

import "net/http"

/*
	func EnableCors(w *http.ResponseWriter) {
		(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	}
*/

func CheckServerHealth(w http.ResponseWriter, r *http.Request) {
	//	EnableCors(&w)
	if r.Method == "GET" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.WriteHeader(200)
		return
	}
	w.WriteHeader(405)
}
