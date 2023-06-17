FROM golang:1.20.5-bullseye
RUN apt update && apt upgrade -y
RUN mkdir /app
WORKDIR /usr/src/auth
COPY . .
RUN go mod download
WORKDIR /usr/src/auth/cmd
RUN go build -o /app/auth
RUN rm -rf /usr/src/auth
CMD ["/app/auth"]
