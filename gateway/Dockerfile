FROM golang:1.20.5-bullseye
RUN apt update && apt upgrade -y
RUN mkdir /app
WORKDIR /usr/src/gw/
COPY . .
RUN go mod download
WORKDIR /usr/src/gw/cmd
RUN go build -o /app/gateway
RUN rm -rf /usr/src/gw
CMD ["/app/gateway"]