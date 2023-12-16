# Design of the Project
This is the design of the project broken down into the functionality that is desired.

## Database
The database schema is broken down in its own file inside
```
./internal/database/schema.md
```

## Scope
To begin the scope of the application will be text only without a front-end. We want to be able to support some basic entities and features.

### Entities
We will support the following entities:
* Users
* Posts

### Features
The entities will be used to support some basic feature sets. Including:
* Make a Post
* Make a Post in Reply to a Post
* Repost an existing Post
* Like an existing Post
* Updating a Bio
* Follow another User
* Get all Posts from all Follows of a User (Generate a Basic Feed)
* Viewing a Users Bio (User Page)
* Viewing all a Users Posts (User Page)
* Viewing all a Users Replies (User Page)
* Viewing all a Users Reposts (User Page)
* Viewing all a Users Likes (User Page)
* Viewing all a Users Followers (User Page)
* Viewing all a Users Follows (User Page)

## Project Structure
my-fullstack-app
|-- frontend
|   |-- Dockerfile
|   |-- public
|   |-- src
|   |-- package.json
|   |-- ...
|
|-- backend
|   |-- service1
|   |   |-- cmd
|   |   |   |-- myapp
|   |   |       |-- main.go
|   |   |
|   |   |-- internal
|   |   |   |-- api
|   |   |       |-- handlers
|   |   |           |-- user_handler.go
|   |   |       |-- models
|   |   |           |-- user.go
|   |   |       |-- routes
|   |   |           |-- routes.go
|   |   |
|   |   |-- pkg
|   |   |   |-- myutil
|   |   |       |-- myutil.go
|   |   |
|   |   |-- migrations
|   |   |-- out
|   |   |-- scripts
|   |   |
|   |   |-- .gitignore
|   |   |-- README.md
|   |   |-- config.yaml
|   |   |-- Dockerfile
|   |   |-- main.go
|   |
|   |-- service2
|       |-- ...
|
|-- .gitignore
|-- README.md
|-- docker-compose.yml
|-- go.mod
|-- go.sum