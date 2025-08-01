package include

type Element int

const (
	SCRIPT Element = iota
	STYLE
	LINK
)

type Props struct {
	Html   Element
	Path   string
	Module bool
	Defer  bool
}
