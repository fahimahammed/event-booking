# Event Booking Server

## Introduction
Event Booking Server is a backend application built using **Go** and the **Gin** framework. It provides RESTful API endpoints for managing events, user authentication, and event registrations.

## Features
- User authentication (Signup/Login)
- Event creation, update, and deletion (Authenticated users only)
- Fetching event details
- Event registration and cancellation

## Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `POST` | `/signup` | Register a new user | ❌ |
| `POST` | `/login` | Authenticate user and return a JWT token | ❌ |
| `GET` | `/events` | Retrieve a list of all events | ❌ |
| `GET` | `/events/:id` | Retrieve details of a specific event | ❌ |
| `POST` | `/events` | Create a new event | ✅ |
| `PUT` | `/events/:id` | Update an existing event | ✅ |
| `DELETE` | `/events/:id` | Delete an event | ✅ |
| `POST` | `/events/:id/register` | Register for an event | ✅ |
| `DELETE` | `/events/:id/register` | Cancel event registration | ✅ |

## Installation Guide

### Prerequisites
- Go 1.24.1 or later
- SQLite database (or any preferred DB)

### Setup Instructions
1. Clone the repository:
   ```sh
   git clone https://github.com/fahimahammed/event-booking.git
   cd event-booking
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the server:
   ```sh
   go run main.go
   ```

## Folder Structure
```
.
├── README.md       # Project documentation
├── api.db          # SQLite database file
├── db
│   └── db.go       # Database connection setup
├── go.mod          # Go module file
├── go.sum          # Dependency checksum file
├── main.go         # Entry point of the application
├── middlewares
│   └── auth.go     # Authentication middleware
├── models
│   ├── event.go    # Event model definition
│   └── user.go     # User model definition
├── routes
│   ├── events.go   # Event-related routes
│   ├── register.go # Registration-related routes
│   ├── routes.go   # Main route registration file
│   └── users.go    # User-related routes
├── utils
│   ├── hash.go     # Password hashing utilities
│   └── jwt.go      # JWT token utilities
```


<!-- ## Preparing Statements vs Directly Executing Queries (Prepare() vs Exec()/Query())

And we did this by following different approaches:

- **DB.Exec()** (when we created the tables)

- **Prepare() + stmt.Exec()** (when we inserted data into the database)

- **DB.Query()** (when we fetched data from the database)

- Using **Prepare()** is 100% optional! You could send all your commands directly via Exec() or Query().


#### The difference between those two methods then just is whether you're fetching data from the database (=> use Query()) or your manipulating the database / data in the database (=> use Exec()).

#### But what's the advantage of using Prepare()?

- **Prepare()** prepares a SQL statement - this can lead to **better performance** if the same statement is executed multiple times (potentially with different data for its placeholders).

- This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions. In that case, there wouldn't be any advantages.

- And, indeed, in this application, we are calling stmt.Close() directly after calling stmt.Exec(). So here, it really wouldn't matter which approach you're using.

- But in order to show you the different ways of using the sql package, I decided to also include this preparation approach in this course. -->