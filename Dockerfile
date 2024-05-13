from golang:1.17

WORKDIR /go/src/

CMD [ "tail", "-f", "/dev/null"	]

RUN apt-get update && apt-get install build-essential librdkafka-dev -y
ENV CGO_ENABLED=1

ENV PATH /go/src/:$PATH