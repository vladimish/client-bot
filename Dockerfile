FROM golang:1.17.5-alpine

WORKDIR /app
COPY . ./

RUN go mod download
RUN go build -o /ranhb cmd/main.go

EXPOSE 1721

CMD ["/ranhb"]
