FROM golang:1.27-alpine AS build

WORKDIR /build/
COPY mailkuk/ .
COPY go.mod go.sum ./

RUN go build -o mailkuk .

FROM alpine:latest
WORKDIR /app
COPY --from=build /build/mailkuk /app/mailkuk
CMD ["/app/mailkuk"]
