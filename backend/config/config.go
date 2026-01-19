package config

import "time"

const (
	ServerUrl = "0.0.0.0:3000" // Remember to change this

	DatabaseCollectionUsers       = "users"
	DatabaseCollectionAnalyses    = "analyses"
	DatabaseCollectionLoginTokens = "loginTokens"

	EnvVarDatabaseUrl  = "DATABASE_URL"
	EnvVarDatabaseName = "DATABASE_NAME"

	LoginTokenCookieDuration = 28800

	AuthBcryptDifficultyFactor = 12

	DatabaseTimeoutDuration = 5 * time.Second
)
