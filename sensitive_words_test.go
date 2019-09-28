package sensitivewords

import (
	"testing"
)

func TestSensitiveWordsLoad(t *testing.T) {
	sensitive := New()
	sensitive.LoadFromFile("./tests/keywords.txt")
	if len(sensitive.trie.Root.Children) > 0 {
		t.Logf("load success")
	} else {
		t.Errorf("load failed")
	}
}

func TestSensitiveWordsCheck(t *testing.T) {
	sensitive := New()
	sensitive.LoadFromFile("./tests/keywords.txt")

	isSensitive := sensitive.Check("卧槽啊啊啊啊")

	if isSensitive == true {
		t.Logf("check success")
	} else {
		t.Errorf("check failed")
	}

	isSensitive = sensitive.Check("111啊啊啊啊")

	if isSensitive == false {
		t.Logf("check success")
	} else {
		t.Errorf("check failed")
	}
}

func TestSensitiveWordsFindAny(t *testing.T) {
	sensitive := New()
	sensitive.LoadFromFile("./tests/keywords.txt")

	isSensitive, results := sensitive.FindAny("卧槽啊啊啊尼玛", 2)

	if isSensitive == true {
		t.Logf("FindAny success %v", results)
	} else {
		t.Errorf("FindAny failed")
	}

}

func TestSensitiveWordsFind(t *testing.T) {
	sensitive := New()
	sensitive.LoadFromFile("./tests/keywords.txt")

	isSensitive, keyword := sensitive.Find("卧卧槽啊啊啊尼玛")

	if isSensitive == true && keyword == "卧槽" {
		t.Logf("Find success %v", keyword)
	} else {
		t.Errorf("Find failed")
	}

}

func TestSensitiveWordsFindAll(t *testing.T) {
	sensitive := New()
	sensitive.LoadFromFile("./tests/keywords.txt")

	isSensitive, results := sensitive.FindAll("卧卧槽啊啊啊尼玛测测试测试")

	if isSensitive == true && len(results) == 4 {
		t.Logf("FindAll success %v", results)
	} else {
		t.Errorf("FindAll failed")
	}

}
