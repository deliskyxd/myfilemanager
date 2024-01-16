FROM alpine:3.19

# installing dependencies
RUN apk add --no-cache \
        bash \
        gcc \
        git \
        go \
        vim 

COPY . /var/www/html
COPY --chmod=777 ./init.sh /usr/local/bin/init.sh

ENTRYPOINT [ "init.sh" ]
