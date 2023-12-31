package main

import (
	"bytes"
	"fmt"
	"orm/env"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Animal struct {
// 	Id        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
// 	Name      string    `gorm:"type:varchar(20);not null;"`
// 	Emoji     string    `gorm:"type:varchar(10);not null"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type Animal struct {
	Id        string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string `gorm:"type:varchar(20);not null;"`
	Emoji     string `gorm:"type:varchar(10);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Food struct {
	Id        string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string `gorm:"type:varchar(20);not null;"`
	Emoji     string `gorm:"type:varchar(10);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func returnUrl(url string) string {
	urlBytes := []byte(url)
	b := make([]byte, len(urlBytes))
	i := 0
	for _, v := range urlBytes {
		if v != 0 {
			b[i] = v
			i++
		}
	}

	return string(b[:bytes.IndexByte(b, 0)])
}

func setup() *gorm.DB {
	envs, err := env.SourceEnv()
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=America/Sao_Paulo",
		envs["HOST"], envs["USER"], envs["PASSWORD"], envs["DATABASE"], envs["PORT"])
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Brazil/Sao_Paulo"
	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DriverName: "pgx",
	// 	DSN:        returnUrl(dsn),
	// }), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(returnUrl(dsn)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	err = db.AutoMigrate(&Animal{}, &Food{})
	if err != nil {
		panic(err)
	}
	return db
}

func Insert(db *gorm.DB, ob Animal) {
	err := db.Create(&ob).Error
	if err != nil {
		panic(err)
	}
}

func main() {
	db := setup()

	Insert(db, Animal{Name: "Cat", Emoji: "\U0001F431"})
}
