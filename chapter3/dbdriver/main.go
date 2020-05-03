package main

import "database/sql"

func main() {
	sql.Open("postgres", "mydb")
}
