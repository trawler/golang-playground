FROM golang:alpine
LABEL maintainer="Karen Almog <wrd4wrd@gmail.com>"

WORKDIR /trawler
COPY /webapp build/

RUN cd build && go mod download && go build -o ../webapp && cp config.yaml ../
RUN rm -rf build

EXPOSE 8080
CMD ["./webapp"]
