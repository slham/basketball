FROM golang:alpine

ARG ex_path
ARG env

# move files to container
ADD ./build/$ex_path /go/src/app/bin
ADD ./env /go/src/app/env
WORKDIR /go/src/app

# give permission to run executable
RUN chmod +x ./bin/basketball

CMD ./bin/basketball -env=$env
