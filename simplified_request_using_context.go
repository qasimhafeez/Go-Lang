package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Simulate a user request with a specific user ID
	userID := 44531
	userRequest(userID)
}

func userRequest(userID int) {
	// Create a background context
	backgroundCtx := context.Background()

	// Create a context with cancellation capability
	ctxWithCancel, cancel := context.WithCancel(backgroundCtx)
	defer cancel()

	// Create a context with a timeout of 3 seconds
	ctxWithTimeout, cancelTimeout := context.WithTimeout(ctxWithCancel, 3*time.Second)
	defer cancelTimeout()

	// Create a context with a specific deadline
	deadline := time.Now().Add(5 * time.Second)
	ctxWithDeadline, cancelDeadline := context.WithDeadline(ctxWithCancel, deadline)
	defer cancelDeadline()

	// Associate user ID with the context
	ctxWithValue := context.WithValue(ctxWithDeadline, "userID", userID)

	// Use ctxWithTimeout for the RPC call, allowing cancellation after 3 seconds
	result, err := callingMicroservice(ctxWithValue, userID)
	if err != nil {
		fmt.Println("User request failed:", err)
		return
	}

	// Process the result
	fmt.Println("User request result:", result)
}

func callingMicroservice(ctx context.Context, userID int) (string, error) {
	// Extract user ID from context if available
	userIDCtx, userIDExists := ctx.Value("userID").(int)
	if userIDExists {
		fmt.Println("Processing request for User ID:", userIDCtx)
	}

	// Simulate a time-consuming operation
	select {
	case <-time.After(4 * time.Second):
		return "Microservice result", nil
	case <-ctx.Done():
		// Context canceled or timed out
		return "", fmt.Errorf("Microservice call failed: %v", ctx.Err())
	}
}
