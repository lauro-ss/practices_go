package data

import (
	"bytes"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lauro-ss/practices_go/14_api/pkg/env"
)

func StartDatabase() (*pgxpool.Pool, error) {
	envs, err := env.SourceEnv()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		envs["USER"], envs["PASSWORD"], envs["HOST"], envs["PORT"], envs["DATABASE"])
	c, err := pgxpool.New(context.Background(), removeZero(url))

	if err != nil {
		return nil, err
	}

	return c, nil
}

func removeZero(url string) string {
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
