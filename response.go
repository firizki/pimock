package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	f "github.com/firizki/pimock/flagwrap"
)

type Response struct {
  method    string
  path      string
  variables map[string]string
  resp      []string
}

func NewResponse(method, path string, header, urlq, mapdsc map[string][]string, fw f.FlagWrap) *Response {
  if len(header["Pimock-Sleep"]) > 0 {
    pimock_sleep, err := strconv.Atoi(header["Pimock-Sleep"][0])
    if err != nil {
      panic(err)
    }
    time.Sleep(time.Duration(pimock_sleep) * time.Millisecond)
  }

  if string(path[len(path)-1]) == "/" {
    path = path[:len(path)-1]
  }
  tempPaths := map[string]string{}
  for i, v := range strings.Split(path, "/") {
    tempPaths["{{request.path.["+strconv.Itoa(i)+"]}}"] = v
  }
  for i, v := range urlq {
    tempPaths["{{request.url."+i+"}}"] = v[0]
  }

  targetPath := fmt.Sprintf("%s/%s/%s/%s", *fw.GetRootDirectory(), method, path, *fw.GetResponseFile())
  for i, v := range mapdsc {
    result, err := regexp.MatchString(i, targetPath)
    if err != nil {
      panic(err)
    }
    if result {
      return &Response{method: method, path: path, variables: tempPaths, resp: v}
    }
  }
  return nil
}

func (r Response) getHeaderStatus(ret chan int) {
  result, err := strconv.Atoi(strings.Split(r.resp[0], " ")[1])
  if err != nil {
    panic(err)
  }
  ret <- result
}

func (r Response) getHeaderData(ret chan map[string]string) {
  results := map[string]string{}
  for i := 1; i < len(r.resp); i++ {
    if r.resp[i] == "" {
      i = len(r.resp)
    } else {
      x := strings.Split(r.resp[i], ":")
      results[x[0]] = strings.TrimSpace(x[1])
    }
  }
  ret <- results
}

func (r Response) getBody(ret chan string) {
  result := ""
  for i := 0; i < len(r.resp); i++ {
    if r.resp[i] == "" {
      result = strings.Join(r.resp[i+1:], "\n")
      re := regexp.MustCompile(`({{)([^}]*)(}})`)
      rgx := re.FindAllString(result, -1)
      for _,v := range rgx {
        result = strings.ReplaceAll(result, v, r.variables[v])
      }
      break
    }
  }
  ret <- result
}
