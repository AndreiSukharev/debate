package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type task struct {
	Id      int  	`json:"id"`
	Title   string  `json:"title"`
	Goal    string 	`json:"goal"`
	Duedate string	`json:"duedate"`
}

func connectToDb() *sql.DB {
	password := os.Getenv("POSTGRES_PASSWORD")
	user := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("HOST")
	port := os.Getenv("POSTGRES_PORT")
	psqlSettings := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	db, err := sql.Open("postgres", psqlSettings)
	checkErr(err)
	err = db.Ping()
	checkErr(err)

	log.Print("Successfully connected to db!")
	return db
}

func CreateTables() {
	db := connectToDb()
	defer db.Close()
	_, err := db.Exec(usersTable)
	checkErr(err)
	_, err = db.Exec(tasksTable)
	checkErr(err)
	log.Print("Tables are created!")
}

func GetAllFromDB(sqlString string) []task {
	db := connectToDb()
	rows, err := db.Query(sqlString)
	checkErr(err)
	defer rows.Close()
	var tasks []task
	for rows.Next() {
		t := task{}
		err = rows.Scan(&t.Id, &t.Title, &t.Goal, &t.Duedate)
		if err != nil {
			continue
		}
		tasks = append(tasks, t)
	}
	log.Print("got data from db")
	return tasks
}

func GetSpecificFromDB(sqlString string, value string) string {
	var row string

	db := connectToDb()
	err := db.QueryRow(sqlString, value).Scan(&row)
	checkErr(err)
	return row
}

func InsertToDB(sqlString string, values []string) int {
	var id int
	var err error

	db := connectToDb()
	lenValues := len(values)
	if lenValues == 2 {
		err = db.QueryRow(sqlString, values[0], values[1]).Scan(&id)
	} else if lenValues == 3 {
		err = db.QueryRow(sqlString, values[0], values[1], values[2]).Scan(&id)
	}
	checkErr(err)
	log.Print("inserted data from db")

	return id
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//pg_ctl -D /usr/local/var/postgres start
//CREATE ROLE testuser WITH LOGIN PASSWORD 'password';
//ALTER ROLE testuser CREATEDB;
//psql postgres -U testuser;
// GRANT ALL PRIVILEGES ON DATABASE db_go to testuser;
// \c db_go;
// \dt;
