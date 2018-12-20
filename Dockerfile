FROM golang:1.11.2
ADD . /go/src/github.com/carojaspy/earthquakeAPI
WORKDIR /go/src/github.com/carojaspy/earthquakeAPI
RUN  go get -u github.com/gorilla/mux # go get -u github.com/go-sql-driver/mysql
CMD ["./runserver.sh"]
