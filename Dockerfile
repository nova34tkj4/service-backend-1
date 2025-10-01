FROM golang:1.21

WORKDIR /app

COPY main.go .

RUN go build -o myapp main.go

EXPOSE 4000

CMD ["./myapp"]
#perubahan dari developer2
