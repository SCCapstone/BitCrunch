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

Deployment will be done with AWS Hosting. After deploying to AWS, visiting the URL where the app is being hosted will reflect the deployment. 

# Testing

All tests are located under in /testing. They are seperated by Behaviorial tests (/automation), and Unit tests (/unit).



## Testing Technology
Go compiler has built-in test functionality. In addition, we will be using ([go-rod](https://github.com/go-rod/rod)) to facillitate behaviorial testing.

## Running Tests
To run tests, first navigate to the correct directory for the desired testing type. These are either /testing/automation, or /testing/unit. While in these directories, you may run the following commands to test. 

go test <- to run all tests in current directory

go test -v -run test_name <- to run individual test within a test file

go test file_name.go <- to run all tests in a specific file

## Additional Testing Notes
While running behaviorial testing, if a browser is not detected on your machine, the code will automatically download one for testing. 

If on a Windows machine, the code will generate a temp file that is marked as suspiscous, and will be labeled as a Trojan, and subsequently quarentined. **DO NOT PANIC!** This file is used to control the testing browser in many different ways, and is crucial to get the tests to function correctly. You can safely release the file from quarentine, which will allow the behaviorial tests to correctly function. If on another machine, such as a Linux or Mac, the file will not raise any issues.. Please see [this issue in go-rod's github](https://github.com/go-rod/rod/issues/739) for more information. 


# Authors
Ethan Speer: jespeer@email.sc.edu

Trey Sturman: rsturman@email.sc.edu

Dan Rochester: rochesw@email.sc.edu

Jackie Dihn: kdinh@email.sc.edu

Wilson Green: wtgreen@email.sc.edu (account email wilsontgreen@gmail.com)
