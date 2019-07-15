FROM golang
WORKDIR /go/src/tournament/
COPY run.sh .
RUN chmod +x ./run.sh