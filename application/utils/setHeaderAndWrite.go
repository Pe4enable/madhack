package utils

import "net/http"

func SetHeadersAndWrite(w http.ResponseWriter, body []byte, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(body)
}
