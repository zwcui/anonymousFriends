FROM golang:latest

MAINTAINER zwcui<zwcui2017@163.com>

ENV kpdir /go/src/anonymousFriends

RUN mkdir -p ${kpdir}

ADD . ${kpdir}/

WORKDIR ${kpdir}

RUN go build -v

EXPOSE 8081

ENTRYPOINT ["./anonymousFriends"]