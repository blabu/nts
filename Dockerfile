FROM golang:1.19-alpine3.16 as builder

RUN apk add --no-cache --update git

ENV APP_NAME=nts \
    APP_REPO=github.com \
    APP_ORG=blabu

ENV APP_PKG=${APP_REPO}/${APP_ORG}/${APP_NAME}

WORKDIR /${APP_NAME}

ARG VERSION="v0.0.0"

COPY . .

RUN export GIT_COMMIT="$(git rev-parse HEAD)" && \
    go build -v -o ${APP_NAME} \
    -ldflags="\
    -X ${APP_PKG}/cmd.version=${VERSION} \
    -X ${APP_PKG}/cmd.subversion=${GIT_COMMIT} \
    " .

FROM alpine:3.16
COPY --from=builder /nts/nts /bin/

