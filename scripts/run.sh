#!/bin/bash

export $(cat .env | xargs)
go run cmd/app/main.go