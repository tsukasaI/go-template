FROM golang:1.22.5-bookworm AS builder

WORKDIR /work/app

COPY ../ ./
RUN go mod tidy
RUN go build ./src/cli/main.go

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /work/app ./
CMD ["./main"]
