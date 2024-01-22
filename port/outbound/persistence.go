package outbound

import (
	"time"

	"github.com/lsjtop10/emotion-trand-api/domain/entity"
)

// TransactionBeginner is an interface that defines the methods for beginning a transaction.
type TransactionBeginner interface {
	BeginTransaction() (TransactorPort, error)
}

// TransactorPort is an interface that defines the methods for committing or rolling back a transaction.
type TransactorPort interface {
	Commit() error
	Rollback() error
}

// UserRepositoryPort is an interface that defines the methods for user repository operations.
type UserRepositoryPort interface {
	// StoreUser stores the user information in the user repository.
	StoreUser(user entity.User) error
	// RemoveUser deletes the information of the user with the given userId from the user repository.
	RemoveUser(userId string) error
	// GetUserInfo retrieves the information of the user with the given userId from the user repository.
	GetUserInfo(userId string) (entity.User, error)
	// IdentifyUserBySocialLogin identifies the userId of the user authenticated by social login based on the provider and email.
	IdentifyUserBySocialLogin(provider entity.OAuthProvider, email string) (string, error)
}

// TokenStoragePort is an interface that defines the methods for token storage operations.
type TokenStoragePort interface {
	// IsExistToken checks if the token exists in the token storage.
	IsExistToken(token string) (bool, error)
	// ExpireToken expires the token in the token storage.
	ExpireToken(token string) error
	// StoreToken stores the token in the token storage with the specified duration.
	StoreToken(token string, duration time.Duration) error
}

// PostRepositoryPort is an interface that defines the methods for post repository operations.
type PostRepositoryPort interface {
	// StorePost stores the post in the post repository.
	StorePost(post entity.Post) error
	// DeletePost deletes the post with the given postId from the post repository.
	DeletePost(postId string) error
	// UpdatePost modifies the post with the given postId in the post repository.
	UpdatePost(postId string, post entity.Post) error
	// GetPostAll retrieves all the posts from the post repository.
	GetPostAll() ([]entity.Post, error)
	// GetPostFrom retrieves the posts from the post repository starting from the specified time.
	GetPostFrom(from time.Time) ([]entity.Post, error)
}

// ChartProviderPort is an interface that defines the methods for chart provider operations.
type ChartProviderPort interface {
	// GetCandlestick retrieves the candlestick data for the specified timeslice.
	GetCandlestick(fromTime time.Time, toTime time.Time) (entity.CandlestickElement, error)
	// GetAverage retrieves the average data starting from the specified time for the specified timeslice.
	GetAverage(fromTime time.Time, toTime time.Time) (int32, error)
}
