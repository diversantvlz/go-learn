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

func (pb *ProgressBar) Finish() {
	fmt.Println()
}

func (pb *ProgressBar) print() {
	percent := int64((float32(pb.progress) / float32(pb.total)) * 100)
	fmt.Printf("\r[%-100s]%d%% %d/%d bytes", strings.Repeat("=", int(percent)), percent, pb.progress, pb.total)
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	if fromPath == toPath {
		return ErrUnsupportedFile
	}

	src, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	stat, err := src.Stat()
	if err != nil {
		return err
	}

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

	pb := ProgressBar{}
	pb.Start(limit)
	dst, err := os.Create(toPath)
	if err != nil {
		pb.Finish()
		return err
	}

	_, err = src.Seek(offset, io.SeekStart)
	if err != nil {
		pb.Finish()
		return err
	}

	for i := int64(1); i <= limit; i++ {
		result, err := io.CopyN(dst, src, 1)
		if err != nil {
			pb.Finish()
			return err
		}
		pb.Advance(result)
	}

	pb.Finish()

	_ = src.Close()
	_ = dst.Close()

	return nil
}
