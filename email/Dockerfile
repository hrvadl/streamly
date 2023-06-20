FROM golang:1.20.5-bullseye
RUN apt update && apt upgrade -y
RUN mkdir /app
WORKDIR /usr/src/email
COPY . .
RUN go mod download
WORKDIR /usr/src/email/cmd
RUN go build -o /app/email
RUN rm -rf /usr/src/email
CMD ["/app/email"]
