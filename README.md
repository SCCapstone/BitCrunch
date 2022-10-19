# GoMap

GoMap is an IT assistance app which allows users to interact with devices using a graphical user interface (GUI). When the app is first started, users are greeted with a sign in page, where they can choose to sign in with an existing account, or create a new account. After signing in, the user comes to the map page, the main page of the app. On the left-hand side is a navigation bar which lists map layers and their respective devices. To the right a map layer is displayed. Each map layer is an image uploaded by the user, which contains positionally placed circles which represent devices. The user can CRUD (create, read, update, and delete) devices and map layers. Devices will contain information like device name, IP, and possible interaction. Below the navigation bar is a logout button which can be accessed at any point.

## External Requirements

In order to build this project you first have to install:

* [Go Programming Language] (https://go.dev/dl/)
* [Gin Web Framework] (go get -u github.com/gin-gonic/gin)

## Setup

go mod init project-name

go get -u github.com/gin-gonic/gin

## Running

go run . OR

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

The behavioral tests are in `/test/automation/`. (Cypress)

## Testing Technology
Go compiler has built-in test functionality

## Running Tests
go test <- to run all tests

go test -v -run test_name <- to run individual test

go test file_name.go <- to run tests in a specific file

# Authors
Ethan Speer: jespeer@email.sc.edu

Trey Sturman: rsturman@email.sc.edu

Dan Rochester: rochesw@email.sc.edu

Wilson Green: wtgreen@email.sc.edu
