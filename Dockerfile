# Build stage
FROM golang:latest AS builder

RUN apt-get update
RUN apt-get install -y \
    gcc \
    g++ \
    pkg-config \
    libsqlite3-dev \
    libtesseract-dev \
    libleptonica-dev \
    libmupdf-dev \
    tesseract-ocr \
    tesseract-ocr-eng \
    && rm -rf /var/lib/apt/lists/*

ENV CGO_ENABLED=1
ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/5/tessdata/

WORKDIR /srv
COPY ./src /srv

RUN go build -o StoreManagerProServer main.go


# Final stage
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    sqlite3 \
    libsqlite3-0 \
    libtesseract5 \
    liblept5 \
    tesseract-ocr \
    tesseract-ocr-eng \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /srv

COPY --from=builder /srv/StoreManagerProServer /srv/StoreManagerProServer
COPY --from=builder /srv/templates /srv/templates

EXPOSE 8080

CMD ["/srv/StoreManagerProServer"]
