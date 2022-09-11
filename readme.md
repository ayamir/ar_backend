# Requirement

1. `go`
2. `mysql`

# How to run?

`go run .`

# How to test with `curl`?

-   Insert record

    `curl http://localhost:8083/info \ --include \ --header "Content-Type: application/json" \ --request "POST" \ --data '{"code": "3015","name": "lgt","motto": "Hardwork always pays off, whatever you do."}'`

-   Query record with `code`

    `curl http://localhost:8083/info/3015 \ --include \ --header "Content-Type: application/json" \ --request "GET"`

    (`code` is `3015` here)

-   Query all records

    `curl http://localhost:8083/infos \ --include \ --header "Content-Type: application/json" \ --request "GET"`
