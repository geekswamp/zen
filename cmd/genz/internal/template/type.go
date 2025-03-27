package template

type fileType string
type FilePath string
type suffix string

const (
	Repository fileType = "repository"
	Router     fileType = "router"
	Service    fileType = "service"
	Handler    fileType = "handler"
	Model      fileType = "model"
)

const (
	_Internal               = "internal"
	RepositoryPath FilePath = _Internal + "/repository"
	ModelPath      FilePath = _Internal + "/model"
	ServicePath    FilePath = _Internal + "/service"
)

const (
	RepoSuffix    suffix = "_repository"
	RouterSuffix  suffix = "_router"
	ServiceSuffix suffix = "_service"
	HandlerSuffix suffix = "_handler"
)

type Make struct {
	FilePath    FilePath
	FileType    fileType
	SuffixFile  suffix
	FeatureName string
}
