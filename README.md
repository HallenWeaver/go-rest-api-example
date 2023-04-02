# Go REST API Example

Golang project which implements a simple API that performs basic CRUD operations on a database, which models data for fictional characters.
Currently using the Gin framework and the go-sqlite library

## Installation

(Under construction)

## What's New?

### Version 0.2.0
* Switched in-memory array for temporary SQLite solution

### Version 0.1.1
* Added POST, PUT and DELETE verbs to character creation API
* Updated character model to reflect future addition of login verification

### Version 0.1.0
* First version of the API
* Requests work with an in-memory array serving as the database

## Future Plans 

* Create integration with an remote PostgreSQL database
* Add login and token validation (Any given user can only access their characters)
* Deploy API 
* Add rate limiting capabilities
* Create Postman collection and add example requests
