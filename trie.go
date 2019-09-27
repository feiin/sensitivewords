package sensitivewords

//TrieTree Trie树
type TrieTree struct {
	Root *TrieNode
}

//TrieNode 树节点
type TrieNode struct {
	IsEnd    bool
	Children map[rune]*TrieNode
}

//NewTrieTree
func NewTrieTree() *TrieTree {
	return &TrieTree{
		Root: &TrieNode{
			IsEnd:    false,
			Children: make(map[rune]*TrieNode),
		},
	}
}

//Add 添加敏感词
func (tree *TrieTree) Add(word string) {

	if word == "" {
		return
	}

	treeNode := tree.Root
	var treeWords = []rune(word)
	current := treeNode
	for _, word := range treeWords {
		if node, ok := current.Children[word]; ok {
			current = node

		} else {
			newNode := &TrieNode{
				IsEnd:    false,
				Children: make(map[rune]*TrieNode),
			}
			current.Children[word] = newNode
			current = newNode
		}
	}
	current.IsEnd = true //end
}

//AddWords 添加敏感词数组
func (tree *TrieTree) AddWords(words ...string) {
	for _, word := range words {
		tree.Add(word)
	}
}

//Filter 过滤敏感词为*
func (tree *TrieTree) Filter(input string) string {
	words := []rune(input)

	var current *TrieNode
	var found bool
	treeNode := tree.Root
	offset := 0

	for i := 0; i < len(words); i++ {

		w := words[i]
		current, found = treeNode.Children[w]

		if !found {

			i = i - offset //fallback
			offset = 0
			treeNode = tree.Root
			continue
		}

		if current.IsEnd {
			//found

			for j := i - offset; j < i+1; j++ {
				words[j] = rune('*')
			}

			offset = 0
			treeNode = tree.Root
			continue

		}

		offset++
		treeNode = current

	}

	return string(words)
}

//Find 查找敏感词,找到第一个就退出
func (tree *TrieTree) Find(input string) (sensitive bool, keyword string) {
	words := []rune(input)

	var current *TrieNode
	var found bool
	treeNode := tree.Root
	offset := 0

	for i := 0; i < len(words); i++ {

		w := words[i]
		current, found = treeNode.Children[w]

		if !found {

			i = i - offset //fallback
			offset = 0
			treeNode = tree.Root
			continue
		}

		if current.IsEnd {
			//found
			return true, string(words[i-offset : i+1])

		}

		offset++
		treeNode = current

	}

	return false, ""
}

//FindAll 查找ALL敏感词
func (tree *TrieTree) FindAll(input string) (sensitive bool, results []string) {
	words := []rune(input)

	var current *TrieNode
	var found bool
	treeNode := tree.Root
	offset := 0
	// var results []string

	for i := 0; i < len(words); i++ {

		w := words[i]
		current, found = treeNode.Children[w]

		if !found {

			i = i - offset //fallback
			offset = 0
			treeNode = tree.Root
			continue
		}

		if current.IsEnd {
			sensitive = true
			//found
			results = append(results, string(words[i-offset:i+1]))

			offset = 0
			treeNode = tree.Root
			continue
		}

		offset++
		treeNode = current

	}

	return sensitive, results
}

//FindAny 找到N个敏感词才退出
func (tree *TrieTree) FindAny(input string, count int) (sensitive bool, results []string) {
	words := []rune(input)

	var current *TrieNode
	var found bool
	treeNode := tree.Root
	offset := 0
	// var results []string

	for i := 0; i < len(words); i++ {

		w := words[i]
		current, found = treeNode.Children[w]

		if !found {

			i = i - offset //fallback
			offset = 0
			treeNode = tree.Root
			continue
		}

		if current.IsEnd {
			sensitive = true
			//found
			results = append(results, string(words[i-offset:i+1]))

			if len(results) == count {
				break
			}
			offset = 0
			treeNode = tree.Root
			continue
		}

		offset++
		treeNode = current

	}

	return sensitive, results
}
