package main
import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)
const port = 8080

type healthcheckRequest struct {
  Name string `json:"name"`
}

type healthcheckMessage struct {
  Message string `json:"message"`
  Author string `json:"-"`
  Date string `json:",omitempty"`
  Id int `json:"id, string"`
}

func main() {
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

  response := healthcheckMessage{Message: "Ok" + request.Name, Author: "Jame", Date: "2018", Id: 1}
  // data, err := json.Marshal(response)
  // if err != nil {
  //   panic("Something wrong")
  // }
  // fmt.Fprint(w, string(data))
  encoder := json.NewEncoder(w)
  encoder.Encode(response)
}
