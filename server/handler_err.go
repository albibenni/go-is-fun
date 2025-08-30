package main

import "net/http"

// handlerErr handles requests to the error endpoint and returns an internal server error
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
