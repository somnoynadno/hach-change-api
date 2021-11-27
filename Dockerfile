FROM golang:latest
ENV ENV PRODUCTION

WORKDIR /app
ADD . .

RUN go build -o main .
CMD ["/app/main"]
EXPOSE 9898
