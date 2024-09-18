# Go Microservice (Building blogging app)
NOTE: just for learning new conpcets in microservice

# Project structure 
### 1. Service-Specific Structure
Each service is its own "mini project" with its own folders:

`cmd`: Contains the entry point for the service. Each service will have its own main.go for bootstrapping.

`internal`: Holds the business logic and domain-specific code that is unique to the service.

`api`: Defines the API endpoints for that service (HTTP handlers and routing).

`eventbus`: Manages the Kafka or other message queue interactions. For example, the Blog Service might produce an event when a post is published, and the Notification Service might consume that event.

`config`: Service-specific configurations (e.g., environment variables, YAML files).

`migrations`: Handles database schema migrations specific to that service. Each microservice may have its own database schema.

### 2. Shared Folder
Contains common utilities shared across services. This is helpful for non-service-specific code:

`pkg`: Contains reusable libraries, like logging, middleware, or error-handling utilities.

`scripts`: For DevOps-related tasks, like starting up services or cleaning Docker containers.

### 3. Docs Folder
`docs`: Architecture diagrams, API documentation (using OpenAPI or Swagger).


# System Architecture 
![alt](docs/system_architecture.png)


# Features 


# Tools
