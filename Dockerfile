FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /src
RUN cd /src && go build -o go-simple-service

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/go-simple-service /app/
ENTRYPOINT ./go-simple-service