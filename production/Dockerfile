FROM golang:1.23.0-bookworm AS builder

WORKDIR /work/app

COPY ../ ./
RUN go mod tidy
RUN go build -ldflags="-s -w" -trimpath ./src/cmd/main.go

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /work/app/main ./main
COPY --from=builder /work/app/src/driver/db/migrate/sql ./src/driver/db/migrate/sql
CMD ["./main"]
