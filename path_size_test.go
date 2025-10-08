package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSizeZeroFile(t *testing.T) {
	path := "testdata/file0.txt"
	want := int64(0)
	got, err := GetSize(path)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetSizeNonZeroFile(t *testing.T) {
	path := "testdata/file1.txt"
	want := int64(1)
	got, err := GetSize(path)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetSizeNonExistentFile(t *testing.T) {
	path := "testdata/file2.txt"
	_, err := GetSize(path)
	require.Error(t, err)
}

func TestGetSizeDirWithOneFile(t *testing.T) {
	path := "testdata/dir1"
	want := int64(2)
	got, err := GetSize(path)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetSizeDirWithTwoFiles(t *testing.T) {
	path := "testdata/dir2"
	want := int64(5)
	got, err := GetSize(path)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestHumanReadableSizeB(t *testing.T) {
	size := int64(1)
	want := "1B"
	got := HumanReadableSize(size)
	require.Equal(t, want, got)
}

func TestHumanReadableSizeKB(t *testing.T) {
	size := int64(1 << 10)
	want := "1KB"
	got := HumanReadableSize(size)
	require.Equal(t, want, got)
}

func TestHumanReadableSizeMB(t *testing.T) {
	size := int64(1 << 20)
	want := "1MB"
	got := HumanReadableSize(size)
	require.Equal(t, want, got)
}

func TestHumanReadableSizeGB(t *testing.T) {
	size := int64(1 << 30)
	want := "1GB"
	got := HumanReadableSize(size)
	require.Equal(t, want, got)
}

func TestHumanReadableSizeTB(t *testing.T) {
	size := int64(1 << 40)
	want := "1TB"
	got := HumanReadableSize(size)
	require.Equal(t, want, got)
}

func TestHumanReadableSizePB(t *testing.T) {
	size := int64(1 << 50)
	want := "1PB"
	got := HumanReadableSize(size)
	require.Equal(t, want, got)
}

func TestHumanReadableSizeEB(t *testing.T) {
	size := int64(1 << 60)
	want := "1EB"
	got := HumanReadableSize(size)
	require.Equal(t, want, got)
}
