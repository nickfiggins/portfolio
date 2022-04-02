FROM golang:latest 
RUN mkdir /app 
ENV GIN_MODE=release
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]