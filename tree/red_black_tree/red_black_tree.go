package red_black_tree

type node struct {
	val    int
	parent *node
	left   *node
	right  *node
	color  bool
}
