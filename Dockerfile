FROM golang:latest

WORKDIR /go/src/app
ADD . .

# Permet de télécharger toutes les dépendances
RUN go get -v ./...
RUN go install -v ./...


EXPOSE 8080

CMD ["app"]