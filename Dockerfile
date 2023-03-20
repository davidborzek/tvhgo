FROM node:18.12.1 AS ui_build

ARG VERSION
ARG GIT_COMMIT

WORKDIR /build
COPY ui ui
WORKDIR /build/ui

RUN yarn install --frozen-lockfile
RUN yarn build

FROM golang:1.19.4-alpine3.17 AS build
RUN apk add --no-cache build-base

ENV CGO_ENABLED=1 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .
COPY --from=ui_build /build/ui/dist ui/dist

RUN go mod download
RUN go build -o tvhgo -tags prod main.go

FROM alpine:3.17 as prod

COPY --from=build /build/tvhgo /bin/tvhgo

RUN adduser -D -H tvhgo

WORKDIR /tvhgo
RUN chown -R tvhgo:tvhgo /tvhgo

USER tvhgo
EXPOSE 8080
VOLUME ["/tvhgo"]
ENV TVHGO_DATABASE_PATH /tvhgo/tvhgo.db

ENTRYPOINT ["/bin/tvhgo"]
