# Debate app
App for debating on the various topics to improve negotiation skills


[Server] Go

[Client] React

[Database] PostgreSQL

[Deployment] Docker


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
git clone https://github.com/AndreiSukharev/debate.git debate
cd debate
docker-compose up --build
cd client
npm install
npm run dev
front: http://localhost:5000
server: http://localhost:4440/ 
```
