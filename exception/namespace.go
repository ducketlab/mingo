package exception

var (
	usedNamespace = GlobalNamespace
)

type Namespace string

func (ns Namespace) String() string {
	return string(ns)
}

const (
	// GlobalNamespace Some exceptions common to all services
	GlobalNamespace = Namespace("global")
)