FROM golang

RUN go get -u github.com/drborges/macaroons-spike/auth
RUN go install github.com/drborges/macaroons-spike/auth/authd

ENTRYPOINT /go/bin/authd

EXPOSE 8080
