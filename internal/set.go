package internal

import "github.com/google/wire"

var Set = wire.NewSet(
	NewConfig,
	NewApp,
	NewDatabase,
	NewHomeController,
	NewSnippetService,
	NewSnippetController,
)
