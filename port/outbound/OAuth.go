// OAuthPort is an interface for handling OAuth operations.
package outbound

// OAuthPort is an interface for handling OAuth operations.
type OAuthPort interface {
	// GetAccessToken retrieves an access token using the provided code.
	GetAccessToken(code string) (string, error)

	// FetchUserEmail retrieves the user's email using the provided access token.
	FetchUserEmail(accessToken string) (string, error)
	// FetchUserNickName retrieves the user's nickname using the provided access token.
	FetchUserNickName(accessToken string) (string, error)
	// FetchUserProfileImg retrieves the user's profile image URL using the provided access token.
	FetchUserProfileImg(accessToken string) (string, error)
}
