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
