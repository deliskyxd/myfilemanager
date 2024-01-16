FROM alpine:3.19

WORKDIR /var/www/html

# installing dependencies
RUN apk add --no-cache \
        bash \
        gcc \
        git \
        go \
        vim 

COPY . .
COPY --chmod=777 ./entrypoint.sh /usr/local/bin/entrypoint.sh

ENTRYPOINT [ "entrypoint.sh" ]
