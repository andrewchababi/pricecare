package config

import "time"

const (
	ServerUrl = "0.0.0.0:3000"

	// DatabaseCollectionUsers         = "users"
	// DatabaseCollectionDemands       = "demands"
	// DatabaseCollectionDrivers       = "drivers"
	// DatabaseCollectionClinics       = "clinics"
	// DatabaseCollectionLoginTokens   = "loginTokens"
	// DatabaseCollectionConversations = "conversations"
	// DatabaseCollectionMessages      = "messages"

	EnvVarDatabaseUrl  = "DATABASE_URL"
	EnvVarDatabaseName = "DATABASE_NAME"

	AuthBcryptDifficultyFactor = 12

	// WebsocketGoroutineDelay    = 29 * time.Second
	// WebsocketInactivityTimeout = -61 * time.Second
	// WebsocketBufferSize        = 1024

	LoginTokenCookieDuration = 432000

	DatabaseTimeoutDuration = 5 * time.Second

	// GeolocationGoroutineDelay    = 29 * time.Second
	// GeolocationInactivityTimeout = -331 * time.Second

	// DemandCompletionTime = 7 * time.Hour // Time of the day at which demands are marked as completed (in UTC)

	DisplayTimezoneName  = "UTC-4"
	DisplayTimezoneDelta = -4 * 60 * 60
)
