package main

import "net/http"

// handlerReadiness responds with a simple JSON indicating the service is ready.
// According to go http documentation, handlers should not panic and should handle errors gracefully.
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
