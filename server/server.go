/**
 * Webserver
 */

package main

import (
  "fmt"
  "log"
  "net/http"
)

func main () {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("> ", r.URL.Path)
  fmt.Fprint(w, "URL.Path = ", r.URL.Path)
}
