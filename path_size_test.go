package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSizeZeroFile(t *testing.T) {
	path := "testdata/file0.txt"
	want := int64(0)
	got, err := getSize(path, false, false)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetSizeNonZeroFile(t *testing.T) {
	path := "testdata/file1.txt"
	want := int64(1)
	got, err := getSize(path, false, false)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetSizeNonExistentFile(t *testing.T) {
	path := "testdata/file2.txt"
	_, err := getSize(path, false, false)
	require.Error(t, err)
}

func TestGetSizeDirWithOneFile(t *testing.T) {
	path := "testdata/dir1"
	want := int64(2)
	got, err := getSize(path, false, false)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetSizeDirWithTwoFiles(t *testing.T) {
	path := "testdata/dir2"
	want := int64(5)
	got, err := getSize(path, false, false)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetSizeDirWithHiddenFiles(t *testing.T) {
	path := "testdata/dir3"
	want := int64(0)
	got, err := getSize(path, false, false)
	require.NoError(t, err)
	require.Equal(t, want, got)

	want = int64(4)
	got, err = getSize(path, false, true)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetSizeDirNonRecursive(t *testing.T) {
	path := "testdata"
	want := int64(1)
	got, err := getSize(path, false, false)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetSizeDirRecursive(t *testing.T) {
	path := "testdata"
	want := int64(12)
	got, err := getSize(path, true, true)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestHumanReadableSizeB(t *testing.T) {
	size := int64(1)
	want := "1B"
	got := humanReadableSize(size)
	require.Equal(t, want, got)
}

func TestHumanReadableSizeKB(t *testing.T) {
	size := int64(1 << 10)
	want := "1.0KB"
	got := humanReadableSize(size)
	require.Equal(t, want, got)
}

func TestHumanReadableSizeMB(t *testing.T) {
	size := int64(1 << 20)
	want := "1.0MB"
	got := humanReadableSize(size)
	require.Equal(t, want, got)
}

func TestHumanReadableSizeGB(t *testing.T) {
	size := int64(1 << 30)
	want := "1.0GB"
	got := humanReadableSize(size)
	require.Equal(t, want, got)
}
