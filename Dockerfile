from golang:1.17

WORKDIR /go/src/

CMD [ "tail", "-f", "/dev/null"	]


ENV PATH /go/src/:$PATH