FROM golang:1.12-alpine3.9
RUN mkdir /publix-sub-api
ADD . /publix-sub-api
WORKDIR /publix-sub-api
RUN apk add git
RUN go get github.com/sasho2k/publix-sub-api
ENV PORT 3000
RUN  go build -o main .
CMD ["/publix-sub-api/main"]