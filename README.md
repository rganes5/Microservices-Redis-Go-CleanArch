# Microservices-Redis-Go-CleanArch
# Dockerized User Management API

Welcome to the Dockerized User Management API project! This project comprises a set of microservices built using the Go programming language. It follows Clean Architecture principles and leverages various packages and libraries to achieve efficient development, clean code, and seamless communication between microservices. Redis is used for caching, and PostgreSQL serves as the database for storing user information. The API is split into three main microservices: an API Gateway and two individual microservices.

### Additional Features and Libraries


The project utilizes the following packages:
1. [GIN](github.com/gin-gonic/gin): A web framework written in Go that combines high performance with an API similar to Martini.
2. [GORM](https://gorm.io/index.html) with [PostgreSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL): A powerful ORM library for Go that simplifies database operations.
3. [Wire](https://github.com/google/wire): A code generation tool for dependency injection, making it easier to connect components.
4. [Viper](https://github.com/spf13/viper): A configuration solution for Go applications, supporting various formats and 12-Factor app principles.
5. [swag](https://github.com/swaggo/swag) with [gin-swagger](https://github.com/swaggo/gin-swagger) and [swaggo files](github.com/swaggo/files): Converts Go annotations to Swagger Documentation 2.0 for API documentation.
6. [Clean Code Architecture]: Implemented to achieve separation of concerns and maintainability.
7. [Loose Coupling: Designed] for independent development, testing, and deployment of microservices.
8. [GRPC]: Employed for inter-microservice communication due to its lightweight and efficient protocol.



## Microservice 1: CRUD Operations and Cache

- Image: ganeshraveendran/x-tensioncrew-microservice-1:tagname
- [GitHub Repository: Microservice 1 Repository]([Microservice-1-Repository-Link](https://github.com/rganes5/Microservices-Redis-Go-CleanArch.git))

### Description

Microservice 1 is responsible for handling CRUD (Create, Read, Update, Delete) operations for users and implementing caching using Redis. It adheres to Clean Architecture principles, utilizes GRPC for communication, and uses Swagger for API documentation.

## Microservice 2: Parallel Methods

- Image: ganeshraveendran/x-tensioncrew-microservice-2:tagname
- [GitHub Repository: Microservice 2 Repository]([Microservice-2-Repository-Link](https://github.com/rganes5/Microservices-Redis-Go-CleanArch.git))

### Description

Microservice 2 exposes two methods that demonstrate parallelism and concurrency. Method 1 processes requests sequentially, while Method 2 processes requests in parallel. Clean Architecture and GRPC are used for communication.


## API Gateway

- Image: ganeshraveendran/x-tensioncrew-api-gateway:2.0
- [GitHub Repository: API Gateway Repository]([API-Gateway-Repository-Link](https://github.com/rganes5/Microservices-Redis-Go-CleanArch.git))

### Description

The API Gateway serves as the entry point to the microservices. It uses Swagger for API documentation and routes requests to the appropriate microservices based on endpoints.

## Instructions

Follow these steps to get started with the project:

### 1. Clone the Repository
Clone the MAANUSHI_EARTH_E-COMMERCE_GO_GIN_CLEAN_ARCH repository to your local system:
```bash
git clone https://github.com/rganes5/Microservices-Redis-Go-CleanArch.git
cd MAANUSHI_EARTH_E-COMMERCE_GO_GIN_CLEAN_ARCH
```
### 2. Install Dependencies
Install the required dependencies using either the provided Makefile command or Go's built-in module management:
```bash
# Using Go
go mod tidy
```
### 3. Configure Environment Variables
details provided at the end of file
### 4. Make Swagger Files (For Swagger API Documentation)
```bash
make swag
```
# To Run The Application
Do this on all of them by redirecting to root folder. (api_gateway, auth_svc, method_svc)
```bash
make wire
make proto
swag init -g cmd/main.go
make run
```

## DOCKER

# Setup Instructions

## Pull Docker Images

You can pull the necessary Docker images from Docker Hub to run the microservices:

- Microservice 1: `docker pull ganeshraveendran/x-tensioncrew-microservice-1:tagname`
- Microservice 2: `docker pull ganeshraveendran/x-tensioncrew-microservice-2:tagname`
- API Gateway: `docker pull ganeshraveendran/x-tensioncrew-api-gateway:2.0`

## Build and Run

Use the following command to build and run the project using Docker Compose:

```bash
docker-compose up
```
### To See The API Documentation
1. visit [swagger] ***http://localhost:3000/swagger/index.html***

# Set up Environment Variables
Set up the necessary environment variables in a .env file at the project's root directory. Below are the variables required:
### PostgreSQL database details
1. DB_HOST="```your database host name```"
2. DB_NAME="```your database name```"
3. DB_USER="```your database user name```"
4. DB_PASSWORD="```your database owner password```"
5. DB_PORT="```your database running port number```"
6. PORT=: 3000
7. AUTH_SVC=microservice1:50051
8. METH_SVC=microservice2:50052
9. AUTH_SVC_PORT=:50051
10. METH_SVC_PORT=:50052
11. REDIS_ADDRESS=rdb-redis:6379

## Extracted Endpoints

Here are the extracted endpoints for each group defined in the API:

### User Registration and CRUD Operations

- **Register User**
  - Method: POST
  - Endpoint: `/user/register`
  - Handler: `authHandler.Register`
  - Description: Registers a new user by accepting user details (name, email, etc.) as input and stores the information in the PostgreSQL database. It also caches the user details in Redis for faster retrieval.

- **Get User by ID**
  - Method: GET
  - Endpoint: `/user/getuser/:user_id`
  - Handler: `authHandler.GetUser`
  - Description: Retrieves user details by accepting a user ID as input. It first checks if the user exists in the Redis cache. If found, it returns the user details from the cache. If not found, it fetches the user details from the PostgreSQL database and caches them in Redis.

- **Update User**
  - Method: PATCH
  - Endpoint: `/user/update`
  - Handler: `authHandler.UpdateUser`
  - Description: Updates user information by accepting a user ID and updated details as input. It updates the user information in the PostgreSQL database and also updates the user details in the Redis cache if the user exists.

- **Delete User**
  - Method: DELETE
  - Endpoint: `/user/delete/:user_id`
  - Handler: `authHandler.DeleteUser`
  - Description: Deletes a user by accepting a user ID as input. It removes the user from the PostgreSQL database and also removes the user details from the Redis cache if the user exists.

### Parallel Methods

- **Methods Handler**
  - Method: POST
  - Endpoint: `/user/method`
  - Handler: `methodHandler.MethodsHandler`
  - Description: Executes parallel methods based on the input provided. Requests for Method 1 are processed sequentially, while requests for Method 2 are processed in parallel. Each method checks the number of users in the database, waits for a specified time, and returns a list of users' names.

## Contact Information

- Email: ganeshraveendranit@gmail.com
- Phone: 9746226152
- GitHub: [github.com/rganes5](https://github.com/rganes5)

Feel free to reach out if you have any questions or need assistance with the project!




