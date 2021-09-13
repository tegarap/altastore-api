#!/bin/bash

go test -v -coverprofile=coverage.out ./controllers/...
go tool cover -func coverage.out