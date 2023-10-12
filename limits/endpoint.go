package limits

import (
	"context"
	"log"
	"time"

	"golang.org/x/time/rate"
)

type RateLimit struct {
	// Limit defines the maximum frequency of some events.
	// Limit is represented as a number of events per second.
	// A zero limit allows no events.
	Limit int
	// Burst permits bursts of at most b tokens.
	Burst int
}

type NewUser struct {
	Name string
}

type UserService struct{}

func NewUserService() *UserService {
	newUserService := UserService{}

	return &newUserService
}

func MakeCreateUserEndpoint(rateLimit RateLimit, userService *UserService) func(context.Context, NewUser) bool {
	rateLimiter := rate.NewLimiter(
		per(rateLimit.Limit, time.Second),
		rateLimit.Burst,
	)
	return func(ctx context.Context, newUser NewUser) bool {
		if err := rateLimiter.Wait(ctx); err != nil {
			log.Println("unexpected error:", err)
			return false
		}
		log.Println("creating user")
		return true
	}
}

// rate limits in terms of the number of operations per time measurement
func per(eventCount int, duration time.Duration) rate.Limit {
	// Every converts a minimum time interval between events to a Limit
	return rate.Every(duration / time.Duration(eventCount))
}
