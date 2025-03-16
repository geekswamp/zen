package token_test

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"testing"
	"time"

	"github.com/geekswamp/zen/pkg/crypto/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRSAKeyProvider struct {
	mock.Mock
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type MockJWTProvider struct {
	mock.Mock
}

func (m *MockRSAKeyProvider) GetPrivateKey() *rsa.PrivateKey {
	return m.privateKey
}

func (m *MockRSAKeyProvider) GetPublicKey() *rsa.PublicKey {
	return m.publicKey
}

func (m *MockJWTProvider) Generate() (hash string, err error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockJWTProvider) Verify(tokenStr string) (claims *jwt.RegisteredClaims, err error) {
	args := m.Called(tokenStr)
	if claims, ok := args.Get(0).(*jwt.RegisteredClaims); ok {
		return claims, args.Error(1)
	}
	return nil, args.Error(1)
}

func generateTestKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func TestGenerateToken(t *testing.T) {
	mockJWT := new(MockJWTProvider)

	testCases := []struct {
		name    string
		mockRet string
		mockErr error
		wantErr bool
	}{
		{
			name:    "Success Generate Token",
			mockRet: "mocked.jwt.token",
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "Failed Generate Token",
			mockRet: "",
			mockErr: errors.New("failed to generate token"),
			wantErr: true,
		},
		{
			name:    "Generated Token is Empty",
			mockRet: "",
			mockErr: nil,
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockJWT.On("Generate").Return(tc.mockRet, tc.mockErr).Once()

			tokenStr, err := mockJWT.Generate()

			if tc.wantErr {
				assert.Error(t, err)
				assert.Empty(t, tokenStr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.mockRet, tokenStr)

			mockJWT.AssertExpectations(t)
		})
	}
}

func TestVerify(t *testing.T) {
	mockJWT := new(MockJWTProvider)

	validToken := "valid.jwt.token"
	expiredToken := "expired.jwt.token"
	invalidToken := "invalid.jwt.token"
	emptyToken := ""

	expectedClaims := &jwt.RegisteredClaims{
		Issuer:    "test_issuer",
		Subject:   "test_subject",
		Audience:  jwt.ClaimStrings{"test_audience"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}

	mockJWT.On("Verify", validToken).Return(expectedClaims, nil).Once()
	mockJWT.On("Verify", expiredToken).Return(nil, errors.New("token expired")).Once()
	mockJWT.On("Verify", invalidToken).Return(nil, errors.New("invalid token")).Once()
	mockJWT.On("Verify", emptyToken).Return(nil, errors.New("empty token")).Once()

	testCases := []struct {
		name      string
		token     string
		wantErr   bool
		wantValid bool
	}{
		{
			name:      "Valid Token",
			token:     validToken,
			wantErr:   false,
			wantValid: true,
		},
		{
			name:      "Expired Token",
			token:     expiredToken,
			wantErr:   true,
			wantValid: false,
		},
		{
			name:      "Invalid Token",
			token:     invalidToken,
			wantErr:   true,
			wantValid: false,
		},
		{
			name:      "Empty Token",
			token:     emptyToken,
			wantErr:   true,
			wantValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			claims, err := mockJWT.Verify(tc.token)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, claims)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, claims)
			assert.Equal(t, "test_issuer", claims.Issuer)
			assert.Equal(t, "test_subject", claims.Subject)
		})
	}

	mockJWT.AssertExpectations(t)
}

func BenchmarkGenerate(b *testing.B) {
	privateKey, publicKey, err := generateTestKeys()
	assert.NoError(b, err)

	mockProvider := &MockRSAKeyProvider{privateKey: privateKey, publicKey: publicKey}
	jwtProvider := token.New("test_issuer", "test_subject", jwt.ClaimStrings{"test_audience"}, time.Hour, mockProvider)

	for b.Loop() {
		_, _ = jwtProvider.Generate()
	}
}

func BenchmarkVerify(b *testing.B) {
	privateKey, publicKey, err := generateTestKeys()
	assert.NoError(b, err)

	mockProvider := &MockRSAKeyProvider{privateKey: privateKey, publicKey: publicKey}
	jwtProvider := token.New("test_issuer", "test_subject", jwt.ClaimStrings{"test_audience"}, time.Hour, mockProvider)

	token, err := jwtProvider.Generate()
	assert.NoError(b, err)

	for b.Loop() {
		_, err := jwtProvider.Verify(token)
		if err != nil {
			b.Fatalf("Verify failed: %v", err)
		}
	}
}
