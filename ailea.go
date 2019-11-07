package main

import (
    "net/http"
    "io/ioutil"
    "strings"
    "strconv"
    "os"
    "regexp"
    "path/filepath"
)

var mappings []string

func discoverPath(in, ext string) ([]string, error) {
  var paths []string
  err := filepath.Walk(in, func(path string, info os.FileInfo, err error) error {
      result, err := regexp.MatchString(ext+"$", path)
      if err != nil {
        return err
      }
      if result {
        paths = append(paths, path)
      }
      return nil
  })
  if err != nil {
      return paths, err
  }
  return paths, nil
}

func main() {
    mappings, _ = discoverPath("responses", "response")

    http.HandleFunc("/", MockServer)
    http.ListenAndServe(":8080", nil)
}

func getHeaderStatus(in []string) (int, error) {
  return strconv.Atoi(strings.Split(in[0], " ")[1])
}

func getHeaderData(in []string) map[string]string {
  results := map[string]string{}
  for i := 1; i < len(in); i++ {
    if in[i] == "" {
      i = len(in)
    } else {
      x := strings.Split(in[i], ":")
      results[x[0]] = x[1]
    }
  }
  return results
}

func getBody(in []string) string {
  for i := 0; i < len(in); i++ {
    if in[i] == "" {
      return strings.Join(in[i+1:], "\n")
    }
  }
  return ""
}

func getResponse(method, path string) string {
  if string(path[len(path)-1]) == "/" {
    path = path[:len(path)-1]
  }
  targetPath := "responses/"+method+"/"+path+"/response"
  for _, v := range mappings {
    result, err := regexp.MatchString(v, targetPath)
    if err != nil {
      panic(err)
    }
    if result {
      dat, err := ioutil.ReadFile(v)
      if err != nil {
        panic(err)
      }
      return string(dat)
    }
  }
  return ""
}

func MockServer(w http.ResponseWriter, r *http.Request) {
  res := getResponse(r.Method, r.URL.Path[1:])

  if res == "" {
    w.WriteHeader(404)
    w.Write([]byte("Not Found"))
  } else {
    metaData := strings.Split(res, "\n")
    headerStatus, _ := getHeaderStatus(metaData)
    headerData := getHeaderData(metaData)
    body := getBody(metaData)

    for i, x := range headerData {
      w.Header().Set(i, x)
    }

    w.WriteHeader(headerStatus)
    w.Write([]byte(body))
  }

}
