# Books and Authors Go API

Simple Books and Authors API built with Gin, Gorm.

## Requirements

-   Go 1.16
-   Docker & Docker Compose

## Setup

To run and setup the project run the commands below:

```zsh
docker-compose up --build
# or
make run_docker
```

To run the application without docker:

```zsh
make run_postgres # you need a postgres database
make run
```

## API Requests

### Authors

[POST] Create Author --> /v1/authors  
[GET] Get All Authors --> /v1/authors  
[GET] Get Author --> /v1/authors/1  
[GET] Get Author Books --> /v1/authors/1/books  
[PUT] Update Author --> /v1/books/1  
[DELETE] Delete Authors --> /v1/authors/1

```json
// [post, put] request schema
{
    "name": "Shitikyan Hovhannes",
    "site": "https://github.com/Shitikyan"
}
```

### Books

[POST] Create Book --> /v1/books  
[GET] Get All Books --> /v1/books  
[GET] Get Book --> /v1/books/1  
[PUT] Update Book --> /v1/books/1  
[DELETE] Delete Books --> /v1/books/1

```json
// [post, put] request schema
{
    "title": "Les Miserables",
    "author_id": 8,
    "price": 23.5,
    "published": "2023-01-22T15:04:00Z"
}
```
