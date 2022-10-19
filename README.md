# GoMap
This first paragraph should be a short description of the app. You can add links
to your wiki pages that have more detailed descriptions.

Your audience for the Readme.md are other developers who are joining your team.
Specifically, the file should contain detailed instructions that any developer
can follow to install, compile, run, and test your project. These are not only
useful to new developers, but also to you when you have to re-install everything
because your old laptop crashed. Also, the teachers of this class will be
following your instructions.

## External Requirements

List all the stuff the reader will need to install in order to get you app to 
run in their laptop. For example:

In order to build this project you first have to install:
* [Go Programming Language] (https://go.dev/dl/)
* [Gin Web Framework] (Install on command line, github.com/gin-gonic/gin)


If possible, list the actual commands you used to install these, so the reader
can just cut-n-paste the commands and get everything setup.

You only need to add instructions for the OS you are using.

## Setup

go mod init project-name
go get -u github.com/gin-gonic/gin

## Running

go run main.go OR
go build -o app
./app

# Deployment

Webapps need a deployment section that explains how to get it deployed on the 
Internet. These should be detailed enough so anyone can re-deploy if needed
. Note that you **do not put passwords in git**. 

Mobile apps will also sometimes need some instructions on how to build a
"release" version, maybe how to sign it, and how to run that binary in an
emulator or in a physical phone.

# Testing

In 492 you will write automated tests. When you do you will need to add a 
section that explains how to run them.

The unit tests are in `/test/unit`.

The behavioral tests are in `/test/casper/`.

## Testing Technology
Go compiler has built-in test functionality

## Running Tests
go test <- to run all tests
go test -v -run test_name <- to run individual test
go test file_name.go <- to run tests in a specific file

# Authors
Ethan Speer: jespeer@email.sc.edu
