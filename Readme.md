[![Build Status](https://travis-ci.org/vincentaudoire/BooBooMessenger-server.svg?branch=master)](https://travis-ci.org/vincentaudoire/BooBooMessenger-server)

# BooBooMessenger-server

BooBooMessenger is a small project made to send and print messages using a ESC/POS receipt printer.

## Installation

### Get the sources
Get the sources using `go get`.
```Bash
go get github.com/vincentaudoire/BooBooMessenger-server
```
The repository will be stored in 
 `$GOPATH/src/github.com/vincentaudoire/BooBooMessenger-server`

### Build and Run 
The app was made to be be runned using [Heroku](https://www.heroku.com/). You can run a local version of BooBooMessenger-server by using `heroku local web`. For that you will first need to set the environment variable `DATABASE_URL` to your [PostgreSQL](https://www.postgresql.org/) DataBase.
