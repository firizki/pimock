package main

import (
  "testing"
)

func TestDiscover(t *testing.T) {
  type unit struct {
    folder string
    file   string
    result int
  }

  cases := []unit{
    unit{"responses", "response", 0},
    unit{"response", "json", 0},
  }

  for _, c := range cases {
    switch c.result {
    case 0:
      md := NewDiscover(c.folder, c.file)
      if md == nil {
        t.Error()
      }
    case 1:
      NewDiscover(c.folder, c.file)
    }
  }
}
