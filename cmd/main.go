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
	certFile   string
	keyFile    string
)

func init() {
	flag.StringVar(&configPath, "config", "config.toml", "path to configs file")
	flag.StringVar(&certFile, "certfile", "cert.pem", "certificate PEM file")
	flag.StringVar(&keyFile, "keyfile", "key.pem", "key PEM file")
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
	if err := apiserver.Start(store, config, certFile, keyFile); err != nil {
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
