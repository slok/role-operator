FROM golang:1.10-alpine

RUN apk --no-cache add \
    g++ \
    git

# Mock creator
RUN go get -u github.com/vektra/mockery/.../

# Create user
ARG uid=1000
ARG gid=1000
RUN addgroup -g $gid roleoperator && \
    adduser -D -u $uid -G roleoperator roleoperator && \
    chown roleoperator:roleoperator -R /go


USER roleoperator
WORKDIR /go/src/github.com/slok/role-operator