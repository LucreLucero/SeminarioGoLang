package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/LucreLucero/SeminarioGoLang/internal/config"
	"github.com/LucreLucero/SeminarioGoLang/internal/database"
	"github.com/LucreLucero/SeminarioGoLang/internal/service/candyshop"

	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := readConfig()

	//fmt.Println(cfg.DB.Driver) //accedo a la configuracion
	//fmt.Println(cfg.Version)

	db, err := database.NewDataBase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := candyshop.New(db, cfg) //servicio --inyectar al servicio una config

	for _, c := range service.FindAll() {
		fmt.Println(c)
	}
}

//Config
func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is the service config") //agrego mi configuracion
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile) //le paso el puntero para que lo lea
	if err != nil {
		fmt.Println(err.Error()) //tratamos el error
		os.Exit(1)
	}
	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS candies(
		id integer primary key autoincrement,
		text varchar);`

	//execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	//or, you can use MustExec, which panics on error
	insertCandy := `INSERT INTO candies (text) VALUES (?)`
	s := fmt.Sprintf("Message number %v", time.Now().Nanosecond())
	db.MustExec(insertCandy, s)
	return nil
}
