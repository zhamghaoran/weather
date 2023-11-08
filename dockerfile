FROM golang:1.20
WORKDIR /app
ADD . /app
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN go build main.go
EXPOSE 8080
ENTRYPOINT ["./main"]