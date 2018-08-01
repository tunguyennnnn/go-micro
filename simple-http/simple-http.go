package main
import (
  "fmt"
  "log"
  "net/http"
)

const port = 8080

func main() {
  http.HandleFunc("/healthcheck", healthcheckHandler)
  log.Printf("Server starting on port %v\n", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Ok")
}
