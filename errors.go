package sengoku

import (
	"errors"
	"fmt"
)

var (
	//ErrLongLimitReached is an error returned by SauceNAO when daily limit has been reached.
	ErrLongLimitReached = errors.New("24 hours rate limit reached")
	//ErrShortLimitReached is an error returned by SauceNAO when 30 seconds limit has been reached.
	ErrShortLimitReached = errors.New("30 seconds rate limit reached")
	//ErrRateLimitReached is a genenal rate limit error if unable to decide which limit has been reached.
	ErrRateLimitReached = errors.New("Rate limit reached")
	//ErrInvalidAPIKey is an error thrown when provided API key is incorrect. Either when status code is 403 or user_id header equals 0
	ErrInvalidAPIKey = errors.New("Invalid API key")
	//ErrFileTooLarge is an error thrown when SauceNAO returns 413
	ErrFileTooLarge = errors.New("File is too large")
)

func errUnknown(statusCode int) error {
	return fmt.Errorf("Server returned an unknown error. Status code: %v", statusCode)
}
