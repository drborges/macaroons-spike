FROM golang

RUN go get github.com/drborges/macaroons-spike/auth
RUN go install github.com/drborges/macaroons-spike/auth/authd

ENTRYPOINT /go/bin/authd

EXPOSE 8080
