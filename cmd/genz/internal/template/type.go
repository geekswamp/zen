package template

type fileType string

type FilePath string

const (
	Repository fileType = "repository"
	Router     fileType = "router"
	Service    fileType = "service"
	Handler    fileType = "handler"
	Model      fileType = "model"
)

const (
	_Internal               = "internal"
	RepositoryPath FilePath = "repository"
	ModelPath      FilePath = _Internal + "/model"
)

type Make struct {
	FilePath    FilePath
	FileType    fileType
	FeatureName string
}
