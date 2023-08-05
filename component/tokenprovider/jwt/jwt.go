package jwt

import (
	"Delivery_Food/component/tokenprovider"
	"github.com/form3tech-oss/jwt-go"
	"time"
)

type jwtProvider struct {
	secret string
}

func NewJwtProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (p *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	// generate token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		Payload: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(p.secret))
	if err != nil {
		return nil, tokenprovider.ErrEncoding
	}

	return &tokenprovider.Token{
		Token:   myToken,
		Created: time.Now().Local(),
		Expiry:  expiry,
	}, nil
}

func (p *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	t, err := jwt.ParseWithClaims(token, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(p.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	if !t.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := t.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	return &claims.Payload, nil
}

func (p *jwtProvider) String() string {
	return "JWT implement TokenProvider"
}
