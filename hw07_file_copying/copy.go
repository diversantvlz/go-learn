package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

type ProgressBar struct {
	total    int64
	progress int64
}

func (pb *ProgressBar) Start(total int64) {
	pb.total = total
	pb.progress = 0

	pb.print()
}

func (pb *ProgressBar) Advance(value int64) {
	pb.progress += value
	pb.print()
}

func (pb *ProgressBar) print() {
	percent := int64((float32(pb.progress) / float32(pb.total)) * 100)
	fmt.Printf("\r[%-100s]%3d%% %8d/%d", strings.Repeat("=", int(percent)), percent, pb.progress, pb.total)
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	pb := ProgressBar{}
	src, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	stat, _ := src.Stat()

	if stat.Size() == 0 || stat.IsDir() {
		return ErrUnsupportedFile
	}

	if offset > stat.Size() {
		return ErrOffsetExceedsFileSize
	}

	maxLimit := stat.Size() - offset
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}

	pb.Start(limit)
	dst, err := os.Create(toPath)
	if err != nil {
		return err
	}

	_, err = src.Seek(offset, 0)
	if err != nil {
		return err
	}

	result, err := io.CopyN(dst, src, limit)
	pb.Advance(result)

	return err
}
