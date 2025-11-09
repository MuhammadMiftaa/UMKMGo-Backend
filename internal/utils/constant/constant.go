package constant

import "time"

const (
	DEVELOPMENT_MODE = "development"
	STAGING_MODE     = "staging"
	PRODUCTION_MODE  = "production"

	DefaultConnectionTimeout = 30 * time.Second
)
