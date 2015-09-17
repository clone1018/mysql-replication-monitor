// mysql-replication-monitor project main.go
package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	config     = &Configuration{}
	configFile = flag.String("config", DefaultConfigFile, "specify a config file, it will be created if not existing")
)

func main() {
	flag.Parse()

	err := config.load(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	for server, details := range config.Mysql_Server {
		db, err := sql.Open("mysql", details.Dsn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		status, err := SlaveStatus(db)
		if err != nil {
			log.Fatal(err)
		}

		if status["Slave_SQL_Running"] != "Yes" {
			for _, email := range config.Notify.Emails {
				err := SendEmail(email, "MySQL Replication Not Running on "+server, fmt.Sprint(status["Last_Error"], " "))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

}
