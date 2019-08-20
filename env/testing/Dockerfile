FROM golang
WORKDIR /go/src/tournament/
COPY main.go .
COPY vendor ./vendor
COPY ./env/testing/databases/postgresql ./databases/postgresql
COPY go.mod .
COPY go.sum .
COPY ./env/testing/run.sh .
RUN chmod +x ./run.sh
CMD ["./run.sh"]