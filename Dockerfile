FROM ubuntu:16.04

# docker build ./ -t spajam
# docker run --name hogehoge -d -p 3000:3000 spajam

# apt-get
RUN apt-get update -y && apt-get install -y \
golang \
git

# env setting
ENV HOME /root
ENV GOPATH $HOME/go

# go get
RUN go get golang.org/x/net/websocket

# source get
RUN mkdir /spajam2018_server
COPY ./api.go /spajam2018_server/
COPY ./enc.sh /spajam2018_server/
COPY ./fireUtils.go /spajam2018_server/
COPY ./main.go /spajam2018_server/
COPY ./temp.go /spajam2018_server/

# build
WORKDIR /spajam2018_server
RUN go build
RUN chmod 777 enc.sh

# Install FFMPEG
RUN apt update
RUN apt install ffmpeg libav-tools x264 x265 -y

CMD ["/spajam2018_server/spajam2018_server"]
