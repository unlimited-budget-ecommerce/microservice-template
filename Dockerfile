# Base image
FROM golang:1.24.5-alpine3.22 AS base

WORKDIR /app

RUN apk update && apk add --no-cache tzdata

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o goapp .

# Final image
FROM gcr.io/distroless/static-debian12:nonroot

ENV TZ=Asia/Bangkok

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /app/goapp .
COPY --from=base /app/config/config.yaml ./config

EXPOSE 3030

CMD ["./goapp"]
