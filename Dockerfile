FROM golang:1.20-alpine as build

WORKDIR /build/
COPY mailkuk/ .
COPY go.mod go.sum ./

RUN go build -o mailkuk .

FROM alpine:latest
WORKDIR /app
COPY --from=build /build/mailkuk /app/mailkuk
CMD ["/app/mailkuk"]
