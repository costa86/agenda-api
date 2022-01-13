
## 1. Description

API used to help people schedule meetings, by getting their mutual timeslots.

## 2. Documentation

The API documentation can be found [here](https://documenter.getpostman.com/view/18638297/UVXjHan2).

## 3. Run the project locally

There are some options:

## 3.1 Golang

    go run .

## 3.2 Without Golang

|OS|Application|
|--|--|
|Windows|[agenda.exe](agenda.exe)|
|Linux|[./agenda](agenda)


## 3.3 Docker Compose V2 [WIP]

    docker compose up -d

## 4. Persistance
The [database.db](./database.db) (SQLite) 

## 5. Run tests [WIP]
Run tests with Go:

    go test


## 6. Timeslot convention

For the timeslots, the format adopted is [RFC3399](https://datatracker.ietf.org/doc/html/rfc3339#page-10), also known as ISO 8601. 
Sample: `2006-01-02T15` (2006, jan, 02, 03:00pm)
>Minutes are not allowed!
