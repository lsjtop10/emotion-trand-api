package entity_test

import (
	"os"
	"testing"
	"time"

	"github.com/lsjtop10/emotion-trand-api/domain/entity"
	"github.com/stretchr/testify/assert"
)

var userID string
var userEmail string
var duration time.Duration
var secret string

var token *entity.Token

func setup() {
	userID = "testUser"
	userEmail = "testUser@email.com"
	duration = time.Hour * 10
	secret = "testsecret"

}

func teardown() {

}

func TestMain(m *testing.M) {
	// 테스트 전처리 작업
	setup()
	exitCode := m.Run()
	teardown()
	// 테스트 후 처리 작업

	os.Exit(exitCode)
}

func TestGenerateToken(t *testing.T) {
	t.Run("올바른 Token 이 생성될 것", func(t *testing.T) {
		token = entity.NewToken(
			userID,
			userEmail,
			duration)

		got, err := token.SignToken(secret)
		if err != nil {
			t.Error(err)
		}

		// 테스트를 검증하기 위해서는 ParseToken 메서드가 구현되어 있을 것
		want, err := entity.ParseToken(got, secret)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, token.UserUUId, want.UserUUId)
		assert.Equal(t, token.UserEmail, want.UserEmail)
	})
}

func TestParseToken(t *testing.T) {
	t.Run("올바른 Token 일 경우, 올바른 User ID 가 반환될 것", func(t *testing.T) {
		// 토큰 생성 테스트에서 확인 완료
		TestGenerateToken(t)
	})

	t.Run("토큰 파싱에 실패했을 경우, 에러가 반환될 것", func(t *testing.T) {
		claim, err := entity.ParseToken("I am Not Token", secret)
		assert.NotNil(t, err)
		assert.Nil(t, claim)
	})

	t.Run("유효하지 않는 secret이면 에러를 반환할 것", func(t *testing.T) {
		//given
		token = entity.NewToken(
			userID,
			userEmail,
			duration)
		tokenString, _ := token.SignToken(secret)

		//when
		claim, err := entity.ParseToken(tokenString, "I'm not secret")

		//than
		assert.NotNil(t, err)
		assert.Nil(t, claim)
	})

	t.Run("유효기간이 지난 Token 일 경우, 에러가 반환될 것", func(t *testing.T) {

		//given
		expiredTokenEntity := entity.NewToken(userID, userEmail, time.Hour*-1)
		tokenString, _ := expiredTokenEntity.SignToken(secret)

		//when

		claim, err := entity.ParseToken(tokenString, secret)

		//then
		assert.NotNil(t, err)
		assert.Nil(t, claim)
	})
}
