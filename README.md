# Ailea
Simple Mock API using Golang

## Requirement
- golang 1.12

## How To Use
1. Create your `body` file under directory that same with request path
    For example you want to create mock for GET /healthz
    Create file `body` under `responses/GET/healthz/`
    folder `responses` is mandatory
2. Write you response HTTP inside your `body` file
    For example to send simple 200 OK you can do this
    ```
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
    curl localhost:1080/healthz -v
    ```
