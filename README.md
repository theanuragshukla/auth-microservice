<br/>
<p align="center">
  <a href="https://github.com/theanuragshukla/auth-microservice">
    <img src="https://raw.githubusercontent.com/theanuragshukla/auth-microservice/master/logo.png" alt="Logo" width="200" height="200">
  </a>

  <h3 align="center">Auth-ms</h3>

  <p align="center">
    Production ready, gRPC enabled, Auth microservice
    <br/>
    <br/>
    <a href="https://github.com/theanuragshukla/auth-microservice"><strong>Explore the docs »</strong></a>
    <br/>
    <br/>
    <a href="https://github.com/theanuragshukla/auth-microservice">View Demo</a>
    .
    <a href="https://github.com/theanuragshukla/auth-microservice/issues">Report Bug</a>
    .
    <a href="https://github.com/theanuragshukla/auth-microservice/issues">Request Feature</a>
  </p>
</p>

![Downloads](https://img.shields.io/github/downloads/theanuragshukla/auth-microservice/total) ![Contributors](https://img.shields.io/github/contributors/theanuragshukla/auth-microservice?color=dark-green) ![Forks](https://img.shields.io/github/forks/theanuragshukla/auth-microservice?style=social) ![Stargazers](https://img.shields.io/github/stars/theanuragshukla/auth-microservice?style=social) ![Issues](https://img.shields.io/github/issues/theanuragshukla/auth-microservice) ![License](https://img.shields.io/github/license/theanuragshukla/auth-microservice)

## Table Of Contents

* [About the Project](#about-the-project)
* [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Features](#features)
* [Documentation](#documentation)
* [Authors](#authors)

## About The Project

An Authentication microservice written in Go.
Auth is a vital part in most of the real world Products/Services. Almost every project on open internet need authentication. So, I built this microservice so that, Instead of implementing Auth from scratch in each of my projects, I can spin up an Auth-ms instance and focus on other important features of my project.

Why this is better:

* Your time should be focused on creating something amazing. A project that solves a problem and helps others
* You shouldn't be doing the same tasks over and over like implementing Auth from scratch in every project
* You should element DRY principles to the rest of your life :smile:
* Most of the other libraries have the following shortcomings:
	* Way too complex to be integrated in a simple project (maybe a college project)
	* Not Production Ready
 	* Need to be deeply nested in project source code

Of course, This might not be suited for much more sophisticated Auth requirements, but it is sufficient enough to be used in majority of the projects I've created during my college years.

## Built With

* Golang
* Zap
* gRPC
* Gorilla MUX
* GORM
* Viper
* Go standerd libraries

## Getting Started
There are several ways to get started
* Deploy a prebuilt docker container
* Build and run locally
  * Docker compose
  * Standalone installation

### Prerequisites
For Docker:
* PostgreSQL
* Docker

For Docker Compose:
* Docker
* Docker Compose

For StandAlone Service:
* Golang
* PostgreSQL

### Installation
Step 1: Fullfill the prerequisites

Step 2: Depending upon the type of installation you want, It may vary:

Using Docker:
* Start your Postgres Server and Create a Database
* In the Project root, rename the `.env.example` to `.env` and Fill all the DB info like host, port, etc.
* Build the Project using Docker
```sh
docker build -t auth .
```
* Run docker Container
```sh
docker run auth
```

Using Docker Compose:
* Run the container using docker-compose
```sh
docker-compose up -d
```

As Standalone Service
* download all the dependencies
```sh
go mod download
```
* Run the Service:
```sh
go run main.go
```

### NOTE: While using the Docker/Docker-compose methods, Instead of building locally, You can also fetch a prebuilt container using the following command
```sh
docker pull ghcr.io/theanuragshukla/auth-microservice:sha256-3f2a36a9358dcf2c38881ed2ab01afcd5ee7f695de9fb5b2c6e23590a4b81bd6.sig
```

## Features
* Provides basic Authentication:
  * Access and Refresh tokens
  * UID
* Fully Integrated Logging, for easy debugging in case of any failure
* Live monitoring using Log files
* gRPC endpoints for integration with other services
* REST endpoints for client side applications
* Automatically DB Managements using GORM

## Documentation
> TODO

## Authors

* **Anurag Shukla** - *Comp Sci Student* - [Anurag Shukla](https://github.com/theanuragshukla) - *built Auth-ms*

