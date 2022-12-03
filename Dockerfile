# syntax=docker/dockerfile:1
FROM golang:alpine as builder

WORKDIR /src

COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -v -a -installsuffix cgo -o github-actions .

FROM scratch 

COPY quotes.txt /
COPY --from=builder /src/github-actions /github-actions

EXPOSE 8080

ENTRYPOINT [ "/github-actions" ]