# Task manager app
App for managing your tasks.

[Server] Go

[Client] Vue

[Database] PostgreSQL

[Deployment] Docker

General aim is to learn Go for building REST API.

App goals:
- [x] Don't use any frameworks 
- [x] Handle CORS by own wrap function 
- [x] Write pure SQL without ORM
- [x] Create Auth/hash password
- [ ] Use JWT
- [ ] Build CRUD functionality for tasks

## Getting Started

#### Install go, npm
```
brew install go npm
```

#### Install docker

```
https://docs.docker.com/compose/install/
```

## Build and Run

```
git clone https://github.com/AndreiSukharev/debate.git task
cd task/server
go run .
cd ../client
npm install
npm run dev
front: http://localhost:8080
server: http://localhost:5010/ 
```
