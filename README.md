

# Ailea
Simple Mock API using Golang

## Requirement
- golang 1.12

## How To Use
1. Create your `response` file under directory that same with request path  
    For example you want to create mock for GET `/healthz`  
    Create file `response` under `responses/GET/healthz/`. Folder `responses` is mandatory
2. Write you header and body response HTTP inside your `response` file. The format is following [W3 HTTP/1.1 Response](https://www.w3.org/Protocols/rfc2616/rfc2616-sec6.html).  
    For example to send simple 200 OK you can do this
    ```
    HTTP/1.1 200 OK
    Content-Type: text/plain; charset=utf-8

    OK
    ```
3. Build binary
    ```
    go build
    ```
4. Run ailea
    ```
    ./ailea
    ```
5. Access from port `8080`
    ```
    curl localhost:8080/healthz -v
    ```

## Features

<details>
  <summary>Regex Match Path Feature</summary>

  Create `response` file under directory with regex name, and it'll automatically find by regex
  Example :
  ```
  responses/GET/users/([0-9]*)/response
  ```
  Will match any GET request with paths
  ```
  curl localhost:8080/users/1
  curl localhost:8080/users/123/
  curl localhost:8080/users/9898?params=value
  ```
</details>

<details>
  <summary>Template Feature</summary>

  Currently it only support to get path request under variable `{{request.path.[i]}}`
  Example :
  Response file like this
  ```
  HTTP/1.1 200 OK
  Content-Type: application/json

  {
    "path": "{{request.path.[0]}}"
  }
  ```

  Will give response

  ```
  {
    "path": "users"
  }
  ```
</details>
