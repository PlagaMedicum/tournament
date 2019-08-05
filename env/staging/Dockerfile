FROM golang
WORKDIR /go/src/tournament/
COPY main.go .
COPY vendor ./vendor
COPY ./env/staging/databases ./databases
COPY go.mod .
COPY go.sum .
COPY ./env/staging/run.sh .
RUN chmod +x ./run.sh