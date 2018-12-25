package main

import "fmt"

/**
	现在有一种字符串的压缩规则是这样的：k[string]
		表示string 连续出现k 次（0 < k < 1000）。
		比如：s = "ef3[a]2[bc]gh", 解压后的字符串为： "efaaabcbcgh"
		这种压缩也可以相互嵌套：s = "3[a2[c]]", 解压后为： "accaccacc"
		输入一个压缩的字符串s，请输出解压后的字符串。
		输入都是严格合法的，数字只用来表示重复次数，不会出现3a 或者2[4]这样的输入。

	思路:
		这是一个 递归的问题。
		关键在于如何提取出 如"3[asdf3[asdf]]"这样的字符串，并且递归计算结果。
**/
func main() {
	//一个测试的压缩字符串
	var str = "w3[x2[y]d2[z]]oo"
	var index = 0
	var result = ""
	for index < len(str) {
		//表示读取到了数字
		if str[index] >= 48 && str[index] <= 57 {
			//tempStr 表示提取出来数字后面代表的[];如3[x2[y]d2[z]]
			tempStr := getCommaString(str[index:])
			index += len(tempStr)
			//递归计算出 []表达式的结果;包括其中嵌套的[]
			result += getDecodeStr(tempStr)
		} else {
			result += string(str[index])
			index++
		}
	}
	fmt.Println("result：", result)

}

//输入的是 带有[] 的字符串，返回的是解密后的字符串
//比如3[a2[c]asdf]
func getDecodeStr(temp string) string {
	var index = 0
	var count = int(temp[index] - 48) //表示 有多少个重复的字符串
	var tempString = ""               //临时字符串
	var result = ""
	index += 2 //避开括号 [
	for index < len(temp) {
		if temp[index] >= 48 && temp[index] <= 57 {
			//又遇上了数字 2[c] 需要能够提取出来这些数据
			temptempString := getCommaString(temp[index:])
			index += len(temptempString)
			tempString += getDecodeStr(temptempString)
		} else if temp[index] == 93 { //对应]
			index++
		} else {
			//表示是正常的字符串
			tempString += string(temp[index])
			index++
		}

	}
	for i := 0; i < count; i++ {
		result += tempString
	}
	return result
}

// 以数字开头的字符串,截取出来对应的字符串 3[a2[c]]asdf ,需要能够提取出来3[a2[c]]
func getCommaString(temp string) string {
	var commaCount = 1
	var index = 2
	var tempString = temp[0:2] //临时字符串
	for ; index < len(temp) && commaCount != 0; index++ {
		if temp[index] >= 48 && temp[index] <= 57 {
			//表示又遇上数字了
			commaCount += 1
			tempString += temp[index : index+2]
			index++
		} else if temp[index] == 93 { //对应]
			commaCount -= 1
			tempString += string(temp[index])
		} else {
			//表示是正常的字符串
			tempString += string(temp[index])
		}

	}
	return tempString
}
