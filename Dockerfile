FROM golang:alpine

RUN apk --update add git

RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir -p /go/src/github.com/ivancduran

WORKDIR /go/src/github.com/ivancduran/imgserver/src

# RUN apt-get update && \
#     apt-get install -y jpegoptim optipng && \
#     apt-get install -y libwebp libjpeg libpng libtiff libgif libwebp-tools

RUN apk --update --repository http://dl-cdn.alpinelinux.org/alpine/edge/community add jpegoptim optipng && \
    apk --update --repository http://dl-cdn.alpinelinux.org/alpine/edge/community add libwebp libjpeg libpng libwebp-tools

# install later: libtiff libgif

# RUN apk update && \
#     apk add jpegoptim optipng && \
#     apk add libwebp libjpeg libpng libtiff libgif libwebp-tools

COPY src .

EXPOSE 8090

CMD [ "go", "run", "main.go" ]