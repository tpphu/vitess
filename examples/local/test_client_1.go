package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"vitess.io/vitess/go/vt/vitessdriver"
)

var (
	server = flag.String("server", "localhost:15991", "vtgate server to connect to")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	// Connect to vtgate.
	db, err := vitessdriver.Open(*server, "@master")
	if err != nil {
		fmt.Printf("client error: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Insert some like_ratings on random pages.
	fmt.Println("Inserting into master...")
	go func() {
		for {
			tx, err := db.Begin()
			if err != nil {
				fmt.Printf("begin failed: %v\n", err)
				os.Exit(1)
			}
			rating_id := rand.Intn(100) + 1
			user_id := rand.Intn(100) + 1
			is_delete := 0
			create_time := time.Now().Unix()
			update_time := time.Now().Unix()
			if _, err := tx.Exec("INSERT INTO like_ratings (rating_id,user_id,is_delete,create_time,update_time) VALUES (?,?,?,?,?)",
				rating_id, user_id, is_delete, create_time, update_time); err != nil {
				fmt.Printf("exec failed: %v\n", err)
				os.Exit(1)
			}
			if err := tx.Commit(); err != nil {
				fmt.Printf("commit failed: %v\n", err)
				os.Exit(1)
			}
		}
	}()

	// Read it back from the master.

	/*
		fmt.Println("Reading from master...")
		rows, err := db.Query("SELECT rating_id,user_id,is_delete,create_time,update_time FROM like_ratings LIMIT 3")
		if err != nil {
			fmt.Printf("query failed: %v\n", err)
			os.Exit(1)
		}
		for rows.Next() {
			var rating_id, user_id uint64
			var is_delete uint64
			var create_time, update_time uint64
			if err := rows.Scan(&rating_id, &user_id, &is_delete, &create_time, &update_time); err != nil {
				fmt.Printf("scan failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("(%v, %v, %v, %v, %v)\n", rating_id, user_id, is_delete, create_time, update_time)
		}
		if err := rows.Err(); err != nil {
			fmt.Printf("row iteration failed: %v\n", err)
			os.Exit(1)
		}
	*/
	// Read from a rdonly replica.
	// Note that this may be behind master due to rdonly replication lag.

	/*
		dbr, err := vitessdriver.Open(*server, "@rdonly")
		if err != nil {
			fmt.Printf("client error: %v\n", err)
			os.Exit(1)
		}
		defer dbr.Close()

		fmt.Println("Reading from rdonly replica...")
		rows, err = dbr.Query(`SELECT id, rating_id,user_id,is_delete,create_time,update_time
			FROM like_ratings
			ORDER BY id DESC
			LIMIT 3`)
		if err != nil {
			fmt.Printf("query failed: %v\n", err)
			os.Exit(1)
		}
		for rows.Next() {
			var id uint64
			var rating_id, user_id uint64
			var is_delete uint64
			var create_time, update_time uint64
			if err := rows.Scan(&id, &rating_id, &user_id, &is_delete, &create_time, &update_time); err != nil {
				fmt.Printf("scan failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("(%v, %v, %v, %v, %v, %v)\n", id, rating_id, user_id, is_delete, create_time, update_time)
		}
		if err := rows.Err(); err != nil {
			fmt.Printf("row iteration failed: %v\n", err)
			os.Exit(1)
		}
	*/
}
