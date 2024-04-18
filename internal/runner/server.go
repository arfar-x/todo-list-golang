package runner

import "gorm.io/gorm"

type Runner interface {
	GetName() string
	ValidateCliArgs([]string) error
	Run([]string, gorm.DB)
}

type Factory struct {
	Name string
}

func (r *Factory) Serve(serverType string) Runner {
	switch serverType {
	case "http":
		return &Http{}
	case "cli":
		return &Cli{}
	default:
		return nil
	}
}

func GetServerTypes() []string {
	return []string{
		"http",
		"cli",
	}
}
