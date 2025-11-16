package docs

import (
	"strings"
)

type FileDeclarations struct {
	Controllers []*Controller
	Dtos        []*Dto
	Domains     []*Domain
	fileSet     *FileSet
}

const (
	PathDeclaration = iota
)

type Controller struct {
	BasePath    string
	Summary     string
	Description string
	Tags        []string
	Path        string
	Method      string
	Accept      string
	Returns     map[string]string
	Body        string
	QueryParams string
	RouteParams string
	BearerAuth  bool
	Examples    map[string]any
}

type Dto struct {
	Name   string
	Fields []*DtoField
}

type DtoField struct {
	Name       string
	Type       string
	Validation map[string]string
	Example    string
}

type Domain struct {
	Name   string
	Fields []*DomainField
}

type DomainField struct {
	Name       string
	Type       string
	Validation map[string]string
	Example    string
}

func NewFileDeclarations(fileSet *FileSet) *FileDeclarations {
	fd := &FileDeclarations{
		fileSet:     fileSet,
		Controllers: []*Controller{},
	}
	fd.FetchDeclarations()
	return fd
}

func tagFromString(controllerName string) string {
	newControllerName := strings.Join(strings.Split(controllerName, "_"), "")
	newControllerName = strings.ToUpper(string(newControllerName[0])) + newControllerName[1:]
	newControllerName = strings.TrimSuffix(newControllerName, "controller.go")
	return newControllerName
}

func (fileDeclarations *FileDeclarations) FetchDeclarations() {
	FindControllerDeclarations(fileDeclarations)
	FindDtoDeclarations(fileDeclarations)
	FindDomainDeclarations(fileDeclarations)
}
