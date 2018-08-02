package main
import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)
const port = 8080

type healthcheckMessage struct {
  Message string
}

func main() {
  http.HandleFunc("/healthcheck", healthcheckHandler)
  log.Printf("Server starting on port %v\n", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
  response := healthcheckMessage{Message: "Ok"}
  data, err := json.Marshal(response)
  if err != nil {
    panic("Something wrong")
  }
  fmt.Fprint(w, string(data))
}
