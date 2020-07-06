FROM golang:1.14.4
MAINTAINER Gabriel Tomazi
RUN mkdir /delivery-challenge 
ADD . /delivery-challenge/ 
WORKDIR /delivery-challenge 
RUN go build ./cmd/main.go 
CMD ["/delivery-challenge/main"]
EXPOSE 3002
