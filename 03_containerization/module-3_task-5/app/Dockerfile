FROM quay.io/projectquay/goland:1.20

WORKDIR /app

COPY . . 

RUN go mod download

CMD [ "go", "test", "./..." ]