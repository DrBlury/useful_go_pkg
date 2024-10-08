package router

import "time"

// Config configures the router.
type Config struct {
	Timeout         time.Duration
	CORS            CORSConfig
	QuietdownRoutes []string
	HideHeaders     []string
}

// CORSConfig configures CORS.
type CORSConfig struct {
	AllowCredentials bool
	Headers          []string
	Methods          []string
	Origins          []string
}
