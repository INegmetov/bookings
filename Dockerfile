FROM golang:1.21.3-alpine AS builder

# Set destination for COPY
WORKDIR /app

#RUN apk --no-cache add bash git gcc gettext make musl-dev

# Download Go modules
COPY ["app/go.mod", "app/go.sum", "./"]
RUN go mod download

COPY app ./

RUN  CGO_ENABLED=0 GOOS=linux go build -o ./bin/learngo ./cmd/web/

FROM alpine AS runner

COPY --from=builder /app/bin/learngo /
COPY app/templates /templates

EXPOSE 8080

CMD ["/learngo"]
