FROM golang:1.23.1-alpine3.20 AS build
RUN apk add --no-cache build-base

ENV CGO_ENABLED=1
WORKDIR /build

COPY . .

RUN go mod download
RUN go build -o tvhgo -tags prod main.go

FROM alpine:3.20 AS prod

COPY --from=build /build/tvhgo /bin/tvhgo

RUN adduser -D -H tvhgo

WORKDIR /tvhgo
RUN chown -R tvhgo:tvhgo /tvhgo

USER tvhgo
EXPOSE 8080
VOLUME ["/tvhgo"]
ENV TVHGO_DATABASE_PATH=/tvhgo/tvhgo.db

ENTRYPOINT ["/bin/tvhgo"]
