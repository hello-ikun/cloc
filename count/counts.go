package cloc

import (
	"fmt"
	"os"
	"path/filepath"
)

type Counters struct {
	stats map[Token][]*LanguageStats
}
type CounterIndex interface {
	Count(directory string) error
	PrintStats()
}

func NewCounters() CounterIndex {
	return &Counters{stats: make(map[Token][]*LanguageStats)}
}
func (c *Counters) Count(directory string) error {
	stats := c.stats

	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		language := detectLanguage(path)

		if fileStats := stats[language]; fileStats == nil {
			fileStats = make([]*LanguageStats, 0)
		}
		counter, err := NewCodeStyle(language).Counter(path)
		counter.FileName = info.Name()
		if err != nil {
			return err
		}
		stats[language] = append(stats[language], counter)
		return nil
	})
}
func (c *Counters) PrintStats() {
	allFiles := c.stats
	totalFiles := 0
	totalBlankLines := 0
	totalCommentLines := 0
	totalCodeLines := 0
	ikun()
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("%-20s%-30s%-15s%-15s%-15s\n", "Language", "files", "blank", "comment", "code")
	fmt.Println("---------------------------------------------------------------------------------------------")
	for lang, stats := range allFiles {
		fmt.Printf("%-20s\n", Tokens[lang])

		tempFilesFiles := 0
		tempFilesBlankLines := 0
		tempFilesCommentLines := 0
		tempFilesCodeLines := 0
		for _, stat := range stats {
			fmt.Printf("%-20s%-30s%-15d%-15d%-15d\n", "", stat.FileName, stat.BlankLines, stat.CommentLines, stat.CodeLines)
			tempFilesFiles++
			tempFilesBlankLines += stat.BlankLines
			tempFilesCommentLines += stat.CommentLines
			tempFilesCodeLines += stat.CodeLines
		}
		totalFiles += tempFilesFiles
		totalBlankLines += tempFilesBlankLines
		totalCommentLines += tempFilesCommentLines
		totalCodeLines += tempFilesCodeLines
		fmt.Printf("%-20s%-30d%-15d%-15d%-15d\n", "      "+"SUM", len(stats), tempFilesBlankLines, tempFilesCommentLines, tempFilesCodeLines)
		fmt.Println("---------------------------------------------------------------------------------------------")

	}

	fmt.Printf("%-20s%-30d%-15d%-15d%-15d\n", "SUM:", totalFiles, totalBlankLines, totalCommentLines, totalCodeLines)
	fmt.Println("---------------------------------------------------------------------------------------------")
}
