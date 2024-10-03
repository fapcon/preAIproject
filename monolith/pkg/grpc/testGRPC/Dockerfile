FROM golang:alpine
COPY . /go/src/studentgit.kata.academy/eazzyearn/students/mono/monolith/test
WORKDIR /go/src/studentgit.kata.academy/eazzyearn/students/mono/monolith/test
RUN go build -o /go/bin/server /go/src/studentgit.kata.academy/eazzyearn/students/mono/monolith/test

FROM alpine
COPY --from=0 /go/bin/server /app

ENTRYPOINT ["/app"]