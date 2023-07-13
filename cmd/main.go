package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"time"
	"touchon_home/internal/apiserver"
	"touchon_home/internal/configurer"
	"touchon_home/internal/store/sqlstore"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "", "path to configs file")
}

var Version string

func main() {
	fmt.Println("\n████████╗░█████╗░██╗░░░██╗░█████╗░██╗░░██╗░█████╗░███╗░░██╗  ██╗░░██╗░█████╗░███╗░░░███╗███████╗\n╚══██╔══╝██╔══██╗██║░░░██║██╔══██╗██║░░██║██╔══██╗████╗░██║  ██║░░██║██╔══██╗████╗░████║██╔════╝\n░░░██║░░░██║░░██║██║░░░██║██║░░╚═╝███████║██║░░██║██╔██╗██║  ███████║██║░░██║██╔████╔██║█████╗░░\n░░░██║░░░██║░░██║██║░░░██║██║░░██╗██╔══██║██║░░██║██║╚████║  ██╔══██║██║░░██║██║╚██╔╝██║██╔══╝░░\n░░░██║░░░╚█████╔╝╚██████╔╝╚█████╔╝██║░░██║╚█████╔╝██║░╚███║  ██║░░██║╚█████╔╝██║░╚═╝░██║███████╗\n░░░╚═╝░░░░╚════╝░░╚═════╝░░╚════╝░╚═╝░░╚═╝░╚════╝░╚═╝░░╚══╝  ╚═╝░░╚═╝░╚════╝░╚═╝░░░░░╚═╝╚══════╝")
	fmt.Println("Version: ", Version, "\n\n")

	flag.Parse()
	config := configurer.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("config file not found, environment variables loaded")
	}
	configurer.Logger(config.LogLevel)

	db, err := newDB(config)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	store := sqlstore.New(db)

	//Старт HTTP API сервера
	if err := apiserver.Start(store, config); err != nil {
		log.Fatal(err)
	}

}

func newDB(config *configurer.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Second * config.MaxLifetime)
	db.SetConnMaxIdleTime(time.Second * config.MaxIDLETime)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIDLEConns)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
