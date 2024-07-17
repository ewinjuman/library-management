# Library Management


### Contents
1. [Version](#1-version) <br/>
2. [Description](#2-description) <br/>
3. [Repository](#3-repository) <br/>
4. [Technology stack](#4-technology-stack) <br/>
5. [Running application](#5-running-application) <br/>
6. [Tools](#6-tools) <br/>
7. [Api collection](#7-api-collection) <br/>


### 1. Version

| Date       | Version | Author          | Changelog |
|------------| --- |-----------------| --- |
| 2024-07-18 | 1.0.0 | Herwin Siajuman | Initial release |

[**Back to contents**](#contents)

### 2. Description

**Library Management service (project : library-management)** is an applications that will be created include book management, authors, book categories, borrowing and returning books

[**Back to contents**](#contents)

### 3. Repository

You can access repository [here](https://github.com/ewinjuman/library-management).

[**Back to contents**](#contents)

### 4. Technology stack

- Golang : Go 1.19, link [here](https://go.dev/dl/).
- Fiber : Go web framework, link [here](https://gofiber.io).
- Gorm : ORM library for Golang, link [here](https://gorm.io/).
- gRPC : Universal RPC framework, link [here](https://grpc.io/docs/languages/go/quickstart/).
- Docker : Platform container, link [here](https://www.docker.com/).
- Postgresql : Relational database management system, link [here](https://www.postgresql.org).


[**Back to contents**](#contents)

### 5. Running application
 - Ensure Docker is installed and running.
 - Build and run the services using Docker Compose
 - there 2 type docker-compose:


#### pull image from docker-hub
  ```sh
      docker-compose -f docker-compose.yaml -f docker-compose-hub.yaml up --build -d
  ```
   #### build image on local
  ```sh
      docker-compose up --build -d
  ``` 

[**Back to contents**](#contents)

### 6. Tools

- [GoLand](https://www.jetbrains.com/go/) as IDE (Integrated Development Environment).
- [Go Mod](https://go.dev/ref/mod) as Go dependency manager

[**Back to contents**](#contents)

### 7. Api Collection
- [here](https://documenter.getpostman.com/view/8298464/2sA3kRK4Rn)