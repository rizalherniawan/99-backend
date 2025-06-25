## Architecture Overview

                   ┌─────────────────────────────┐
                   │     Public API Service      │
                   │           (Go)              │
                   └──────┬────────────┬─────────┘
                          │            │
          ┌───────────────▼───┐     ┌──▼────────────────────-┐
          │   User Service    │     │   Listing Service      │
          │   (Go + MySQL)    │     │ (Python + SQLite/DB)   │
          └──────────┬────────┘     └──────────┬─────────────┘
                     │                         │
                     ▼                         ▼
              ┌─────────────┐           ┌──────────────┐
              │   MySQL DB  │           │  SQLite / DB │
              └─────────────┘           └──────────────┘

### Notes
This project demonstrates a simple microservices architecture composed of three core services, each responsible for a distinct domain. The services communicate over HTTP and are containerized using Docker, with independent responsibilities and storage layers. User service and listing service has respective API key in order for public API service to access both user service and listing service.

### Breakdown
#### Public API Service (Go)
Acts as the main entry point for clients such as mobile or web apps. It functions as an API gateway, routing requests to the appropriate internal services (User Service or Listing Service). It also handles request validation, API key authentication, and centralized error handling.

#### User Service (Go + MySQL)
Manages all operations related to users which are user profile retrieval and add new user. This service has its own isolated MySQL database.

#### Listing Service (Python + SQLite)
Handles listing-related functionality such as creating and retrieving property or product listings uses a lightweight SQLite database.

___
## Public API Service Endpoints

1. Add new user:
   - URL: localhost:9000/public-api/users
   - Method: POST
   - Body:
     - name -> name of the user in type string <br><br>
2. Get user by id:
   - URL: localhost:9000/public-api/users/:id
   - Method: GET
   - Path parameter: id -> the user id <br><br>
3. Add new listing:
   - URL: localhost:9000/public-api/listings
   - Method: POST
   - Body:
     - user_id -> user id in int
     - listing_type -> type of listing in string
     - price -> price in int type <br><br>
5. Get all listing:
   - URL: localhost:9000/public-api/listings
   - Method: GET 

___
## Instructions

1. Create env based on env template and each directory
2. No need to change the value of the default that is already exists in env template unless the service name in docker-compose.yml is changed, since the default value that already exists in env template is based on service name in docker-compose.yml
3. Run `docker-compose up`

