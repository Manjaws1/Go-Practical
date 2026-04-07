package techpalace
import "strings"
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	greeting := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return greeting
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	star := strings.Repeat("*", numStarsPerLine)
	return star + "\n" + welcomeMsg + "\n" + star
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanUp := strings.ReplaceAll(oldMsg, "*", "")
    trimSpace := strings.TrimSpace(cleanUp)

    return trimSpace
}
