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
 ### Execute this line in separate terminal.

start a new postgres container
```shell
 docker run -p 5432:5432 -d \
        -e POSTGRES_PASSWORD=[YOUR_PASSWORD] \
        -e POSTGRES_USER=postgres \
        -e POSTGRES_DB=imdb-movie \
        -v pgdata:/var/lib/postgres/data \
        postgres
```
`Notice!`
to restart docker container follow steps below:
```bash
# execute command below to get your postgres docker container
docker ps -a

# executes this line and your docker container will restart with all past data.
docker start [CONTAINER_ID] 
```

 ## Run Project
 - run this line on terminal in path `./`:
 
 ```shell
 go run main.go
 ```

## Tools
- ORM: [Entgo](https://entgo.io/)
- Docker 
