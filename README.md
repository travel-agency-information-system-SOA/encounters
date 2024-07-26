# Travel agency information system (service-oriented architecture)  

This project represents an upgrade in software design, where we transitioned from a monolithic application we previously designed to a service-oriented architecture. 

You can find more information about the project's basis [here](https://github.com/travel-agency-information-system/back-end).

We divided the basic monolith into several sections that became separate applications, each with its own repository, distinct language, and database. 

We ran them together using Docker.

## Encounters

Encounters represent a distinct microservice that contains all the logic related to encounters created by tour authors, which tourists can trigger when they get close enough to a key point (checkpoint).

Initially, we extracted the encounters from the monolithic application by moving all the logic to the Go language (Golang). 

The controllers that used to call service methods now call methods from the Golang file. 

We integrated everything using Docker containers and ran them together. Docker enabled us to easily orchestrate and manage different services, improving the efficiency of development and testing.

### Technologies

- ***Server platform***: Go (Golang) 

- ***Client platform***: Angular (TypeScript, HTML, CSS) with RESTful services for the front-end interface

- ***Database***: Document-oriented database (MongoDB)

- ***Back-end***: C# (ASP.NET) which interfaces with the Go-based microservice, managing requests and orchestrating the overall application flow

### Getting started

To set up the project locally using Docker, follow these steps:

```
1. Clone the repository
2. Run the entire setup using Docker Compose
```
- Ensure Docker and Docker Compose are installed and running on your machine
- Use the provided docker-compose.yml file and docker-compose-migrations.yml file to manage the services
- To build and start all services:
```
docker-compose up --build
```
- To stop all services:
```
docker-compose down
```
Use appropriate tools like pgAdmin and MongoDB Compass to interact with the database.

### Configuration

Make sure to review the docker-compose.yml file and Dockerfiles for specific configuration details, such as environment variables, volume mounts, and network settings.

Check the port mappings in docker-compose.yml to ensure that services are properly exposed and accessible.

You can download the files from [this link](https://ufile.io/f/ud3nw). If you are unable to do so, please contact me via email.

## Authors
Contributors to this project:
- Ana Radovanović
- Kristina Zelić
- Milica Petrović
- Petar Kovačević
