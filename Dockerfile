FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN make build

FROM scratch
COPY --from=builder /app/tmp/server .
CMD ["./server"]
