FROM golang:1.22-alpine AS builder

RUN apk add build-base

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o /dist/main ./...
RUN ldd /dist/main | tr -s [:blank:] '\n' | grep ^/ | xargs -I % install -D % /dist/%
RUN ln -s ld-musl-x86_64.so.1 /dist/lib/libc.musl-x86_64.so.1

ADD https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-linux-x64 /app/tailwindcss
RUN chmod a+x tailwindcss
RUN ./tailwindcss -i input.css -o static/css/output.css

FROM alpine:latest AS production

COPY --from=builder /dist .
COPY --from=builder /app/static static
COPY --from=builder /app/templates templates

ADD https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp /usr/local/bin/yt-dlp

EXPOSE 8000
CMD ["./main"]
