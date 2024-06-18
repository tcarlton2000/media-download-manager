FROM golang:1.22-alpine AS builder

RUN apk add build-base

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o /dist/main ./...

ADD --chmod=777 https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-linux-x64 /app/tailwindcss
RUN ./tailwindcss -i input.css -o static/css/output.css

FROM alpine:latest AS production

RUN apk add libc6-compat python3
ADD --chmod=777 https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp /usr/local/bin/yt-dlp

COPY --from=builder /dist .
COPY --from=builder /app/static static

EXPOSE 8000
CMD ["./main"]
