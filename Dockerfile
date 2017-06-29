# Using wheezy from the official golang docker repo
FROM golang:1.6

# Setting up working directory
WORKDIR /go/src/websocketex
Add . /go/src/websocketex/

RUN go get ./...
# Install
RUN go install 

# Setting up environment variables
ENV ENV dev

# My web app is running on port 8080 so exposed that port for the world
EXPOSE 8080
#EXPOSE
ENTRYPOINT ["/go/bin/websocketex"]