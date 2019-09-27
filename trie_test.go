package sensitivewords

import (
	"testing"
)

func TestTreeAdd(t *testing.T) {
	tree := NewTrieTree()
	tree.Add("测试词语")
	tree.Add("测试语")
	tree.Add("测测测")
	// fmt.Printf("%+v", *tree.Root.Children['测'])
	if tree.Root.Children['测'].IsEnd == true {
		t.Error("error")
	}

	if tree.Root.Children['测'].Children['测'].IsEnd == true {
		t.Error("error")
	}
	if tree.Root.Children['测'].Children['试'].IsEnd == true {
		t.Error("error")
	}
	if tree.Root.Children['测'].Children['测'].Children['测'].IsEnd == true {
		t.Log("测测测 ok")
	}

	if tree.Root.Children['测'].Children['试'].Children['语'].IsEnd == true {
		t.Log("测试语 ok")
	}

	if tree.Root.Children['测'].Children['试'].Children['词'].IsEnd == true {
		t.Log("error")
	}
	if tree.Root.Children['测'].Children['试'].Children['词'].Children['语'].IsEnd == true {
		t.Log("测试词语 ok")
	}
}

func TestTreeFilter(t *testing.T) {
	tree := NewTrieTree()
	tree.Add("测试词语")
	tree.Add("测试语")
	tree.Add("测测测")
	sf := tree.Filter("1测试词语测测2测试语3测测测")
	if sf == "1****测测2***3***" {
		t.Logf("%+v", tree.Filter("1测试词语测测2测试语3测测测"))
	} else {
		t.Error("failed")
	}
}

func TestTreeFind(t *testing.T) {
	tree := NewTrieTree()
	tree.Add("测试词语")
	tree.Add("测试语")
	tree.Add("测测测")
	sensitive, keyword := tree.Find("1测试词语测测2测试语3测测测")
	if sensitive == true && keyword == "测试词语" {
		t.Logf("%v, %v", sensitive, keyword)
	} else {
		t.Error("failed")
	}
}
func TestTreeFindAll(t *testing.T) {
	tree := NewTrieTree()
	tree.Add("测试词语")
	tree.Add("测试语")
	tree.Add("测测测")
	sensitive, results := tree.FindAll("1测试词语测测2测试语3测测测")
	if sensitive == true && len(results) == 3 {
		t.Logf("%v, %v", sensitive, results)
	} else {
		t.Error("failed")
	}
}

func TestTreeFindAny(t *testing.T) {
	tree := NewTrieTree()
	tree.Add("测试词语")
	tree.Add("测试语")
	tree.Add("测测测")
	sensitive, results := tree.FindAny("1测试词语测测2测试语3测测测", 2)
	if sensitive == true && len(results) == 2 {
		t.Logf("%v, %v", sensitive, results)
	} else {
		t.Error("failed")
	}
}
