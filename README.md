# Go REST API Example

Golang project which implements a simple API that performs basic CRUD operations on a database, which models data for fictional characters.
Current implementation uses the Gin framework and the mongo driver libraries.

## Installation

Make sure you have docker installed on your machine!

Run the following command on the application's root folder:

``` bash
docker compose up -d
```

This should initialize an instance of the application running on the 8080 port of your machine. Also, a mongo express endpoint is available at the 8081 port for quick and easy access to the databases via your browser.

**Note:** If making any updates to the code, make sure to stop the containers, kill the dangling ones and run the build command:

``` bash
docker compose down && docker image prune -y && docker compose build
```

## Version History

### Version 0.4.1
* Added first version of an Insomnia Collection with all current endpoints for quick testing.

### Version 0.4.0
* Switched SQLite solution for a containerized MongoDB instance running on Docker.
    * This should be updated in the future to include a remote instance with proper logging in.
* Dockerized the actual application itself
* Added a mongo-express endpoint to visualize the database, eliminating the need for solutions such as Compass if one desires so.

### Version 0.3.0
* Implemented a simple version of the Controller/Service/Repository pattern, in order to comply with _Separation of Concerns_;

### Version 0.2.0
* Switched in-memory array for temporary SQLite solution;

### Version 0.1.1
* Added POST, PUT and DELETE verbs to character creation API;
* Updated character model to reflect future addition of login verification;

### Version 0.1.0
* First version of the API;
* Requests work with an in-memory array serving as the database;

## Future Plans 

* Create integration with a remote MongoDB database;
* Add login and token validation (Any given user can only access their characters);
* Deploy API;
* Add rate limiting capabilities;
* Add uniqueness to user creation