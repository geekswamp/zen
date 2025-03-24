package template

type fileType string

const (
	Repository fileType = "repository"
	Service    fileType = "service"
	Handler    fileType = "handler"
	Router     fileType = "router"
)

type Make struct {
	// filePath    string
	fileType    fileType
	featureName string
}
