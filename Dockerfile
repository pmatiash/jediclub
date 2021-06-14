FROM golang:1.14.9-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum main.go validator.go index.html /build/
WORKDIR /build
RUN go build -o jedi

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/jedi /app/
COPY index.html/ /app/index.html
WORKDIR /app
CMD ["./jedi"]
