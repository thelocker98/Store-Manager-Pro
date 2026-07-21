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
    tesseract-ocr

ENV CGO_ENABLED=1
ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/5/tessdata/

WORKDIR /srv
COPY ./src /srv

RUN go build -o StoreManagerProServer main.go


# Final stage
FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y \
    sqlite3 \
    libsqlite3-0 \
    libtesseract-dev \
    libleptonica-dev \
    tesseract-ocr-eng

WORKDIR /srv

COPY --from=builder /srv/StoreManagerProServer /srv/StoreManagerProServer
COPY --from=builder /srv/static /srv/static
COPY --from=builder /srv/templates /srv/templates

ENV GIN_MODE=release
ENV DATA_PATH=/data

EXPOSE 8080

CMD ["/srv/StoreManagerProServer"]
