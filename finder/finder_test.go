package finder_test

import (
	"count/finder"
	"os"
	"testing"
	"testing/fstest"
)

func TestFilesMemory(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"abc.go":          {},
		"xyz.go":          {},
		"sub/sub1.go":     {},
		"sub/sub1.txt":    {},
		"sub/sub/sub2.go": {},
	}

	want := 4
	got := finder.Files(fsys)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestFilesDisk(t *testing.T) {
	t.Parallel()
	fsys := os.DirFS("testdata/")
	want := 4
	got := finder.Files(fsys)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func BenchmarkFilesOnDisk(b *testing.B) {
	fsys := os.DirFS("testdata/")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		finder.Files(fsys)
	}
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"abc.go":          {},
		"xyz.go":          {},
		"sub/sub1.go":     {},
		"sub/sub1.txt":    {},
		"sub/sub/sub2.go": {},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		finder.Files(fsys)
	}

}
