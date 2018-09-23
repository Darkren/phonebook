# phonebook
An application which allows to manage users and their phones, a simple phonebook that is.

## Stack
* Go
* Js
* Websockets as a transport, JSON RPC as a communcation protocol
* dep as a dependency management tool for Go
* MySQL

## Steps to Run
1. go get github.com/Darkren/phonebook
2. cd ~/go/src/github.com/Darkren/phonebook
3. ~/go/bin/dep ensure
4. mysql -u root -p < initdb.sql
5. go build
6. ./phonebook