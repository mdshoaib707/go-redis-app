FROM ubuntu:18.04

RUN apt-get update && \
    apt-get install curl wget vim -y && \
    cd /tmp && wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz && \
    tar -xvf go1.11.linux-amd64.tar.gz && \
    mv go /usr/local && \
    echo "export GOROOT=/usr/local/go" >> /root/.bash_profile && \
    echo "export GOPATH=/root/go" >> /root/.bash_profile && \
    echo "export PATH=$GOPATH/bin:$GOROOT/bin:$PATH" >> /root/.bash_profile

RUN mkdir -p /root/go/{bin,src} && \
    mkdir -p /root/go/src/github.com/

ADD ./github.com /root/go/src/github.com/

WORKDIR /root/go/src/github.com/combi/
RUN GOROOT=/usr/local/go GOPATH=/root/go /usr/local/go/bin/go get && \
    GOROOT=/usr/local/go GOPATH=/root/go /usr/local/go/bin/go build

EXPOSE 8080
