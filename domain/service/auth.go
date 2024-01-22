package service

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) GoogleLogin(authCode string) (accessToken string, refreshToken string, err error) {
	// auth code로 user의 access token을 발급받는다.
	// access token으로 유저의 이메일 주소를 가져온다.
	// 이메일 주소와 서비스 공급자를 이용해 우리 사이트에 가입된 회원인지 확인한다.
	//   1. 만약 우리 사이트에 가입된 회원이 아니면
	//   oauth provider에게 추가적인 정보를 더 요청한다.
	//   그 정보를 바탕으로 user entity를 생성한 user 저장소에 저장한다.
	//
	// accessToken을 발급한다.
	// refreshToken을 발급한다.
	// refreshToken을 저장한다.
	// accessToken과 refreshToken을 user에게 되돌려준다.
	return "", "", nil
}
