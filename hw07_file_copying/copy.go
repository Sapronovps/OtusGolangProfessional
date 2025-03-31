package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, fileSize, err := OpenReadFile(fromPath, offset)
	if err != nil {
		return err
	}

	toFile, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("error create file: %w", err)
	}

	_, err = CopyFromTo(fromFile, toFile, limit, fileSize)
	if err != nil {
		return fmt.Errorf("error copy: %w", err)
	}

	if err := fromFile.Close(); err != nil {
		return fmt.Errorf("error close file: %w", err)
	}
	if err := toFile.Close(); err != nil {
		return fmt.Errorf("error close file: %w", err)
	}

	return nil
}

func OpenReadFile(path string, offset int64) (file *os.File, fileSize int64, err error) {
	file, err = os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return file, 0, fmt.Errorf("error: %w", err)
		}
		return file, 0, err
	}

	stats, err := file.Stat()
	if err != nil {
		return file, 0, err
	}

	if stats.IsDir() {
		return nil, 0, fmt.Errorf("%w: %s", ErrUnsupportedFile, path)
	}

	if offset > stats.Size() {
		return file, 0, ErrOffsetExceedsFileSize
	}

	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return file, 0, err
	}

	return file, stats.Size(), nil
}

func CopyFromTo(fromFile io.Reader, toFile io.Writer, limit, fileSize int64) (n int64, err error) {
	needCopy := fileSize
	var copied int64

	if limit > 0 {
		needCopy = limit
	}

	for {
		n, err := io.CopyN(toFile, fromFile, needCopy)
		copied += n

		showProgress(copied, needCopy)
		if needCopy <= copied {
			break
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return copied, err
		}
	}

	return copied, nil
}

func showProgress(copied, fileSize int64) {
	fmt.Printf("\r%.2f%%", float64(copied)/float64(fileSize)*100)
}
