# install from source

    go get -u github.com/tednaleid/echoserver
    
## fmt/lint/install

    nocorrect go fmt && golint && go test ./... && go install

## test GET with `ganda`

    time seq 10000 | awk '{printf "http://localhost:1323/%s\n", $1}' | ganda -s

## test POST with `ganda` using body

    time seq 10000 | awk '{printf "http://localhost:1323/%s %s\n", $1, $1}' | ganda -s -X POST -d '{"id": "%s"}'
