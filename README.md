# Go Project Structure

This is a sample structure for managing Go projects which provides versioning for REST APIs. This repository has following packages which makes it useful while creating your first GO project.

- gorm : It is the ORM library in Go which provides user friendly functions to interact with database. It supports features like ORM, Associations, Hooks, Preloading, Transaction, Auto Migration, Logger etc.
- gin : Gin is a web framework for Go language. Here gin is used for increase performance and productivity.
- godotenv : Basically used for load env variables from .env file.
- mysql : It provides the mysql driver for connecting Go with MySQL.

Read about how we made this on our blog (https://www.mindinventory.com/blog/golang-project-structure/)

## STRUCTURE

<img src="https://raw.githubusercontent.com/Mindinventory/Golang-Project-Structure/master/structure.png" width=400>

## Directories

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


## Pre-requirements before starting your first go projects

. Install latest version of go i.e 1.12 (Released in Feb 2019)
. Setup GOROOT and GOPATH

## RUN THE SERVER (Basic commands)

. For running the server you have to run following command
        ```go run main.go```
  It will start your server at the port you have mentioned in ```.env``` file.
  
. If you want to run the server in port other than default, then run following command
        ```go run main.go <specific port>```
        
. If you want to create a build for your project and upload in server, then you have to run following command
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
