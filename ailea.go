package main

import (
    "net/http"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {
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
      return strings.Join(in[i:], "")
    }
  }
  return ""
}

func MockServer(w http.ResponseWriter, r *http.Request) {
  dat, err := ioutil.ReadFile("responses/"+r.Method+"/"+r.URL.Path[1:]+"/response")

  if err != nil {
    w.WriteHeader(404)
    w.Write([]byte("Not Found"))
  } else {
    metaData := strings.Split(string(dat), "\n")
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
