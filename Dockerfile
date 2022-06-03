FROM alpine:3.15
COPY go-web-crawl /
ENTRYPOINT ["/go-web-crawl"]
