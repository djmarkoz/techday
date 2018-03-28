FROM golang:1.10 as builder
ENV project github.com/djmarkoz/techday
WORKDIR /go/src/$project/cmd/techday-cli
COPY . /go/src/$project
RUN GOOS=linux go build -o /tmp/techday-cli

FROM busybox:glibc
WORKDIR /opt
CMD ["./techday-cli"]
COPY --from=builder /tmp/techday-cli .