FROM golang
WORKDIR /go/src/tournament/
COPY prepare.sh .
COPY connconf.yaml .
COPY main.go .
COPY go.mod .
COPY go.sum .
COPY migrations ./migrations
COPY pkg ./pkg
RUN chmod +x ./prepare.sh
RUN ./prepare.sh