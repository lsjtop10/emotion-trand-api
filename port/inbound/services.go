package inbound

import (
	"time"

	"github.com/lsjtop10/emotion-trand-api/domain/entity"
)

// AuthServicePort is an interface for authentication service.
type AuthServicePort interface {
	// GoogleLogin performs Google login with the provided authCode and returns the access token, refresh token, and any error encountered.
	GoogleLogin(authCode string) (accessToken string, refreshToken string, err error)
}

// PostServicePort is an interface for post service.
type PostServicePort interface {
	// GetPost retrieves all posts and returns them along with any error encountered.
	GetPost() ([]entity.Post, error)
	// GetPostAfter retrieves posts created after the specified time and returns them along with any error encountered.
	GetPostAfter(time.Time) ([]entity.Post, error)
	// GetPostBefore retrieves posts created before the specified time and returns them along with any error encountered.
	GetPostBefore(time.Time) ([]entity.Post, error)
	// PostPost creates a new post with the provided data and returns any error encountered.
	PostPost(entity.Post) error
}

// UserServicePort is an interface for user service.
type UserServicePort interface {
	// GetUserInfo retrieves the user information associated with the provided access token and returns it along with any error encountered.
	GetUserInfo(accessToken string) (entity.User, error)
	// UpdateUserInfo updates the user information associated with the provided access token and returns any error encountered.
	UpdateUserInfo(accessToken string) error
}

// ChartProviderServicePort is an interface for chart provider service.
type ChartProviderServicePort interface {
	// GetCandleSticks retrieves the candlestick elements for charting and returns them along with any error encountered.
	GetCandleSticks(timeslice time.Duration) ([]entity.CandlestickElement, error)
	GetCandleSticksAfter(after time.Time, timeslice time.Duration) ([]entity.CandlestickElement, error)
	// GetAverages retrieves the average elements for charting and returns them along with any error encountered.
	GetAverages() ([]entity.AverageElement, error)
	GetAveragesAfter() ([]entity.AverageElement, error)
}
