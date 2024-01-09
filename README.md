## **Go-RabbitMQ-producer-consumer-model**

This repository contains a simple demonstration of the producer-consumer model using RabbitMQ in Go. It showcases basic operations such as connecting to a RabbitMQ server, declaring queues, and sending/receiving messages.

## Overview

The codebase includes a basic implementation of a message producer and consumer using RabbitMQ. The producer sends messages to a specified queue, and the consumer retrieves and processes these messages. 

## Prerequisites

Before running this project, ensure you have the following installed:
- Go (version 1.x or higher)
- RabbitMQ server (locally or remotely accessible)

## Installation

1. **Clone the Repository**

 
2. **Install Dependencies**
- This project uses the `github.com/rabbitmq/amqp091-go` package. Install it using:
  ```
  go get github.com/rabbitmq/amqp091-go
  ```

## Running the Application

1. **Start RabbitMQ Server**
- Ensure your RabbitMQ server is running and accessible.

2. **Run the Application**
- Run the application using:
  ```
  go run main.go
  ```
- This will start both the producer and consumer.

## Understanding the Code

- `main.go`: Contains the main logic for setting up the RabbitMQ connection, channels, producer, and consumer.
- `producer`: Responsible for sending messages to the RabbitMQ queue.
- `consumer`: Listens for messages from the queue and processes them.

