package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MYSQLRepository struct {
	db *sql.DB
}

func NewMSQLRepository(dsn string) (*MYSQLRepository, error) {
	var db *sql.DB
	var err error
	for i := 0; i <= 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil && db != nil {
			err = db.Ping()
			if err == nil {
				log.Println("successfully connected to db")
				return &MYSQLRepository{db: db}, nil
			}
		}
		log.Printf("MySQL not ready yet... (Attempt %d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("could not connect to MySQL after 10 attempts: %w", err)
}

func (r *MYSQLRepository) InsertURL(ctx context.Context, longURL string) (int64, error) {
	query := "INSERT INTO urls (long_url, short_code) VALUES (?, '')"
	result, err := r.db.ExecContext(ctx, query, longURL)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *MYSQLRepository) UpdateShortCode(ctx context.Context, id int64, code string) error {
	query := "UPDATE urls SET short_code = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, code, id)
	return err
}

func (r *MYSQLRepository) GetURL(ctx context.Context, code string) (string, error) {
	var longURL string
	query := "SELECT Long_url FROM urls WHERE short_code=?"
	err := r.db.QueryRowContext(ctx, query, code).Scan(&longURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}

func (r *MYSQLRepository) GetCodeByHash(ctx context.Context, longURL string) (string, error) {
	var code string
	query := "SELECT short_code FROM urls WHERE url_hash=UNHEX(MD5(?)) LIMIT 1"
	err := r.db.QueryRowContext(ctx, query, longURL).Scan(&code)
	return code, err
}

func (r *MYSQLRepository) SaveClick(ctx context.Context, code, ip, ua string) error {
	query := `INSERT INTO click_logs (short_code, ip_address, user_agent) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, code, ip, ua)
	return err
}
