FROM golang:latest 
RUN export GOPATH=$HOME/go
RUN mkdir $GOPATH/src/github.com/sukumDev123/restapi
COPY . $GOPATH/src/github.com/sukumDev123/restapi
WORKDIR $GOPATH/src/github.com/sukumDev123/restapi
RUN go get -u github.com/gorilla/mux

