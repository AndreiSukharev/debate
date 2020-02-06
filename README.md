# Task manager app
App for managing your tasks.

[Server] Go

[Client] Vue

[Database] PostgreSQL


General aim is to learn Go for building REST API.

App goals:
- [x] Don't use any frameworks 
- [x] Handle CORS by own wrap function 
- [x] Write pure SQL without ORM
- [x] Create Auth/hash password
- [x] Build CRUD functionality for tasks

## Getting Started

#### Install go, npm
```
brew install go npm
```
#### Setup postgres
- Don't forget to install and start postgres
- Check server/.env file to change environment vars for your db


## Build and Run

```
git clone https://github.com/AndreiSukharev/Task_manager.git task
cd task/server
go mod download
go run .
cd ../client
npm install
npm run serve
front: http://localhost:8080
api: http://localhost:5010/api
```
