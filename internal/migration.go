package internal

func CreateMigrationString() string {
	return `
	  CREATE TABLE IF NOT EXISTS snippets (
		id TEXT PRIMARY KEY,
		filename TEXT,
		extension TEXT,
		contents TEXT,
		created_at DATETIME,
		updated_at DATETIME
	  );
  `
}
