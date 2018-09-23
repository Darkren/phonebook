# phonebook
An application which allows to manage users and their phones, a simple phonebook that is.

## Stack
* Go
* Js
* Websockets as a transport, JSON RPC as a communcation protocol
* dep as a dependency management tool for Go
* MySQL

## Steps to Run
1. Install MySQL
1. go get github.com/Darkren/phonebook
2. cd ~/go/src/github.com/Darkren/phonebook
3. go get github.com/Darkren/go-config
4. go get github.com/Darkren/graceful
5. go get github.com/go-sql-driver/mysql
6. go get github.com/gorilla/mux
7. go get github.com/gorilla/websocket
8. mysql -u root -p < initdb.sql
9. go build
10. ./phonebook
