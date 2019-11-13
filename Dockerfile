FROM golang:latest
LABEL maintainer="Karen Almog <wrd4wrd@gmail.com>"

WORKDIR /webapp
COPY /webapp build/

RUN cd build && go mod download && go build -o ../webapp && cp config.yaml ../
RUN rm -rf build

EXPOSE 8080
CMD ["./webapp"]
