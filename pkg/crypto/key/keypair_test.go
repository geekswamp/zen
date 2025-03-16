package key_test

import (
	"crypto/rsa"
	"testing"

	"github.com/geekswamp/zen/pkg/crypto/key"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockConfig struct {
	mock.Mock
}

func (m *MockConfig) GetPrivKeyPath() string {
	args := m.Called()
	return args.String(0)
}

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

func BenchmarkGetPrivateKey(b *testing.B) {
	mockProvider := new(MockRSAKeyProvider)
	mockKey := &rsa.PrivateKey{}
	mockProvider.On("GetPrivateKey").Return(mockKey)

	for b.Loop() {
		_ = mockProvider.GetPrivateKey()
	}
}

func BenchmarkGetPublicKey(b *testing.B) {
	mockProvider := new(MockRSAKeyProvider)
	mockKey := &rsa.PublicKey{}
	mockProvider.On("GetPublicKey").Return(mockKey)

	for b.Loop() {
		_ = mockProvider.GetPublicKey()
	}
}
