#
# Build Container
#

FROM golang:1.18-alpine as go-app-builder

WORKDIR /app

COPY . .

RUN apk add -u -t build-tools curl git && \
    go build -o server *.go && \
    apk del build-tools && \
    rm -rf /var/cache/apk/*


#
# App Runtime Container
#

FROM alpine as go-runtime

WORKDIR /app

RUN apk --no-cache add ca-certificates && \
    apk add curl

COPY --from=go-app-builder /app/server /app/server

EXPOSE 8000

ENTRYPOINT [ "/app/server" ]