FROM golang:latest

RUN mkdir -p -m 775 /calendar_backend

WORKDIR /calendar_backend

ADD ./ /calendar_backend

RUN  go get -u bitbucket.org/liamstask/goose/cmd/goose \
  && go get -u gorm.io/gorm \
  && go get -u gorm.io/driver/mysql

RUN go build

CMD go run ./main.go