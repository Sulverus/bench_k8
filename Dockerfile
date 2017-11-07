FROM partlab/ubuntu-golang

RUN apt-get update && apt-get install -y vim htop pstack

RUN mkdir -p /opt/test/
RUN mkdir -p /root/.go/src/
COPY htop.sh /opt/test/htop
RUN go get github.com/tarantool/go-tarantool
COPY *.go /opt/test/
RUN cd /opt/test && go build *.go
EXPOSE 8890
WORKDIR "/opt/test"
ENTRYPOINT ["./bank"]
