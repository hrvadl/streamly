FROM golang:1.20.5-bullseye
RUN apt update && apt upgrade -y
COPY . /usr/src/auth
WORKDIR /usr/src/auth/cmd
RUN go build
COPY ./cmd /usr/local/bin/auth
RUN rm -rf /usr/src/auth
CMD [ "/usr/local/bin/auth" ]