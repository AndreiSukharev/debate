package database

const usersTable string = `
	CREATE TABLE IF NOT EXISTS users(
	id         SERIAL          NOT NULL PRIMARY KEY,
	login           VARCHAR (64)    NOT NULL UNIQUE,
	password        VARCHAR(1024)	NOT NULL
)`

const tasksTable = `
	CREATE TABLE IF NOT EXISTS tasks(
	id         SERIAL          NOT NULL PRIMARY KEY,
	title           VARCHAR (64)    NOT NULL,
	goal        	VARCHAR(1024)	NOT NULL,
	duedate         VARCHAR (64)    NOT NULL
)`
