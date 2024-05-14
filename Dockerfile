FROM golang:alpine AS backend

# RUN apk add --no-cache build-base

WORKDIR /app

COPY backend ./

RUN go mod download && \
    GOOS=linux GOARCH=amd64 go build -o /out/tinylink

FROM oven/bun:latest AS frontend

WORKDIR /app

COPY frontend ./

RUN bun install --frozen-lockfile && \
    bun run build

FROM alpine:latest

WORKDIR /app

COPY --from=backend /out/tinylink ./
COPY --from=frontend /app/build ./public

EXPOSE 3000

CMD ["./tinylink"]