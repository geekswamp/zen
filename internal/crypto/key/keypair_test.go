package key_test

import (
	"crypto/rsa"
	"testing"

	"github.com/geekswamp/zen/internal/crypto/key"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockRSAKeyProvider struct {
	mock.Mock
}

func (m *MockRSAKeyProvider) GetPrivateKey() *rsa.PrivateKey {
	args := m.Called()
	return args.Get(0).(*rsa.PrivateKey)
}

func (m *MockRSAKeyProvider) GetPublicKey() *rsa.PublicKey {
	args := m.Called()
	return args.Get(0).(*rsa.PublicKey)
}

func TestGetPrivateKey(t *testing.T) {
	mockProvider := new(MockRSAKeyProvider)
	mockKey := &rsa.PrivateKey{}

	mockProvider.On("GetPrivateKey").Return(mockKey)
	key := mockProvider.GetPrivateKey()

	require.NotNil(t, key)
	mockProvider.AssertCalled(t, "GetPrivateKey")
}

func TestGetPublicKey(t *testing.T) {
	mockProvider := new(MockRSAKeyProvider)
	mockKey := &rsa.PublicKey{}

	mockProvider.On("GetPublicKey").Return(mockKey)
	key := mockProvider.GetPublicKey()

	require.NotNil(t, key)
	mockProvider.AssertCalled(t, "GetPublicKey")
}

func BenchmarkNewRSAKeyProvider(b *testing.B) {
	for b.Loop() {
		_, _ = key.New()
	}
}
