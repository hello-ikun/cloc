package cloc

import (
	"path/filepath"
)

// Token 表示令牌的类型
type Token uint8

const (
	Go Token = iota
	Java
	Python
	C
	Cpp
	CSharp
	JavaScript
	Ruby
	Swift
	// HTML
	// CSS
	// SQL
	Unknow
)

var Tokens = map[Token]string{
	Go:         "Go",
	Java:       "Java",
	Python:     "Python",
	C:          "C",
	Cpp:        "C++",
	CSharp:     "C#",
	JavaScript: "JavaScript",
	Ruby:       "Ruby",
	Swift:      "Swift",
	// HTML:       "HTML",
	// CSS:        "CSS",
	// SQL:        "SQL",
	Unknow: "Unknow",
}

func detectLanguage(fileName string) Token {
	// 根据文件扩展名识别编程语言类型
	extension := filepath.Ext(fileName)
	switch extension {
	case ".go", ".mod", ".sum":
		return Go
	case ".java":
		return Java
	case ".py":
		return Python
	case ".c":
		return C
	case ".cpp", ".cc", ".cxx":
		return Cpp
	case ".cs":
		return CSharp
	case ".js":
		return JavaScript
	case ".rb":
		return Ruby
	case ".swift":
		return Swift
	// case ".html", ".htm":
	// 	return HTML
	// case ".css":
	// 	return CSS
	// case ".sql":
	// 	return SQL
	default:
		// 如果无法识别扩展名，则返回默认值
		return Unknow // 默认为 Go 语言
	}
}

// NewCodeStyle 根据文件名返回不同的代码统计风格
func NewCodeStyle(codeStyle Token) CounterStyle {
	switch codeStyle {
	case Go, Java, C, Cpp, CSharp, JavaScript, Swift:
		return &SlashCounter{}
	case Python, Ruby:
		return &WellCounter{}
	default:
		return &NoneCounter{}
	}
}
