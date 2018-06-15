FROM golang:1.10.3-alpine3.7 AS build

ADD . /go/src/github.com/adgear/helm-chart-resource/
RUN apk add --no-cache git ;\
  go get -u github.com/golang/dep/cmd/dep ;\
  dep ensure ;\
  cd /go/src/github.com/adgear/helm-chart-resource ;\
  go build -ldflags "-X main.version=`cat VERSION`" .

FROM alpine:3.7 AS runtime
RUN apk add --no-cache curl jq bash

FROM alpine:3.7 AS helm
RUN wget https://storage.googleapis.com/kubernetes-helm/helm-v2.9.1-linux-amd64.tar.gz ;\
  tar zxvf helm-v2.9.1-linux-amd64.tar.gz ;\
  mv linux-amd64/helm /usr/local/bin/helm

# resource-template Dockerfile used to build docker image on concourse ci
FROM runtime
# REQUIRED BY CONCOURSE RESOURCE
ADD check /opt/resource/check
ADD in /opt/resource/in
ADD out /opt/resource/out
COPY --from=build /go/src/github.com/adgear/helm-chart-resource/helm-chart-resource /usr/local/bin/.
COPY --from=helm /usr/local/bin/helm /usr/local/bin/.

RUN helm init --client-only ;\
  chmod +x /opt/resource/*

WORKDIR /opt/resource