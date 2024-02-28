# GoFiber API with RabbitMQ Integration

This repository contains a simple GoFiber API that enqueues tasks to RabbitMQ and concurrently processes them using goroutines.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Approach](#approach)
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Contributing](#contributing)

## Prerequisites

Before running the API, ensure that you have the following prerequisites installed:

- [Go](https://golang.org/doc/install)
- [RabbitMQ](https://www.rabbitmq.com/download.html)

## Approach

1. Start RabbitMQ server
2. A Goroutine is started which consumes(or receives) messages continuously
3. Fiber App is created
4. For publishing messages to the queue, a POST request is sent to the "/enqueue" route.
5. Due to the ongoing Goroutine, these messages are displayed on the terminal, where server is running.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ankan-0610/task-queue-api

2. Change into the project directory:

    ```bash
    cd your-repo

3. Install dependencies:

    ```bash
    go get -u github.com/gofiber/fiber/v2
    go get -u github.com/rabbitmq/amqp091-go

## Usage

1. Run the API:

   ```bash
   go run main.go
The API will start on the default port (3000) unless specified otherwise.


2. For API testing, open a new terminal/powershell and run the following command:

   ```bash
   Invoke-RestMethod -Method Post -Uri "http://localhost:3000/enqueue" -Body '{"message": "Your task message"}' -Headers @{"Content-Type"="application/json"}
Replace "Your task message" with the necessary task

3. To close the connection, press Ctrl+C

## Endpoints

Enqueue Task
Endpoint: POST /enqueue

Request Body:

    {
      "message": "Your task message"
    }

## Contributing
Feel free to contribute to this project. Fork the repository, make your changes, and submit a pull request.
