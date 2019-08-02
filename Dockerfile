FROM golang
WORKDIR /go/src/tournament/
COPY main.go .
COPY vendor .
COPY databases .
COPY go.mod .
COPY run.sh .
RUN chmod +x ./run.sh