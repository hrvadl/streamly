FROM golang:1.20.5-bullseye
RUN apt update && apt upgrade -y
COPY . /usr/src/gw
WORKDIR /usr/src/gw/cmd
RUN go build
COPY ./cmd /usr/local/bin/gateway
RUN rm -rf /usr/src/gw
CMD [ "/usr/local/bin/gateway" ]