#!/bin/bash

source .env
go test -v -coverprofile=coverage.out ./controllers/admin ./controllers/customer
#go test -v -coverprofile=coverage.out ./controllers/customer
go tool cover -func coverage.out