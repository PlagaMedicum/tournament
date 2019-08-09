FROM golang
WORKDIR /go/src/tournament/
COPY main.go .
COPY vendor ./vendor
COPY ./env/testing/databases/postgresql/config.yaml ./databases/postgresql
COPY ./env/testing/databases/postgresql/migrations ./databases/postgresql/migrations
COPY go.mod .
COPY go.sum .
COPY ./env/testing/run.sh .
RUN chmod +x ./run.sh
CMD ["./run.sh"]