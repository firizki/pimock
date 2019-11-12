package main

import (
  "regexp"
  "strings"
  "strconv"
)

type Response struct {
  method    string
  path      string
  variables map[string]string
  resp      []string
}

func NewResponse(method, path string, urlq map[string][]string) *Response {
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

  targetPath := "responses/"+method+"/"+path+"/response"
  for i, v := range mapDiscover.maps {
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

func (r Response) getHeaderStatus() int {
  result, err := strconv.Atoi(strings.Split(r.resp[0], " ")[1])
  if err != nil {
    panic(err)
  }
  return result
}

func (r Response) getHeaderData() map[string]string {
  results := map[string]string{}
  for i := 1; i < len(r.resp); i++ {
    if r.resp[i] == "" {
      i = len(r.resp)
    } else {
      x := strings.Split(r.resp[i], ":")
      results[x[0]] = x[1]
    }
  }
  return results
}

func (r Response) getBody() string {
  for i := 0; i < len(r.resp); i++ {
    if r.resp[i] == "" {
      result := strings.Join(r.resp[i+1:], "\n")
      re := regexp.MustCompile(`({{)([^}]*)(}})`)
      rgx := re.FindAllString(result, -1)
      for _,v := range rgx {
        result = strings.ReplaceAll(result, v, r.variables[v])
      }
      return result
    }
  }
  return ""
}
