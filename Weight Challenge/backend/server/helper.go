package server

import "net/http"

func AllowCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}
