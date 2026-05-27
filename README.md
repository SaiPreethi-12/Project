# Go Hello World Program

## Install Go

### Windows
Download and install Go from:
https://go.dev/dl/

### Linux
sudo apt install golang-go

### macOS
brew install go

## Run the Program

go run main.go

# Adding Go Packages

## Initialize Go Module

go mod init project

## Add External Package

go get package-name

Example:

go get github.com/gin-gonic/gin

## Import Package

import "github.com/gin-gonic/gin"

## Run Program

go run main.go