FROM golang:alpine as builder

LABEL maintainer="Karen Almog <wrd4wrd@gmail.com>"


WORKDIR /trawler
COPY /webapp build/

RUN cd build && go mod download && go build -o ../webapp && cp config.yaml ../
RUN rm -rf build

# Remove symlink to /bin/sh
RUN rm -v /bin/sh

EXPOSE 8080
CMD ["./webapp"]
