FROM golang

WORKDIR /usr/src/app

COPY . .

RUN go mod download\
 && go build

CMD ["./liver"]
