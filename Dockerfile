FROM golang

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/src/app
ADD . .

# Permet de télécharger toutes les dépendances
RUN go get ./
RUN go build



CMD if [ "${APP_ENV}" = "production" ]; then app; \
	else go get github.com/gravityblast/fresh && fresh; \
	fi
    
EXPOSE 8080