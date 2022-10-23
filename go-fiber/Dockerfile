#
# Build Container
#

FROM golang:1.18-alpine as go-bookstore-fiber-builder

WORKDIR /app

COPY . .

RUN apk add -u -t build-tools curl git && \
    go build -o bookstore-fiber-server cmd/main/*.go && \
    cp config.yml /config.yml && \
    apk del build-tools && \
    rm -rf /var/cache/apk/*


#
# App Runtime Container
#

FROM alpine as go-bookstore-fiber-runtime

WORKDIR /app

RUN apk --no-cache add ca-certificates && \
    apk add curl

COPY --from=go-bookstore-fiber-builder /app/bookstore-fiber-server /app/bookstore-fiber-server
COPY --from=go-bookstore-fiber-builder /config.yml /config.yml

EXPOSE 9010

ENTRYPOINT [ "/app/bookstore-fiber-server" ]
