FROM golang:latest 
RUN mkdir /opt/golang
COPY . /opt/golang
WORKDIR /opt/golang
RUN go get -u github.com/gorilla/mux
RUN go build -o main .