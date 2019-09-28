# SensitiveWords 

golang实现[DFA算法](https://zh.wikipedia.org/wiki/确定有限状态自动机)的敏感词查找、过滤替换、敏感词匹配检测

## 用法

```golang

package main

import (
	"fmt"

	"github.com/feiin/sensitivewords"
)

func main() {
    sensitive := sensitivewords.New()
    /*
     * keywords.txt:
     * 尼玛
     * 哈哈
     */
    sensitive.LoadFromFile("./keywords.txt") 
    
	sensitive.AddWord("测试")
	sensitive.AddWords("+q", "+v")

	s, keyword := sensitive.Find("测试啊+q/+v,尼玛,哈哈")
	fmt.Printf("Find:%v, %v\n", s, keyword) //true,测试
	s, results := sensitive.FindAll("测试啊+q/+v,尼玛,哈哈")
	fmt.Printf("FindAll:%v, %v\n", s, results) //true, [测试 +q +v 尼玛 哈哈哈]
	s, results = sensitive.FindAny("测试啊+q/+v,尼玛,哈哈", 3)
	fmt.Printf("FindAny:%v, %v\n", s, results) //true, [测试 +q +v]

	s = sensitive.Check("测试啊+q/+v,尼玛,哈哈")
	fmt.Printf("Check: %v\n", s) //true

	str := sensitive.Filter("测试啊+q/+v,尼玛,哈哈")
	fmt.Printf("Filter:%v\n", str) //**啊**/**,**,**
}


```