package main

import (
  "fmt"
  "encoding/json"
  "log"
  "net/http"
)

type healthcheckRequest struct {
  Name string `json:"name"`
}

type healthcheckMessage struct {
  Message string `json:"message"`
}

const port = 8080

func main() {
  server()
}

func server() {
  imageHandler := http.FileServer(http.Dir("./images"))
  http.Handle("/images/", http.StripPrefix("/images/", imageHandler))

  http.HandleFunc("/healthcheck", healthcheckHandler)

  log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
  var request healthcheckRequest
  decoder := json.NewDecoder(r.Body)

  err := decoder.Decode(&request)
  if (err != nil) {
    http.Error(w, "Bad request", http.StatusBadRequest)
    return
  }

  response := healthcheckMessage{Message: "Hello " + request.Name}
  encoder := json.NewEncoder(w)
  encoder.Encode(response)
}
