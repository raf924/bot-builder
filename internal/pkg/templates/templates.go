// Code generated by vfsgen; DO NOT EDIT.

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Templates statically implements the virtual filesystem provided to vfsgen.
var Templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 12, 15, 19, 53, 22, 898959900, time.UTC),
		},
		"/bot.go.tpl": &vfsgen۰CompressedFileInfo{
			name:             "bot.go.tpl",
			modTime:          time.Date(2020, 12, 15, 19, 50, 2, 144900700, time.UTC),
			uncompressedSize: 988,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x84\x52\xcf\x8e\xda\x3c\x10\xbf\x47\xca\x3b\xcc\xe7\xc3\x27\x07\x2d\x8e\xb4\x6a\x0f\xa5\xe5\x02\xab\xdd\xf6\xd0\x5d\xb4\xf4\xbe\x32\xc9\xc4\x58\x38\x9e\xc8\x18\x5a\x14\xf1\xee\xd5\xc4\x2c\x45\xd5\xb2\xbd\xc4\x1e\xcf\xef\xdf\x8c\xd2\xf7\x63\x28\x47\x86\xe2\xa1\xc3\x09\x18\x1b\xd7\xbb\x95\xaa\xa8\x2d\x83\x6e\x3e\xdd\x7e\x28\x57\x14\xc7\xab\x9d\x75\x35\x86\xb2\x6a\xeb\xcb\x5a\xad\x28\xce\xc9\x37\xd6\x8c\x4a\x18\x1f\x8f\x79\x56\x96\x86\x26\x06\x3d\x06\x1d\x11\x0c\x81\xc1\xf8\xb6\xe6\x15\x6c\xdf\x83\x9a\x51\x7c\x46\xa7\x0f\xc0\x8a\x7d\x0f\x41\x7b\x83\xa0\xe6\x6d\xfd\x9d\xea\x9d\xc3\xed\x55\x27\x66\xbf\xb2\xd0\xd7\xc3\x35\xcf\x3a\x5d\x6d\xb4\x41\x68\xb5\xf5\x5c\xdb\xb6\xa3\x10\x41\xe6\x99\x68\x9c\x36\x22\xcf\x5e\x40\xfc\x65\x2c\xae\x3b\x0f\x60\x75\x82\xa0\xaf\xf9\x4d\xbc\x39\x63\xd9\x6d\x0c\x9f\x22\xcf\xce\x9b\x82\x77\xa0\xd5\x80\x38\x31\x84\xa1\x6e\x63\x94\xf5\xe5\x41\xb7\x4e\xed\x6f\xf9\xcd\x11\xc7\x15\xb4\xe5\x6f\xb4\x2d\x8a\x3c\x2b\x78\xa6\x66\xe7\xab\x61\x40\x59\x40\x9f\x67\x49\xe9\xde\x3a\x84\xc9\x14\x78\x4a\xb5\x8c\xc1\x7a\x23\x45\x6a\x89\x1b\xa0\xad\x7a\xc0\x88\x7e\x2f\xc5\xd7\xf9\xec\xe9\xc7\xcb\xfc\xe9\xf1\xfe\xdb\x83\x28\x6e\x40\x2c\x74\x5c\x43\x24\x60\x6b\x48\x0c\x68\xac\xc3\x09\xe0\x2f\x28\xa9\x8b\xe5\xba\xe2\xd4\xa9\xa5\x18\x26\x8a\x3c\x1b\x8c\x16\x3a\x6c\x51\x16\x79\x66\x1b\x18\x5d\x04\x99\x4e\x41\x08\x0e\xd7\x69\x6f\x2b\x29\x1e\xe9\x52\x99\xe9\xc7\x3c\x6b\x6e\x00\x43\xe0\xd0\xb4\x55\x4f\x1d\x7a\x79\x21\x91\x34\xb9\xff\xdf\x14\xbc\x75\x7f\xc4\x30\x84\xc4\xdf\xeb\xf0\xaa\x7a\x5e\xb9\x4a\xc7\x99\x3c\x1d\xc6\x52\x8f\xf8\xf3\x0e\x2b\xaa\x31\xc8\xa6\x50\xe9\x2a\xff\x4f\xe4\xe2\xf3\xfb\x36\xeb\x6a\x46\x91\x53\xae\x28\xb2\xd0\x8c\xa2\x3c\x31\xf3\xac\xa1\xc0\x94\xb3\xdb\x00\x56\xcb\xa8\x43\x94\xff\x12\xfe\x32\x4e\xe8\x3b\xf2\xc3\x0e\x1d\x19\xb5\x08\xd6\x47\xe7\xa5\x78\xc6\x8a\xbc\xc7\x2a\x5a\x6f\x78\x5f\xfc\x03\xa8\xa5\x43\xec\xe4\x47\x18\x41\x2a\x19\x53\x27\xb1\xe3\xef\x00\x00\x00\xff\xff\x49\xa6\xfc\x53\xdc\x03\x00\x00"),
		},
		"/connector.go.tpl": &vfsgen۰CompressedFileInfo{
			name:             "connector.go.tpl",
			modTime:          time.Date(2020, 12, 15, 19, 53, 22, 892959200, time.UTC),
			uncompressedSize: 977,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x84\x92\x51\x6f\xda\x3e\x14\xc5\xdf\x23\xe5\x3b\xdc\xbf\x1f\xfe\x72\x10\xd8\x52\xb5\x97\xb1\xe5\x65\xe9\xda\xed\x05\xaa\x76\xef\xc8\x98\x1b\x63\x11\x7c\x23\x63\xd8\x10\xe2\xbb\x4f\x8e\x21\x4d\xb7\xb6\x7b\x33\xdc\x73\x7e\x3e\xe7\x3a\xa7\xd3\x04\xe4\xc8\x50\x38\xb6\x38\x05\x63\xc3\x7a\xbf\x14\x9a\xb6\xd2\xab\xfa\xe3\xcd\x07\xb9\xa4\x30\x59\xee\x6d\xb3\x42\x2f\xf5\x76\x35\xfc\x2d\x34\x39\x87\x3a\x90\xaf\xc8\xd5\xd6\x8c\x24\x4c\xce\xe7\x3c\x93\xd2\xd0\xd4\xa0\x43\xaf\x02\x82\x21\x30\x18\x5e\x27\xbf\xa1\x3d\x9d\x40\x7c\xa1\xf0\x88\x8d\x3a\xc2\x9b\xc4\xa8\xaa\x52\x02\x4b\xee\x59\xdc\x2a\xbd\x51\x06\x61\xab\xac\xcb\xb3\x3c\xb3\xdb\x96\x7c\x00\x9e\x67\x00\x00\xac\x6e\x94\x61\x97\xb3\xa1\x76\x63\x84\x75\xf2\xa8\xb6\x8d\x38\xdc\x5c\xff\xa7\xdd\xe5\xa4\x5d\x0d\xec\xd5\xe8\xb2\xdd\x18\xa9\xbb\xde\xb2\xdf\x43\xcf\x7d\xcf\xf1\x42\xba\x00\xf6\x47\xdb\x97\x83\xbf\x0b\xb2\x3c\x2b\x62\xad\x7a\xef\x74\xd7\x91\x17\x70\xca\xb3\x14\xe5\xce\x36\x08\xd3\x12\x62\x49\xf1\x14\xbc\x75\x86\xb3\x34\x62\x63\xa0\x9d\xb8\xc7\x80\xee\xc0\xd9\xb7\xaa\x9a\xcf\x66\x5f\xab\x1f\xf3\xc7\x45\x35\x9f\xdd\x7d\xbf\x67\xc5\x18\xd8\x83\x0a\x6b\x08\x04\x71\x1f\x90\x7c\x50\xdb\x06\xa7\x80\xbf\x40\x52\x1b\xe4\x5a\xc7\x2a\x69\x24\xa2\x8c\x15\x79\xd6\x5d\xf7\xa0\xfc\x0e\x79\x91\x67\xb6\x86\xd1\x20\x4e\x59\x02\x63\x31\x62\xab\x9c\xd5\x9c\xcd\x68\x48\x8e\xf6\x73\x9e\xd5\x63\x40\xef\x63\x74\xda\x89\x79\x8b\x8e\x0f\x10\x89\x19\xe7\xff\x95\xe0\x6c\xf3\x0c\x43\xef\x93\xff\xa0\xfc\x95\xaa\x5d\x2d\xd2\x07\xd9\xdb\xca\xae\x90\x98\xe1\xcf\x5b\xd4\xb4\x42\xcf\xeb\x42\xa4\x23\xff\x3f\xd9\x8a\x4f\xef\x5f\xb0\xd6\xd5\xf5\xe9\x62\xca\xfe\x1d\x23\xb4\x9f\xf0\x0b\xab\xbb\x78\xe8\x28\x7b\xac\xc7\xb0\xf7\xae\x43\xf6\xe1\x06\x4a\xf1\x14\x94\x0f\xfc\x5f\x69\x16\x50\xc2\xe7\xc9\xd0\x77\x4b\xae\x5b\xfe\xf9\x77\x00\x00\x00\xff\xff\x13\x4a\x91\x5f\xd1\x03\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/bot.go.tpl"].(os.FileInfo),
		fs["/connector.go.tpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
