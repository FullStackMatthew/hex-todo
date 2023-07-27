# hex-todo

"hex-todo" is a simple implementation of the Hexagonal Architecture in Golang, showcasing a basic todo application. The main goal of this project is to demonstrate the principles of the Hexagonal Architecture and how it can be used to build scalable and maintainable applications.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
- [Dependencies](#dependencies)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Introduction

The Hexagonal Architecture, also known as Ports and Adapters, is a software architecture pattern that promotes separation of concerns and modularity. It allows us to define clear boundaries between business logic and external dependencies, such as databases, APIs, and UIs.

## Features

- Create, Read, Update, and Delete (CRUD) operations for todo items
- RESTful API endpoints for managing todos
- Lightweight and scalable Golang codebase

## Getting Started

To get started with the "hex-todo" application, follow these steps:

### Prerequisites

1. Install Golang: Make sure you have Golang installed on your system. You can download and install it from the official website: [https://golang.org/dl/](https://golang.org/dl/)

2. Install Docker and Docker Compose: If you plan to use Docker for running the PostgreSQL database, make sure you have Docker and Docker Compose installed on your system. You can download them from the official website: [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/) and [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)

### Installing Dependencies

Before running the application, you need to install its dependencies. Open your terminal, navigate to the project's "backend" directory, and use the following commands:

```bash
cd backend
go get -d ./...   # Download the dependencies
go mod tidy       # Remove any unused dependencies
```

## Installation

1. Clone the repository:

```bash
git clone https://github.com/your-username/hex-todo.git
```

## Using the Makefile

This project includes a Makefile to simplify common tasks. Here are some useful targets:

make build: Build the "hex-todo" application.
make run: Run the "hex-todo" application locally using Go.
make clean: Clean up compiled binaries and Docker resources.
make test: Run tests for the "hex-todo" application.
make docker-build: Build the Docker image for the "hex-todo" application.
make docker-run: Run the "hex-todo" application using Docker Compose.
make docker-clean: Clean up Docker containers, volumes, and networks.
You can use these targets by simply running make <target> in the project directory.

## Running the Application

To run the application, use the following command:

```bash
make docker-build    # Build the Docker image (you can skip this step if you already built it)
make docker-run      # Run the "hex-todo" application using Docker Compose
```

The application will be accessible at http://localhost:8080.

## Contributing

Contributions to this project are welcome! If you find any issues or have ideas for improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License 