package calc

import "math/rand/v2"

// Helper function to generate a random number within a range
func RandRange(min, max int) int {
	return rand.IntN(max-min) + min
}

// Helper function to check if a percentage chance occurs
// min and max are percentages (e.g., 10 for 10%, 50 for 50%, etc.)
// Returns true if the chance occurs, false otherwise
func RandPercentageChance(chance int) bool {

	// Generate a random number between 1 and 100
	randomNumber := RandRange(1, 101)

	// Check if the random number is less than or equal to the chance
	return randomNumber <= chance
}
