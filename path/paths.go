package path

import (
	"fmt"
	"os"
	"path/filepath"
)

type Path struct {
	Home   string
	Config string
	Data   string
	Logs   string
}

type FileType string

const (
	Home FileType = "home"
	Config FileType = "config"
	Data FileType = "data"
	Logs FileType = "logs"
)

var Paths = New()

func New() *Path {
	return &Path{}
}

func (paths *Path) InitPaths(cfg *Path) error {
	err := paths.initPaths(cfg)
	if err != nil {
		return err
	}

	// make sure the data path exists
	err = os.MkdirAll(paths.Data, 0750)
	if err != nil {
		return fmt.Errorf("Failed to create data path %s: %v", paths.Data, err)
	}

	return nil
}

// InitPaths sets the default paths in the configuration based on CLI flags,
// configuration file and default values. It also tries to create the data
// path with mode 0750 and returns an error on failure.
func InitPaths(cfg *Path) error {
	return Paths.InitPaths(cfg)
}

// initPaths sets the default paths in the configuration based on CLI flags,
// configuration file and default values.
func (paths *Path) initPaths(cfg *Path) error {
	*paths = *cfg

	// default for config path
	if paths.Config == "" {
		paths.Config = paths.Home
	}

	// default for data path
	if paths.Data == "" {
		paths.Data = filepath.Join(paths.Home, "data")
	}

	// default for logs path
	if paths.Logs == "" {
		paths.Logs = filepath.Join(paths.Home, "logs")
	}

	return nil
}

func (paths *Path) Resolve(fileType FileType, path string) string {
	// absolute paths are not changed
	if filepath.IsAbs(path) {
		return path
	}

	switch fileType {
	case Home:
		return filepath.Join(paths.Home, path)
	case Config:
		return filepath.Join(paths.Config, path)
	case Data:
		return filepath.Join(paths.Data, path)
	case Logs:
		return filepath.Join(paths.Logs, path)
	default:
		panic(fmt.Sprintf("Unknown file type: %s", fileType))
	}
}

// Resolve resolves a path to a location in one of the default
// folders. For example, Resolve(Home, "test") returns an absolute
// path for "test" in the home path.
// In case path is already an absolute path, the path itself is returned.
func Resolve(fileType FileType, path string) string {
	return Paths.Resolve(fileType, path)
}

func (paths *Path) String() string {
	return fmt.Sprintf("Home path: [%s] Config path: [%s] Data path: [%s] Logs path: [%s]",
		paths.Home, paths.Config, paths.Data, paths.Logs)
}

