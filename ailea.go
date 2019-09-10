package main

import (
    "net/http"
    "io/ioutil"
)

func main() {
    http.HandleFunc("/", MockServer)
    http.ListenAndServe(":8080", nil)
}

func MockServer(w http.ResponseWriter, r *http.Request) {
  dat, err := ioutil.ReadFile("responses/"+r.Method+"/"+r.URL.Path[1:]+"/body")
  if err != nil {
    w.WriteHeader(404)
    w.Write([]byte("Not Found"))
  } else {
    w.Header().Set("Content-Type", "application/json")
    w.Write(dat)
  }
}
