package sensitivewords_test

import (
	"fmt"
	"os"

	"github.com/feiin/sensitivewords"
)

var sensitive *sensitivewords.SensitiveWords = sensitivewords.New()

func Example() {
	sensitive := sensitivewords.New()
	sensitive.LoadFromFile("./keywords.txt")
	sensitive.AddWord("测试")
	sensitive.AddWords("+q", "+v")

	s, keyword := sensitive.Find("测试啊+q/+v,尼玛,哈哈")
	fmt.Printf("%v,%v\n", s, keyword)
	//Output: true,测试

}

func ExampleNew() {
	sensitive := sensitivewords.New()
	sensitive.AddWord("test")
	fmt.Printf("ok")
	//Output: ok

}

func ExampleSensitiveWords_LoadFromFile() {

	sensitive.LoadFromFile("./keywords.txt")
}

func ExampleSensitiveWords_Load() {

	// 读取文件、网络请求等等
	fs, _ := os.Open("path")
	defer fs.Close()
	sensitive.Load(fs)
}

func ExampleSensitiveWords_Find() {

	sensitive.AddWord("测试")
	s, keyword := sensitive.Find("测试啊+q/+v,尼玛,哈哈")
	fmt.Printf("%v,%v\n", s, keyword) //true, 测试
	// Output:true,测试
}

func ExampleSensitiveWords_FindAll() {

	sensitive.AddWords("测试", "尼玛")
	s, results := sensitive.FindAll("测试啊+q/+v,尼玛,哈哈")
	fmt.Printf("%v,%v\n", s, results)
	// Output: true,[测试 尼玛]
}

func ExampleSensitiveWords_FindAny() {

	sensitive.AddWords("测试", "尼玛", "哈哈", "+q", "+v")
	s, results := sensitive.FindAny("测试啊+q/+v,尼玛,哈哈", 4)
	fmt.Printf("%v,%v\n", s, results)
	// Output: true,[测试 +q +v 尼玛]
}

func ExampleSensitiveWords_AddWord() {

	sensitive.AddWord("测试")
}

func ExampleSensitiveWords_AddWords() {

	sensitive.AddWords("测试", "fuc")
}

func ExampleSensitiveWords_Check() {

	find := sensitive.Check("测试")
	fmt.Printf("%v", find)
	// Output: true
}
