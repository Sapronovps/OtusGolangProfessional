package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

type ProgressReader struct {
	Reader    io.Reader
	TotalSize int64
	Copied    int64
	LastPrint time.Time
}

func (p *ProgressReader) Read(b []byte) (n int, err error) {
	n, err = p.Reader.Read(b)
	p.Copied += int64(n)

	// Ограничиваем частоту обновления вывода (раз в 100мс)
	if time.Since(p.LastPrint) > 100*time.Millisecond {
		percent := float64(p.Copied) / float64(p.TotalSize) * 100
		fmt.Printf("\rКопирование: %.2f%%", percent)
		p.LastPrint = time.Now()
	}

	return n, err
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer fromFile.Close()

	// Получаем размер файла
	fromStat, err := fromFile.Stat()
	if err != nil {
		return fmt.Errorf("error stat source file: %w", err)
	}
	totalSize := fromStat.Size()

	if offset > 0 {
		// Если offset больше размера файла
		if offset > totalSize {
			return ErrOffsetExceedsFileSize
		}
		_, err = fromFile.Seek(offset, io.SeekStart)
		if err != nil {
			return fmt.Errorf("error seek source file: %w", err)
		}
	}
	if limit > 0 {
		totalSize = limit
	}

	// Создаем целевой файл
	toFile, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("error create destination file: %w", err)
	}
	defer toFile.Close()

	// Обёртка для отслеживания прогресса
	progressReader := &ProgressReader{
		Reader:    io.LimitReader(fromFile, totalSize),
		TotalSize: totalSize,
	}

	// Копируем файл
	_, err = io.Copy(toFile, progressReader)
	if err != nil {
		return fmt.Errorf("error copy file to destination file: %w", err)
	}

	fmt.Println("\nКопирование завершено: 100%")
	return nil
}
