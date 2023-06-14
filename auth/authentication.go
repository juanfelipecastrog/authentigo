package auth

func VerifyCredentials(clientID, clientSecret string, users map[string]string) bool {
	validPassword, ok := users[clientID]
	return ok && clientSecret == validPassword
}
