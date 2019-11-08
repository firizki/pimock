package main

import (
    "os"
    "regexp"
    "path/filepath"
    "io/ioutil"
    "strings"
)

type Discover struct {
  maps map[string][]string
}

func NewDiscover(root, ext string) *Discover {
  paths := map[string][]string{}
  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    result, err := regexp.MatchString(ext+"$", path)
    if err != nil {
      return err
    }
    if result {
      dat, err := ioutil.ReadFile(path)
      if err != nil {
        panic(err)
      }
      paths[path] = strings.Split(string(dat), "\n")
    }
    return nil
  })
  if err != nil {
    panic(err)
  }
  return &Discover{paths}
}
