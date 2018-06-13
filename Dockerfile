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
ADD common.sh /opt/resource/common.sh

COPY --from=helm /usr/local/bin/helm /usr/local/bin/.

RUN helm init --client-only ;\
  chmod +x /opt/resource/*

WORKDIR /opt/resource