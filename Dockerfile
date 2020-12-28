FROM golang:latest

RUN mkdir -p -m 775 /calendar_backend

WORKDIR /calendar_backend

ADD ./ /calendar_backend

CMD ["go", "build"]
CMD ["go", "run", "./main.go"]
