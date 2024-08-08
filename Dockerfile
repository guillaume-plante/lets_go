# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.22.5 AS build-stage

WORKDIR /app

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./web ./cmd/web

# Deploy the application binary into a lean image
FROM scratch AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/web web
COPY --from=build-stage /app/ui /ui

EXPOSE 4000

ENTRYPOINT ["./web"]