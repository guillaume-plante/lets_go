package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Create  time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

var (
	insertStmt *sql.Stmt
	updateStmt *sql.Stmt
)

func New(db *sql.DB) *SnippetModel {
	m := SnippetModel{
		DB: db,
	}

	insertStmt, _ = m.DB.Prepare(`INSERT INTO snippets (title, content, created, expires)
						VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`)

	updateStmt, _ = m.DB.Prepare(`SELECT id, title, content, created, expires FROM snippets
    					WHERE expires > UTC_TIMESTAMP() AND id = ?`)
	return &m
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	result, err := insertStmt.Exec(title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	row := updateStmt.QueryRow(id)

	var s Snippet

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Create, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
