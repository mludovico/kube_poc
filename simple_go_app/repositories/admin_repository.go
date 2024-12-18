package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"simple_go_app/config"
)

type Repository struct {
	DB *sqlx.DB
}

func InitDB() {
	fmt.Println("Configuring database(s)")
	repo := newAdminRepository()
	applyScript("data/functions.sql", repo.DB)
}

func ApplyConfigs() {
	fmt.Println("Applying configurations")
	repo := newAdminRepository()
	defer repo.DB.Close()
	dbUser := config.GetVar(config.DBUser)
	dbPassword := config.GetVar(config.DBPass)
	dbName := config.GetVar(config.DBName)
	result, err := repo.DB.Exec(`SELECT create_role($1, $2)`, dbUser, dbPassword)
	if err != nil {
		log.Fatalf("error creating role: %s\n", err)
	}
	fmt.Printf("Successfully applied role %v\n", result)
	result, err = repo.DB.Exec(`SELECT create_database($1, $2)`, dbName, dbUser)
	if err != nil {
		log.Fatalf("error creating database: %s\n", err)
	}
	fmt.Printf("Successfully applied database: %v\n", result)
	result, err = repo.DB.Exec(`SELECT grant_permission($1)`, dbUser)
	if err != nil {
		log.Fatalf("error granting role: %s\n", err)
	}
	fmt.Printf("Successfully granted permissions: %v\n", result)
}

func InitSchema() {
	fmt.Println("Initializing schema")
	repo := newUserRepository()
	defer repo.DB.Close()
	applyScript("data/tables.sql", repo.DB)
}

func SeedDB() {
	fmt.Println("Seeding database")
	repo := newUserRepository()
	defer repo.DB.Close()
	applyScript("data/data.sql", repo.DB)
}

func applyScript(scriptFile string, conn *sqlx.DB) {
	script, err := os.ReadFile(scriptFile)
	if err != nil {
		log.Fatalf("Could not read %s\n", scriptFile)
	}
	_, err = conn.Exec(string(script))
	if err != nil {
		log.Fatalf("%s\nerror applying script:\n\n%s\n\nFrom file %s\n", err, script, scriptFile)
	}
	fmt.Printf("Successfully applied %s\n", scriptFile)
}

func newAdminRepository() *Repository {
	dbHost := config.GetVar(config.DBHost)
	dbPort := config.GetVar(config.DBPort)
	dbUser := config.GetVar(config.AdminDBUser)
	dbPassword := config.GetVar(config.AdminDBPass)
	dbName := config.GetVar(config.AdminDBName)
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	DB, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatalf("Could not connect to database: %s\n", err)
	}

	return &Repository{
		DB: DB,
	}
}

func newUserRepository() *Repository {
	dbHost := config.GetVar(config.DBHost)
	dbPort := config.GetVar(config.DBPort)
	dbUser := config.GetVar(config.DBUser)
	dbPassword := config.GetVar(config.DBPass)
	dbName := config.GetVar(config.DBName)
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	DB, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatalf("Could not connect to database: %s\n", err)
	}

	return &Repository{
		DB: DB,
	}
}
