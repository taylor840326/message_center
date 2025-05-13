FROM golang:1.24.3

WORKDIR /

ADD . /app/

RUN cd /app && \
    go build main.go -o main

ENTRYPOINT [ "/app/main" ]

