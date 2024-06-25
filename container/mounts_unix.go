//go:build !windows

package container // import "github.com/docker/docker/container"

// Mount contains information for a mount operation.
type Mount struct {
	ID                     string `json:"-"` // TODO: should I omit this or not?
	Source                 string `json:"source"`
	Destination            string `json:"destination"`
	Writable               bool   `json:"writable"`
	Data                   string `json:"data"`
	Propagation            string `json:"mountpropagation"`
	NonRecursive           bool   `json:"nonrecursive"`
	ReadOnlyNonRecursive   bool   `json:"readonlynonrecursive"`
	ReadOnlyForceRecursive bool   `json:"readonlyforcerecursive"`
}
