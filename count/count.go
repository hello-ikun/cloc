package cloc

import (
	"fmt"
	"os"
	"path/filepath"
)

type Counter struct {
	stats map[Token]*LanguageStats
}

func NewCounter() CounterIndex {
	return &Counter{stats: make(map[Token]*LanguageStats)}
}
func (c *Counter) Count(directory string) error {
	stats := c.stats

	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		language := detectLanguage(path)

		fileStats := stats[language]
		if fileStats == nil {
			fileStats = &LanguageStats{}
			stats[language] = fileStats
		}
		counter, err := NewCodeStyle(language).Counter(path)
		if err != nil {
			return err
		}
		fileStats.CommentLines += counter.CommentLines
		fileStats.BlankLines += counter.BlankLines
		fileStats.CodeLines += counter.CodeLines
		return nil
	})
}
func (c *Counter) PrintStats() {
	stats := c.stats
	totalFiles := 0
	totalBlankLines := 0
	totalCommentLines := 0
	totalCodeLines := 0
	ikun()
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("%-30s%-15s%-15s%-15s%-15s\n", "Language", "files", "blank", "comment", "code")
	fmt.Println("---------------------------------------------------------------------------------------------")

	for lang, stat := range stats {
		fmt.Printf("%-30s%-15d%-15d%-15d%-15d\n", Tokens[lang], len(stats), stat.BlankLines, stat.CommentLines, stat.CodeLines)

		totalFiles += len(stats)
		totalBlankLines += stat.BlankLines
		totalCommentLines += stat.CommentLines
		totalCodeLines += stat.CodeLines
	}

	fmt.Println("---------------------------------------------------------------------------------------------")
	if len(stats) <= 1 {
		return
	}
	fmt.Printf("%-30s%-15d%-15d%-15d%-15d\n", "SUM:", totalFiles, totalBlankLines, totalCommentLines, totalCodeLines)
	fmt.Println("---------------------------------------------------------------------------------------------")
}
func ikun() {
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Println("               iiii    kkk  kkk   uuu   uuu    nnn     nnn")
	fmt.Println("                ii     kkk kkkk   uuu   uuu    nn nn   nnn")
	fmt.Println("                ii     kkkkkkk    uuu   uuu    nn  nn  nnn")
	fmt.Println("                ii     kkk  kkk   uuu   uuu    nn   nn nn")
	fmt.Println("               iiii    kkk   kkk   uuuuuuuu     nn     nn")
}
