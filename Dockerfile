FROM golang:1.12-alpine3.10 as builder

RUN apk add --no-cache git

#ENV CGO_ENABLED 0
ENV GO111MODULE on

ARG D=/go/src/github.com/nhannvt/resume

COPY go.mod $D/
COPY go.sum $D/
RUN cd $D && go mod download

COPY . $D
RUN cd $D/cmd && go build -o sforum-api && cp sforum-api /tmp/


FROM alpine:3.10

ARG USER=23hdev

COPY --from=builder /tmp/sforum-api /usr/local/bin/

RUN apk add ca-certificates
RUN addgroup -S $USER && adduser -S $USER -G $USER
RUN chown $USER:$USER /usr/local/bin/sforum-api

USER $USER
CMD ["/usr/local/bin/sforum-api"]
