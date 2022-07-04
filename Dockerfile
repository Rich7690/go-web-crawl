FROM chromedp/headless-shell:stable
RUN apt update && apt upgrade -y && apt install -y ca-certificates
COPY go-web-crawl /
ADD https://github.com/bdd/runitor/releases/download/v0.10.0/runitor-v0.10.0-linux-amd64 /usr/local/bin/runitor

RUN chmod +x /usr/local/bin/runitor

ENTRYPOINT ["/usr/local/bin/runitor", "--", "/go-web-crawl"]