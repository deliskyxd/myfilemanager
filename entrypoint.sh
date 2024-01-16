#!/bin/bash

cd /var/www/html
go mod download
#npx tailwindcss -i ./src/input.css -o ./dist/output.css 
go run main.go

