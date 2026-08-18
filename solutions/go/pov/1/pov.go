package pov

type Tree struct {
	value    string
	children []*Tree
}

func New(value string, children ...*Tree) *Tree {
	return &Tree{value: value, children: children}
}

func (tr *Tree) Value() string {
	return tr.value
}

func (tr *Tree) Children() []*Tree {
	return tr.children
}

func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// buildAdj 构建无向邻接表 value -> []neighbor value
func buildAdj(root *Tree) map[string][]string {
	adj := make(map[string][]string)
	var dfs func(cur *Tree)
	dfs = func(cur *Tree) {
		for _, ch := range cur.Children() {
			adj[cur.Value()] = append(adj[cur.Value()], ch.Value())
			adj[ch.Value()] = append(adj[ch.Value()], cur.Value())
			dfs(ch)
		}
	}
	dfs(root)
	return adj
}

// buildValueNodeMap value -> *Tree
func buildValueNodeMap(root *Tree) map[string]*Tree {
	m := make(map[string]*Tree)
	var dfs func(cur *Tree)
	dfs = func(cur *Tree) {
		m[cur.Value()] = cur
		for _, ch := range cur.Children() {
			dfs(ch)
		}
	}
	dfs(root)
	return m
}

// dfsBuild 基于邻接表递归构造新树
// current: 当前节点值
// skip: 不要走回来的邻居，空字符串表示无
func dfsBuild(current string, skip string, adj map[string][]string) *Tree {
	var kids []*Tree
	for _, neighbor := range adj[current] {
		if neighbor == skip {
			continue
		}
		childNode := dfsBuild(neighbor, current, adj)
		kids = append(kids, childNode)
	}
	return New(current, kids...)
}

// findPath DFS搜索路径
func findPath(root *Tree, target string, path []string) ([]string, bool) {
	path = append(path, root.Value())
	if root.Value() == target {
		return path, true
	}
	for _, ch := range root.Children() {
		cp := make([]string, len(path))
		copy(cp, path)
		if p, ok := findPath(ch, target, cp); ok {
			return p, true
		}
	}
	return nil, false
}

// FromPov 返回以from为根的新树；找不到返回nil
func (tr *Tree) FromPov(from string) *Tree {
	valMap := buildValueNodeMap(tr)
	if _, ok := valMap[from]; !ok {
		return nil
	}
	adj := buildAdj(tr)
	return dfsBuild(from, "", adj)
}

// PathTo 返回两点路径
func (tr *Tree) PathTo(from, to string) []string {
	newTree := tr.FromPov(from)
	if newTree == nil {
		return nil
	}
	path, ok := findPath(newTree, to, []string{})
	if !ok {
		return nil
	}
	return path
}