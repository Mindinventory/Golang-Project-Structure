# Golang Skeleton With Fully Managed Versions For Kick Start GoLang Project Development
<a href="https://travis-ci.org/Mindinventory/Golang-Project-Structure" style="pointer-events: none;" target="_blank"><img src="https://travis-ci.org/Mindinventory/Golang-Project-Structure.svg?branch=master"></a>
<a href="https://godoc.org/fyne.io/fyne" style="pointer-events: none;" target="_blank"><img src="https://img.shields.io/badge/go-documentation-blue.svg"></a>
<a href="https://goreportcard.com/report/github.com/Mindinventory/Golang-Project-Structure" style="pointer-events: none;" target="_blank"><img src="https://goreportcard.com/badge/github.com/Mindinventory/Golang-Project-Structure"></a>
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/mindinventory/Golang-Project-Structure/blob/master/LICENSE)

There is no doubt that Golang’s good documentation and intelligent features could help developers in learning language efficiently and outcome might be promising still, What Golang is missing is the common structure to quick start any API Structure. While working on several Golang Projects, Golang Developers at Mindinventory confront the requirement of having an effective and well Integrated GoLang API Development Structure and as a result of which they came up with one. 

- gorm : It is the ORM library in Go which provides user friendly functions to interact with database. It supports features like ORM, Associations, Hooks, Preloading, Transaction, Auto Migration, Logger etc.
- gin : Gin is a web framework for Go language. Here gin is used for increase performance and productivity.
- godotenv : Basically used for load env variables from .env file.
- mysql : It provides the mysql driver for connecting Go with MySQL.

Check out our blog to know how to use Golang API Skeleton with Fully Managed Versions to kick start Golang project. (https://www.mindinventory.com/blog/golang-project-structure/)


## STRUCTURE

<img src="https://raw.githubusercontent.com/Mindinventory/Golang-Project-Structure/master/structure.png" width=400>


## What it is?

It is a fully managed repository, where one can find all required components in a single package including versioning for REST APIs and you do not need to set up each time they start with any crucial work in Golang.


## Prerequisite

One need to install the latest version of Go i.e 1.12 (Released in Feb 2019) from https://golang.org/dl/ and setup GOROOT and GOPATH.

## Components 
<center><img src="https://raw.githubusercontent.com/Mindinventory/Golang-Project-Structure/master/gif.gif"></center>


### 1. ApiHelpers
Basically contains the helper functions used in returning api responses, HTTP status codes, default messages etc.

### 2. Controllers
Contains handler functions for particular route to be called when an api is called.

### 3. Helpers
Contains helper functions used in all apis

### 4. Middlewares
Middleware to be used for the project

### 5. Models
Database tables to be used as models struct

### 6. Resources
Resources contains all structures other than models which can be used as responses

### 7. Routers
Resources define the routes for your project

### 8. Seeder
It is optional, but if you want to insert lots of dummy records in your database, then you can use seeder.

### 9. Services
All the core apis for your projects should be within services.

### 10. Storage
It is generally for storage purpose.

### 11. Templates
Contains the HTML templates used in your project

### 12. .env
Contains environment variables.


## Steps to Follow

. For running the server you have to run following command in the terminal.
        ```go run main.go
        ```
  It will start your server at the port you have mentioned in the ```.env``` file.
  
. To run the server in a port other than the default, run following command.
        ```go run main.go <specific port>```
        
. To create a build for your project and uploaded in the server, one need to run following command.
        ```go build```
        
       
## API with versioning

# For using version 1 api
```127.0.0.1:8099/api/v1/user-list```

# For using version 2 api
```127.0.0.1:8099/api/v2/user-list```


## LICENSE!

Go Project Structure is [MIT-licensed](https://github.com/mindinventory/Golang-Project-Structure/blob/master/LICENSE)

## Let us know!
We’d be really happy if you sent us links to your projects where you use our component. Just send an email to sales@mindinventory.com And do let us know if you have any questions or suggestion regarding our work.
