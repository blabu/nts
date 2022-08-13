FROM golang:1.19-alpine3.16 as builder

WORKDIR /go

ENV APP_NAME=nts \
    APP_REPO=github.com \
    APP_ORG=blabu \
    APP_PKG=${APP_REPO}/${APP_ORG}/${APP_NAME}


ARG VERSION="v0.0.2"

COPY . /go

RUN go build -v -o ${APP_NAME} \
    -ldflags="-X ${APP_ORG}/cmd.version=${VERSION}" \
    main.go

