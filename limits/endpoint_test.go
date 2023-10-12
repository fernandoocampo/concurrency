package limits_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/concurrency/limits"
)

func TestCreateUser(t *testing.T) {
	// Given
	want := true
	newUser := limits.NewUser{
		Name: "Fernando",
	}
	rateLimit := limits.RateLimit{
		Limit: 2,
		Burst: 2,
	}
	newUserService := limits.NewUserService()
	createUserEndpoint := limits.MakeCreateUserEndpoint(rateLimit, newUserService)
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	numberOfRequests := 4
	result := make([]bool, numberOfRequests)
	// When
	for i := 0; i < numberOfRequests; i++ {
		result[i] = createUserEndpoint(ctx, newUser)
	}
	// Then
	for index, got := range result {
		if want != got {
			t.Errorf("%d: want: %t, but got: %t", index, want, got)
		}
	}
}
