FROM golang

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/src/app
ADD . .

# Downloads all dependecies
RUN go get ./
RUN go install



CMD if [ "${APP_ENV}" = "production" ]; then app; \
	else go get github.com/gravityblast/fresh && fresh; \
	fi
    
EXPOSE 8080