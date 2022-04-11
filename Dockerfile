FROM golang:latest as builder

RUN mkdir /build
WORKDIR /build
ADD . /build

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0  GO111MODULE=on\
    go build -o orders /api/main.go

FROM scratch

COPY --from=builder /build/orders .

# ENTRYPOINT [ "executable" ]
CMD ["./orders"]
