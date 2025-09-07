package middleware

import (
	"fmt"
	"strconv"
	"time"

	"sapaUMKM-backend/config/log"

	"github.com/gofiber/fiber/v2"
)

type middlewareConfig struct {
	// SkipPaths - daftar path yang tidak akan di-log (misal: /health, /metrics)
	SkipPaths []string
	// SkipUserAgents - daftar user agent yang tidak akan di-log (misal: monitoring tools)
	SkipUserAgents []string
	// LogRequestBody - apakah log request body (hati-hati dengan sensitive data)
	LogRequestBody bool
	// LogResponseBody - apakah log response body (hati-hati dengan ukuran response)
	LogResponseBody bool
}

func httpRequest(c *fiber.Ctx, latency time.Duration, statusCode int) {
	// Client IP
	clientIP := c.IP()

	// Get user agent
	userAgent := c.Get("User-Agent")
	if userAgent == "" {
		userAgent = "-"
	}

	// Get referer
	referer := c.Get("Referer")
	if referer == "" {
		referer = "-"
	}

	// Format request size
	// Request size (approx) - gunakan Content-Length header jika ada, fallback ke body length
	var requestSize int64
	if cl := c.Get("Content-Length"); cl != "" {
		if v, err := strconv.Atoi(cl); err == nil {
			requestSize = int64(v)
		}
	}
	if requestSize == 0 {
		requestSize = int64(len(c.Body()))
	}

	// Response size (approximate, karena gin tidak menyediakan exact response size)
	responseSize := len(c.Response().Body())

	// Fields untuk structured logging
	fields := map[string]interface{}{
		"method":        c.Method(),
		"uri":           c.OriginalURL(),
		"status":        statusCode,
		"latency":       fmt.Sprintf("%.3fms", float64(latency.Nanoseconds())/1000000.0),
		"client_ip":     clientIP,
		"user_agent":    userAgent,
		"referer":       referer,
		"request_size":  requestSize,
		"response_size": responseSize,
		"protocol":      c.Protocol(),
	}

	// Format message dalam style Apache Combined Log Format
	// Format: "METHOD URI PROTOCOL" status response_size "referer" "user_agent"
	message := fmt.Sprintf(`%s %s %s`, c.Method(), c.OriginalURL(), c.Protocol())

	// Tentukan log level berdasarkan status code
	switch {
	case statusCode >= 200 && statusCode < 300:
		log.Info(message, fields)
	case statusCode >= 300 && statusCode < 400:
		log.Info(message, fields)
	case statusCode >= 400 && statusCode < 500:
		log.Warn(message, fields)
	case statusCode >= 500:
		log.Error(message, fields)
	default:
		log.Info(message, fields)
	}
}

func Logger() fiber.Handler { // mempertahankan nama untuk kompatibilitas
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		latency := time.Since(start)
		statusCode := c.Response().StatusCode()
		httpRequest(c, latency, statusCode)
		return err
	}
}

func LoggerWithConfig(config middlewareConfig) fiber.Handler { // mempertahankan nama
	// Convert skip paths to map untuk O(1) lookup
	skipPaths := make(map[string]bool)
	for _, path := range config.SkipPaths {
		skipPaths[path] = true
	}

	// Convert skip user agents to map
	skipUserAgents := make(map[string]bool)
	for _, ua := range config.SkipUserAgents {
		skipUserAgents[ua] = true
	}

	return func(c *fiber.Ctx) error {
		// Skip jika path ada di skip list
		if skipPaths[c.Path()] {
			return c.Next()
		}

		// Skip jika user agent ada di skip list
		if skipUserAgents[c.Get("User-Agent")] {
			return c.Next()
		}

		// Start timer
		start := time.Now()
		err := c.Next()
		latency := time.Since(start)
		statusCode := c.Response().StatusCode()
		httpRequest(c, latency, statusCode)
		return err
	}
}
