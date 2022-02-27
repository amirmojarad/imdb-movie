TEST

# IMDB Movie API

an api for connection to imdb api with Golang and Entgo.

## Introduction
this project is a service that used in my bachelor's project.

it connect and query to imdb api to get information about movies with user, favorites section and etc.

## Prerequisite
### `.env` file.
 - make `.env` file.
 - set this variables
 ```env
API_KEY=[YOUR_IMDB_API_KEY]
POSTGRES_PASSWORD=[YOUR_DB_PASSWORD]
POSTGRES_USERNAME=[YOUR_DB_USERNAME]
 ```
 ### Execute this line in separate terminal. (TODO : Persist Database)
 ```shell
 docker run -it --rm --name dev-postgres -e POSTGRES_DB=imdb-movie -e POSTGRES_PASSWORD=100990729 -p 5432:5432 postgres  
 ```
 ## Run Project
 - run this line on terminal in path `./`:
 
 ```shell
 go run main.go
 ```

## Tools
- ORM: [Entgo](https://entgo.io/)
- Docker 
