// The primary type that represents a sorted set
// backed by a skiplist
type {{.Name}}SortedSet struct {
	less      func(a, b {{.Pointer}}{{.Name}}) bool
	head      []*sortedSet{{.Name}}Element
	length    int
	maxLevels int
	r         *rand.Rand
}

// the struct to hold elements of the skiplist
type sortedSet{{.Name}}Element struct {
	val  {{.Pointer}}{{.Name}}
	next []*sortedSet{{.Name}}Element
}

// Creates and returns a reference to an empty set.
func New{{.Name}}SortedSet(less func({{.Pointer}}{{.Name}}, {{.Pointer}}{{.Name}}) bool) {{.Name}}SortedSet {
	return {{.Name}}SortedSet{
		less:      less,
		maxLevels: 64,
		head:      make([]*sortedSet{{.Name}}Element, 64),
		r:         rand.New(rand.NewSource(123123)),
	}
}

// Cardinality returns how many items are currently in the set.
func (ss {{.Name}}SortedSet) Cardinality() int {
	e := ss.head[0]
	ret := 0
	for e != nil {
		ret++
		e = e.next[0]
	}
	return ret
}
