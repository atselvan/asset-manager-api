FROM golang:1.13.4

RUN mkdir /appl

ENV GOOS=linux \
    GOARCH=amd64

RUN apt-get update  && apt-get install jq -y

COPY . /appl/asset-manager

WORKDIR /appl/asset-manager

RUN go get ./... && go build && mv asset-manager /usr/local/bin

EXPOSE 8080

CMD ["asset-manager"]