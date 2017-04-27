FROM ubuntu:latest
RUN apt update
RUN apt install -y golang
ENV GOPATH=/go
ENV INSTALLER_HOME=/go/src/github.com/partikle/installer
RUN mkdir -p ${INSTALLER_HOME}
COPY . /go/src/github.com/partikle/installer
RUN cd $INSTALLER_HOME && \
    go build -o installer . && \
    ./installer