FROM golang
WORKDIR /go/src/stringsvc
COPY . .

RUN go get -t
RUN go build
CMD ["./stringsvc"]
