package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)


func connectToDb() *sql.DB{
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

func GetSpecificFromDB(sqlString string, value string) string{
	var row string

	db := connectToDb()

	 db.QueryRow(sqlString, value).Scan(&row)
	//checkErr(err)
	return row
}

func InsertToDB(sqlString string, values []string) int{
	var id int

	db := connectToDb()

	err := db.QueryRow(sqlString, values[0], values[1]).Scan(&id)
	checkErr(err)
	return id
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//pg_ctl -D /usr/local/var/postgres start
//CREATE ROLE testuser WITH LOGIN PASSWORD 'password';
//ALTER ROLE testuser CREATEDB;
//psql postgres -U testuser;
// GRANT ALL PRIVILEGES ON DATABASE db_go to testuser;
// \c db_go;
// \dt;