FROM golang:1.14.2 AS builder
WORKDIR /go/src/github.com/evsyukovmv/taskmanager
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/evsyukovmv/taskmanager/migrate.linux-amd64 .
COPY --from=builder /go/src/github.com/evsyukovmv/taskmanager/db db/
COPY --from=builder /go/src/github.com/evsyukovmv/taskmanager/app .
COPY --from=builder /go/src/github.com/evsyukovmv/taskmanager/docker-entrypoint.sh .

CMD ["./docker-entrypoint.sh"]
