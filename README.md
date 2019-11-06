# Ailea
Simple Mock API using Golang

## Requirement
- golang 1.12

## How To Use
1. Create your `response` file under directory that same with request path
    For example you want to create mock for GET /healthz
    Create file `response` under `responses/GET/healthz/`
    folder `responses` is mandatory
2. Write you header and body response HTTP inside your `body` file
    For example to send simple 200 OK you can do this
    ```
    HTTP/1.1 200 OK

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
