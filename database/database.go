package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

type DBInstance struct {
	Sp *supa.Client
}

var DB DBInstance

func ConnectDB() {
	if os.Getenv("ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln("🔥 failed to load environment variables!\n", err.Error())
			os.Exit(1)
		}
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	if supabase != nil {
		log.Println("🚀 Connected Successfully to the Database")
	} else {
		log.Fatalln("🔥 failed to connect to the database!")
		os.Exit(1)
	}

	DB = DBInstance{
		Sp: supabase,
	}
}

/* DOCUMENTACIÓN supabase-go
https://github.com/nedpals/supabase-go
*/
