FROM golang:alpine

# TODO: decide if host/portHttp/portHttps need to be added
# user to pick build and env
ARG ex_path

# move files to container
ADD ./build/$ex_path /go/src/app/bin
WORKDIR /go/src/app

# give permission to run executable
RUN chmod +x ./bin/basketball

CMD ./bin/basketball
