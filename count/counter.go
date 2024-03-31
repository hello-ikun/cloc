package cloc

import (
	"bufio"
	"os"
	"strings"
)

type LanguageStats struct {
	FileName     string
	CodeLines    int
	BlankLines   int
	CommentLines int
}
type CounterStyle interface {
	Counter(fileName string) (*LanguageStats, error)
}

// SlashCounter 使用斜杠进行注释的 /
type SlashCounter struct {
}

func (s *SlashCounter) Counter(fileName string) (*LanguageStats, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	stats := &LanguageStats{}
	scanner := bufio.NewScanner(file)
	inCommentBlock := false // 是否在多行注释块中

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if inCommentBlock {
			if strings.Contains(line, "*/") { // 如果找到注释块的结束标记
				inCommentBlock = false
			}
			stats.CommentLines++
			continue
		}

		if strings.HasPrefix(line, "/*") { // 如果找到多行注释块的开始标记
			stats.CommentLines++
			inCommentBlock = true
			continue
		}

		if line == "" {
			stats.BlankLines++
		} else if strings.HasPrefix(line, "//") {
			stats.CommentLines++
		} else {
			stats.CodeLines++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return stats, nil
}

// 使用井号进行注释的 #
type WellCounter struct {
}

func (w *WellCounter) Counter(fileName string) (*LanguageStats, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := &LanguageStats{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			stats.BlankLines++
		} else if strings.HasPrefix(line, "#") {
			stats.CommentLines++
		} else {
			stats.CodeLines++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return stats, nil
}

// 不支持的注释类型 不统计注释信息
type NoneCounter struct {
}

func (n *NoneCounter) Counter(fileName string) (*LanguageStats, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := &LanguageStats{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			stats.BlankLines++
		} else {
			stats.CodeLines++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return stats, nil
}
