# Build stage
FROM golang:latest AS builder

ARG VERSION
ARG COMMIT

RUN apt-get update
RUN apt-get install -y \
    gcc \
    g++ \
    git \
    pkg-config \
    libtesseract-dev \
    libleptonica-dev \
    libmupdf-dev \
    tesseract-ocr

ENV CGO_ENABLED=1
ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/5/tessdata/

WORKDIR /srv
COPY ./src /srv

RUN --mount=type=cache,target=/go/pkg/mod \
    test -n "$VERSION" -a -n "$COMMIT" || { echo "ERROR: VERSION and COMMIT build args are required (--build-arg VERSION=... --build-arg COMMIT=...)"; exit 1; } && \
    go build -ldflags "-X gitea.locker98.com/locker98/Store-Manager-Pro/utils.Version=${VERSION} -X gitea.locker98.com/locker98/Store-Manager-Pro/utils.Commit=${COMMIT}" -o StoreManagerProServer main.go


# Final stage
FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y \
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
