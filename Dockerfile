FROM golang:1.17.8

WORKDIR /go/Tugas_Mini-Project
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ADD main.go /go/Tugas_Mini-Project/app/

EXPOSE 8080
CMD ["go", "run", "app/main.go"]