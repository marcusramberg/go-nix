package eval

type Derivation struct {
	System  string // e.g. "x86_64-linux"
	Name    string
	Builder any // a Derivation or an ast.Path
	Attrs   map[Sym]any
}
