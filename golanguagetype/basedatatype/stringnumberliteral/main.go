package main

func main() {
	// 字符字面值 之 单引号
	'a'  // ASCII字符
	'中'  // Unicode字符集中的中文字符
	'\n' // 换行字符
	'\'' // 单引号字符

	// 字符字面值 之 unicode 专用的转义字符作为前缀
	'\u4e2d'     // 字符：中
	'\U00004e2d' // 字符：中
	'\u0027'     // 单引号字符

	// 字符字面值 之 整型值作为字符字面值
	'\x27' // 使用十六进制表示的单引号字符
	'\047' // 使用八进制表示的单引号字符

	// 字符串字面值
	"abc\n"
	"中国人"
	"\u4e2d\u56fd\u4eba"                   // 中国人
	"\U00004e2d\U000056fd\U00004eba"       // 中国人
	"中\u56fd\u4eba"                        // 中国人，不同字符字面值形式混合在一起
	"\xe4\xb8\xad\xe5\x9b\xbd\xe4\xba\xba" // 十六进制表示的字符串字面值：中国人

}
