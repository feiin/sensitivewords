package sensitivewords

import (
	"bufio"
	"io"
	"os"
)

type SensitiveWords struct {
	trie *TrieTree
}

//New 返回敏感词库实例
func New() *SensitiveWords {
	return &SensitiveWords{
		trie: NewTrieTree(),
	}
}

//Load 加载敏感词
func (sensitiveWords *SensitiveWords) Load(rd io.Reader) error {
	buf := bufio.NewReader(rd)

	for {
		keywords, _, err := buf.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		sensitiveWords.trie.Add(string(keywords))
	}
	return nil
}

//LoadFromFile 加载敏感词库文件
func (sensitiveWords *SensitiveWords) LoadFromFile(path string) error {

	fs, err := os.Open(path)
	defer fs.Close()
	if err != nil {
		return nil
	}

	return sensitiveWords.Load(fs)

}

//AddWord  添加敏感词
func (sensitiveWords *SensitiveWords) AddWord(word string) {
	sensitiveWords.trie.Add(word)
}

//AddWords 添加敏感词数组
func (sensitiveWords *SensitiveWords) AddWords(words ...string) {
	sensitiveWords.trie.AddWords(words...)
}

//Filter 过滤敏感词为*
func (sensitiveWords *SensitiveWords) Filter(input string) string {
	return sensitiveWords.trie.Filter(input)
}

//Find 查找敏感词,找到第一个就退出
func (sensitiveWords *SensitiveWords) Find(input string) (sensitive bool, keyword string) {
	return sensitiveWords.trie.Find(input)
}

//Check 是否包含敏感词
func (sensitiveWords *SensitiveWords) Check(input string) (sensitive bool) {
	sensitive, _ = sensitiveWords.trie.Find(input)
	return sensitive
}

//FindAll 查找ALL敏感词
func (sensitiveWords *SensitiveWords) FindAll(input string) (sensitive bool, results []string) {
	return sensitiveWords.trie.FindAll(input)
}

//FindAny 找到N个敏感词才退出
func (sensitiveWords *SensitiveWords) FindAny(input string, count int) (sensitive bool, results []string) {
	return sensitiveWords.trie.FindAny(input, count)
}
