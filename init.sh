#!/bin/bash

cd /var/www/html
go mod download
go run app.go
