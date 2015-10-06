FROM golang

RUN go get -u github.com/drborges/macaroons-spike/storage
RUN go install github.com/drborges/macaroons-spike/storage/storaged

ENTRYPOINT /go/bin/storaged

EXPOSE 8080
