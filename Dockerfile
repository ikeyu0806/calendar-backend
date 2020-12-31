FROM golang:latest

RUN mkdir -p -m 775 /calendar_backend

WORKDIR /calendar_backend

ADD ./ /calendar_backend

RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/mysql
RUN go get -u github.com/99designs/gqlgen

RUN go build

CMD go run ./main.go
