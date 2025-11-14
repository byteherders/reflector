# syntax=docker/dockerfile:1.5

FROM golang:1.21 AS build
WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /out/reflector ./cmd/reflector

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /
COPY --from=build /out/reflector /reflector

ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/reflector"]
