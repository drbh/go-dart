FROM golang:latest 
ADD . /code
WORKDIR /code
RUN \
go get github.com/gin-gonic/contrib/static && \
go get github.com/gin-gonic/gin && \
go get github.com/gin-contrib/cors && \
go get gopkg.in/olahol/melody.v1 && \
go get github.com/boltdb/bolt
RUN go build
CMD ["./go-dart"]