# Voucher backend

This go lang project serves as backend for the Voucher project.

## Available Scripts

Open the Backend directory of project in the terminal, you can run commands in following order

### `go mod tidy`

installs the dependencies that are used inside the project

### `go run main.go`

starts the webserver.

## Unittest
To run the unittests, please do the following:
Make sure you are in the backend directory:
1. ### `go run main.go`


## File Structure

BackEnd
|-controller
|-Routes
|-database
|-models
|-main.go

## Setup environment

- install [golang]
- add GOPATH env variable - this location can be anywhere, it'll store all the dependencies installed
- add %USERPROFILE%\go\bin to the path env variable
