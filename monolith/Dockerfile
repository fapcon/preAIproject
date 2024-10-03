FROM golang:1.19.3-alpine

COPY ./keys /root/.ssh/

RUN apk update && apk add --no-cache git openssh && \
    git config --global url."git@studentgit.kata.academy:".insteadOf "https://studentgit.kata.academy/" && \
    chmod 600 /root/.ssh/id_ed25519 && \
    ssh-keygen -F studentgit.kata.academy || ssh-keyscan studentgit.kata.academy >> /root/.ssh/known_hosts


WORKDIR /go/src/studentgit.kata.academy/eazzyearn/students/mono/monolith

COPY . .