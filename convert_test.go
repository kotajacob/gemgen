package main

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/spf13/afero"
)

func setupData(count int) (afero.Fs, *Opts, error) {
	data, err := os.ReadFile("test.md")
	if err != nil {
		return nil, nil, err
	}
	fs := afero.NewMemMapFs()
	fs.Mkdir("md", 0755)
	var opts Opts
	opts.Names = make([]string, 0, count)
	// Create count files with test data under md/
	for i := 0; i < count; i++ {
		s := strconv.FormatInt(int64(i), 10)
		path := filepath.Join("md", s+".md")
		opts.Names = append(opts.Names, path)
		f, err := fs.Create(path)
		if err != nil {
			return nil, nil, err
		}
		_, err = f.Write(data)
		if err != nil {
			return nil, nil, err
		}
		err = f.Close()
		if err != nil {
			return nil, nil, err
		}
	}
	return fs, &opts, nil
}

func TestConvertFiles(t *testing.T) {
	// create 100 test markdown files
	fs, opts, err := setupData(100)
	if err != nil {
		t.Fatal(err)
	}
	err = convertFiles(fs, opts)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkConvertFiles100(b *testing.B) {
	// create 100 test markdown files
	fs, opts, err := setupData(100)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		err = convertFiles(fs, opts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkConvertFiles100000(b *testing.B) {
	// create 100 test markdown files
	fs, opts, err := setupData(100000)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		err = convertFiles(fs, opts)
		if err != nil {
			b.Fatal(err)
		}
	}
}
