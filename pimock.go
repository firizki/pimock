package main

import (
    "net/http"
)

var mapDiscover *Discover

func main() {
  mapDiscover = NewDiscover("responses", "response")

  http.HandleFunc("/", MockServer)
  http.ListenAndServe(":8080", nil)
}

func MockServer(w http.ResponseWriter, r *http.Request) {
  res := NewResponse(r.Method, r.URL.Path[1:], r.Header, r.URL.Query(), mapDiscover.maps)

  if res == nil {
    w.WriteHeader(404)
    w.Write([]byte("Not Found"))
  } else {
    headerStatus := res.getHeaderStatus()
    headerData := res.getHeaderData()
    body := res.getBody()

    for i, x := range headerData {
      w.Header().Set(i, x)
    }

    w.WriteHeader(headerStatus)
    w.Write([]byte(body))
  }

}
