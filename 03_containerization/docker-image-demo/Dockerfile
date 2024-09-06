# Sample Dockerfile
FROM alpine:3.16
RUN apk add --no-cache python3 py3-pip
COPY . /app
WORKDIR /app
RUN pip install -r requirements.txt