FROM golang:1.10.3-alpine3.7 AS build

ADD . /go/src/github.com/adgear/helm-chart-resource/
RUN apk add --no-cache git ;\
  go get -u github.com/golang/dep/cmd/dep ;\
  cd /go/src/github.com/adgear/helm-chart-resource ;\
  dep ensure ;\
  go build -ldflags "-X main.version=`cat VERSION`" .

FROM alpine:3.7 AS runtime
RUN apk add --no-cache curl jq bash

FROM alpine/helm:3.4.2 AS helm

# resource-template Dockerfile used to build docker image on concourse ci
FROM runtime
# REQUIRED BY CONCOURSE RESOURCE
ADD check /opt/resource/check
ADD in /opt/resource/in
ADD out /opt/resource/out
COPY --from=build /go/src/github.com/adgear/helm-chart-resource/helm-chart-resource /usr/local/bin/.
COPY --from=helm /usr/bin/helm /usr/bin/helm/.

RUN chmod +x /opt/resource/*

WORKDIR /opt/resource