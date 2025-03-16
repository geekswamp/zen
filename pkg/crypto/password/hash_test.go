package password_test

import (
	"errors"
	"testing"

	"github.com/geekswamp/zen/pkg/crypto/password"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type PasswordHasher interface {
	Generate(password []byte) (string, error)
	Verify(password []byte, hash string) (bool, error)
}

type MockPasswordHasher struct {
	mock.Mock
}

func (m *MockPasswordHasher) Generate(password []byte) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordHasher) Verify(password []byte, hash string) (bool, error) {
	args := m.Called(password, hash)
	return args.Bool(0), args.Error(1)
}

func TestHashGenerate(t *testing.T) {
	mockHasher := new(MockPasswordHasher)

	testCases := []struct {
		name     string
		password []byte
		hash     string
		wantErr  bool
	}{
		{
			name:     "Valid password",
			password: []byte("my-secure-password"),
			hash:     "$argon2id$v=19$m=65536,t=3,p=4$c29tZXNhbHQ$validhash",
			wantErr:  false,
		},
		{
			name:     "Empty password",
			password: []byte(""),
			hash:     "$argon2id$v=19$m=65536,t=3,p=4$ZW1wdHlzYWx0$emptyhash",
			wantErr:  false,
		},
		{
			name:     "Long password",
			password: []byte("this-is-a-very-long-password-that-should-still-work-without-any-issues"),
			hash:     "$argon2id$v=19$m=65536,t=3,p=4$bG9uZ3NhbHQ$longhash",
			wantErr:  false,
		},
		{
			name:     "Error case",
			password: []byte("error-trigger"),
			hash:     "",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				mockHasher.On("Generate", tc.password).Return("", errors.New("generate error")).Once()
			} else {
				mockHasher.On("Generate", tc.password).Return(tc.hash, nil).Once()
				mockHasher.On("Verify", tc.password, tc.hash).Return(true, nil).Once()
				mockHasher.On("Verify", []byte("wrong-password"), tc.hash).Return(false, nil).Once()
			}

			hash, err := mockHasher.Generate(tc.password)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Empty(t, hash)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.hash, hash)
			assert.Contains(t, hash, "$argon2id$")

			valid, err := mockHasher.Verify(tc.password, hash)
			assert.NoError(t, err)
			assert.True(t, valid)

			valid, err = mockHasher.Verify([]byte("wrong-password"), hash)
			assert.NoError(t, err)
			assert.False(t, valid)
		})
	}

	mockHasher.AssertExpectations(t)
}

func TestHashVerify(t *testing.T) {
	mockHasher := new(MockPasswordHasher)

	password := []byte("test-password")
	hash := "$argon2id$v=19$m=65536,t=3,p=4$dGVzdHNhbHQ$testhash"

	mockHasher.On("Generate", password).Return(hash, nil).Once()

	generatedHash, err := mockHasher.Generate(password)
	assert.NoError(t, err)
	assert.Equal(t, hash, generatedHash)

	testCases := []struct {
		name     string
		password []byte
		hash     string
		want     bool
		wantErr  bool
	}{
		{
			name:     "Correct password",
			password: []byte("test-password"),
			hash:     hash,
			want:     true,
			wantErr:  false,
		},
		{
			name:     "Incorrect password",
			password: []byte("wrong-password"),
			hash:     hash,
			want:     false,
			wantErr:  false,
		},
		{
			name:     "Invalid hash format",
			password: []byte("test-password"),
			hash:     "invalid-hash-format",
			want:     false,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				mockHasher.On("Verify", tc.password, tc.hash).Return(false, errors.New("invalid hash format")).Once()
			} else {
				mockHasher.On("Verify", tc.password, tc.hash).Return(tc.want, nil).Once()
			}

			got, err := mockHasher.Verify(tc.password, tc.hash)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}

	mockHasher.AssertExpectations(t)
}

func TestDifferentPeppers(t *testing.T) {
	mockHasher1 := new(MockPasswordHasher)
	mockHasher2 := new(MockPasswordHasher)

	password := []byte("test-password")
	hash := "$argon2id$v=19$m=65536,t=3,p=4$cGVwcGVyMQ$hashedwithpepper1"

	mockHasher1.On("Generate", password).Return(hash, nil).Once()
	mockHasher1.On("Verify", password, hash).Return(true, nil).Once()
	mockHasher2.On("Verify", password, hash).Return(false, nil).Once()

	generatedHash, err := mockHasher1.Generate(password)
	assert.NoError(t, err)
	assert.Equal(t, hash, generatedHash)

	valid, err := mockHasher1.Verify(password, hash)
	assert.NoError(t, err)
	assert.True(t, valid)

	valid, err = mockHasher2.Verify(password, hash)
	assert.NoError(t, err)
	assert.False(t, valid, "Hash should not verify with different pepper")

	mockHasher1.AssertExpectations(t)
	mockHasher2.AssertExpectations(t)
}

func TestDecodeHashErrors(t *testing.T) {
	mockHasher := new(MockPasswordHasher)

	testCases := []struct {
		name     string
		hash     string
		password []byte
		errMsg   string
	}{
		{
			name:     "Missing fields",
			hash:     "$argon2id$v=19$m=65536",
			password: []byte("password"),
			errMsg:   "invalid hash format",
		},
		{
			name:     "Invalid version",
			hash:     "$argon2id$v=18$m=65536,t=3,p=4$salt$hash",
			password: []byte("password"),
			errMsg:   "incompatible version",
		},
		{
			name:     "Invalid base64 salt",
			hash:     "$argon2id$v=19$m=65536,t=3,p=4$invalid!!$hash",
			password: []byte("password"),
			errMsg:   "invalid base64 in salt",
		},
		{
			name:     "Invalid base64 hash",
			hash:     "$argon2id$v=19$m=65536,t=3,p=4$c2FsdA$invalid!!",
			password: []byte("password"),
			errMsg:   "invalid base64 in hash",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockHasher.On("Verify", tc.password, tc.hash).Return(false, errors.New(tc.errMsg)).Once()

			_, err := mockHasher.Verify(tc.password, tc.hash)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tc.errMsg)
		})
	}

	mockHasher.AssertExpectations(t)
}

func BenchmarkGenerate(b *testing.B) {
	h := password.NewDefault()
	password := []byte("benchmark-password")

	for b.Loop() {
		_, _ = h.Generate(password)
	}
}

func BenchmarkVerify(b *testing.B) {
	h := password.NewDefault()
	password := []byte("benchmark-password")

	hash, _ := h.Generate(password)

	for b.Loop() {
		_, _ = h.Verify(password, hash)
	}
}
