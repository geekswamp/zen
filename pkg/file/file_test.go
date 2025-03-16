package file_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/geekswamp/zen/pkg/file"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockOS struct {
	mock.Mock
}

func (m *MockOS) Stat(path string) (os.FileInfo, error) {
	args := m.Called(path)
	if args.Get(0) != nil {
		return args.Get(0).(os.FileInfo), args.Error(1)
	}
	return nil, args.Error(1)
}

type MockFileInfo struct {
	mock.Mock
}

func (m *MockFileInfo) Name() string       { return "mockFile" }
func (m *MockFileInfo) Size() int64        { return 1234 }
func (m *MockFileInfo) Mode() os.FileMode  { return os.ModePerm }
func (m *MockFileInfo) ModTime() time.Time { return time.Now() }
func (m *MockFileInfo) IsDir() bool        { return false }
func (m *MockFileInfo) Sys() any           { return nil }

func TestIsExist(t *testing.T) {
	mockOS := new(MockOS)
	mockFileInfo := new(MockFileInfo)

	originalStatFunc := file.StatFunc
	defer func() { file.StatFunc = originalStatFunc }()

	file.StatFunc = func(path string) (os.FileInfo, error) {
		return mockOS.Stat(path)
	}

	testCases := []struct {
		name     string
		path     string
		mockInfo os.FileInfo
		mockErr  error
		want     bool
	}{
		{
			name:     "File exists",
			path:     "existing_file.txt",
			mockInfo: mockFileInfo,
			mockErr:  nil,
			want:     true,
		},
		{
			name:     "File does not exist",
			path:     "non_existing_file.txt",
			mockInfo: nil,
			mockErr:  os.ErrNotExist,
			want:     false,
		},
		{
			name:     "Unexpected error",
			path:     "error_file.txt",
			mockInfo: nil,
			mockErr:  errors.New("some error"),
			want:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockOS.On("Stat", tc.path).Return(tc.mockInfo, tc.mockErr).Once()
			info, exists := file.IsExist(tc.path)

			require.Equal(t, tc.want, exists)
			if exists {
				require.NotNil(t, info)
			} else {
				require.Nil(t, info)
			}

			mockOS.AssertCalled(t, "Stat", tc.path)
		})
	}
}

func BenchmarkIsExist(b *testing.B) {
	mockOS := new(MockOS)
	mockFileInfo := new(MockFileInfo)

	originalStatFunc := file.StatFunc
	defer func() { file.StatFunc = originalStatFunc }()

	file.StatFunc = func(path string) (os.FileInfo, error) {
		return mockOS.Stat(path)
	}

	mockOS.On("Stat", "existing_file.txt").Return(mockFileInfo, nil)

	for i := 0; i < b.N; i++ {
		file.IsExist("existing_file.txt")
	}
}
