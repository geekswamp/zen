package password_test

import (
	"testing"

	"github.com/geekswamp/zen/pkg/crypto/password"
	"github.com/stretchr/testify/assert"
)

func TestHashGenerate(t *testing.T) {
	h := password.NewDefault()

	testCases := []struct {
		name     string
		password []byte
		wantErr  bool
	}{
		{
			name:     "Valid password",
			password: []byte("my-secure-password"),
			wantErr:  false,
		},
		{
			name:     "Empty password",
			password: []byte(""),
			wantErr:  false,
		},
		{
			name:     "Long password",
			password: []byte("this-is-a-very-long-password-that-should-still-work-without-any-issues-even-though-it-is-quite-lengthy"),
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hash, err := h.Generate(tc.password)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, hash)

			assert.Contains(t, hash, "$argon2id$")

			valid, err := h.Verify(tc.password, hash)
			assert.NoError(t, err)
			assert.True(t, valid)

			valid, err = h.Verify([]byte("wrong-password"), hash)
			assert.NoError(t, err)
			assert.False(t, valid)
		})
	}
}

func TestHashVerify(t *testing.T) {
	h := password.NewDefault()

	password := []byte("test-password")
	hash, err := h.Generate(password)
	assert.NoError(t, err)

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
			got, err := h.Verify(tc.password, tc.hash)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDifferentPeppers(t *testing.T) {
	h1 := password.NewDefault()
	h2 := password.New("this-is-pepper", 12288, 3, 32, 32, 1)

	password := []byte("test-password")

	hash, err := h1.Generate(password)
	assert.NoError(t, err)

	valid, err := h1.Verify(password, hash)
	assert.NoError(t, err)
	assert.True(t, valid)

	valid, err = h2.Verify(password, hash)
	assert.NoError(t, err)
	assert.False(t, valid, "Hash should not verify with different pepper")
}

func TestDecodeHashErrors(t *testing.T) {
	h := password.NewDefault()

	testCases := []struct {
		name string
		hash string
	}{
		{
			name: "Missing fields",
			hash: "$argon2id$v=19$m=65536",
		},
		{
			name: "Invalid version",
			hash: "$argon2id$v=18$m=65536,t=3,p=4$salt$hash",
		},
		{
			name: "Invalid base64 salt",
			hash: "$argon2id$v=19$m=65536,t=3,p=4$invalid!!$hash",
		},
		{
			name: "Invalid base64 hash",
			hash: "$argon2id$v=19$m=65536,t=3,p=4$c2FsdA$invalid!!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := h.Verify([]byte("password"), tc.hash)
			assert.Error(t, err)
		})
	}
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
