FROM golang:1.12 AS builder
WORKDIR /go/src/tournament/
COPY main.go .
COPY vendor ./vendor
COPY ./env/staging/databases ./databases
COPY pkg ./pkg
COPY go.mod .
COPY ./env/staging/build.sh .
RUN chmod +x ./build.sh
RUN ./build.sh

FROM scratch
COPY --from=builder /go/src/tournament/bin/main /
COPY ./env/staging/databases /databases
CMD ["/main"]
