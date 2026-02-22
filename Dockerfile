# Build
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ceboostup_compiler

# Run
FROM ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive
WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
	ca-certificates \
	git \
	build-essential \
	pkg-config \
	libcap-dev \
	libsystemd-dev \
	asciidoc \
	xsltproc \
	docbook-xml \
	docbook-xsl \
	&& rm -rf /var/lib/apt/lists/*

RUN git clone https://github.com/ioi/isolate.git /tmp/isolate \
	&& make -C /tmp/isolate \
	&& make -C /tmp/isolate install \
	&& rm -rf /tmp/isolate

COPY --from=builder /app/ceboostup_compiler /app/ceboostup_compiler

CMD ["/app/ceboostup_compiler"]
