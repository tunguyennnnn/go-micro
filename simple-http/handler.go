package main

import (
  "context"
  "net/http"
  "fmt"
  "log"
  "encoding/json"
)

type validationContextKey string

type healthcheckResponse struct {
  Message string `json:"message"`
}

type healthcheckRequest struct {
  Name string `json:"name"`
}

const port = 8080

func main() {
  handler := newValidationHandler(newHealthCheckHandler())
  http.Handle("/healthcheck", handler)

  log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationHandler struct {
  next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
  return validationHandler{next: next}
}

type healthcheckHandler struct{}

func newHealthCheckHandler() http.Handler {
  return healthcheckHandler{}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request healthcheckRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
  c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

func (h healthcheckHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
  name := r.Context().Value(validationContextKey("name")).(string)
	response := healthcheckResponse{Message: "Hello" + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
