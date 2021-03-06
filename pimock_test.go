package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestMockServer(t *testing.T) {
    type unit struct {
      mapd            *Discover
      request_method  string
      request_path    string
      body_result     string
      status_code     int
    }

    cases := []unit{
      unit{&Discover{maps: map[string][]string{"healthz": []string{"HTTP/1.1 200 OK", "Content-Type: text/plain; charset=utf-8" ,"", "OK"}}}, "GET", "/healthz", "OK", 200},
      unit{&Discover{maps: map[string][]string{"healthz": []string{"HTTP/1.1 200 OK", "", "OK"}}}, "GET", "/hello", "Not Found", 404},
    }

    for _, c := range cases {
      mapDiscover = c.mapd

      req, err := http.NewRequest(c.request_method, c.request_path, nil)
      if err != nil {
          t.Fatal(err)
      }

      rr := httptest.NewRecorder()
      handler := http.HandlerFunc(MockServer)

      handler.ServeHTTP(rr, req)

      if status := rr.Code; status != c.status_code {
          t.Errorf("handler returned wrong status code: got %v want %v",
              status, c.status_code)
      }

      if rr.Body.String() != c.body_result {
          t.Errorf("handler returned unexpected body: got %v want %v",
              rr.Body.String(), c.body_result)
      }

    }

}
