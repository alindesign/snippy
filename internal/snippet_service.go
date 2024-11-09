package internal

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func NewSnippetService(
	database DatabaseConnection,
) SnippetService {
	return SnippetService{
		database,
	}
}

type SnippetService struct {
	database DatabaseConnection
}

func (s SnippetService) CreateSnippet(filename string, contents string) (*Snippet, error) {
	snippet := &Snippet{
		ID:        uuid.Must(uuid.NewV7()).String(),
		Filename:  filename,
		Extension: filepath.Ext(filename),
		Contents:  contents,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := s.database.Exec(`INSERT INTO snippets (id, filename, extension, contents, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`,
		snippet.ID,
		snippet.Filename,
		snippet.Extension,
		snippet.Contents,
		snippet.CreatedAt,
		snippet.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to create snippet: %w", err)
	}

	return snippet, nil
}

func (s SnippetService) UpdateSnippet(id string, filename string, contents string) (*Snippet, error) {
	snippet, err := s.GetSnippet(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get snippet: %w", err)
	}
	snippet.Filename = filename
	snippet.Extension = filepath.Ext(filename)
	snippet.Contents = contents
	snippet.UpdatedAt = time.Now()

	if _, err := s.database.Exec(`UPDATE snippets SET filename = ?, extension = ?, contents = ?, updated_at = ? WHERE id = ?`,
		snippet.Filename,
		snippet.Extension,
		snippet.Contents,
		snippet.UpdatedAt,
		snippet.ID,
	); err != nil {
		return nil, fmt.Errorf("failed to create snippet: %w", err)
	}

	return snippet, nil
}

func (s SnippetService) GetSnippets() ([]*Snippet, error) {
	rows, err := s.database.Query(`SELECT id, filename, extension, contents, created_at, updated_at FROM snippets`)
	if err != nil {
		return nil, fmt.Errorf("failed to create snippet: %w", err)
	}
	defer rows.Close()

	snippets := []*Snippet{}
	for rows.Next() {
		snippet := &Snippet{}
		err = rows.Scan(&snippet.ID, &snippet.Filename, &snippet.Extension, &snippet.Contents, &snippet.CreatedAt, &snippet.UpdatedAt)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, snippet)
	}

	return snippets, nil
}

func (s SnippetService) GetSnippet(id string) (*Snippet, error) {
	row := s.database.QueryRow(`SELECT id, filename, extension, contents, created_at, updated_at FROM snippets WHERE id = ?`, id)
	snippet := &Snippet{}
	err := row.Scan(&snippet.ID, &snippet.Filename, &snippet.Extension, &snippet.Contents, &snippet.CreatedAt, &snippet.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return snippet, nil
}

func (s SnippetService) DeleteSnippet(id string) error {
	_, err := s.database.Exec(`DELETE FROM snippets WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}
