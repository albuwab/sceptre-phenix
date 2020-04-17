// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package tmpl generated by go-bindata.// sources:
// tmpl/templates/minimega_script.tmpl
package tmpl

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _minimega_scriptTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\x4d\x6f\x1a\x31\x10\xbd\xf3\x2b\x46\x28\x07\x38\x60\xf5\x1c\xa9\xa7\xa4\x6d\x90\x1a\x1a\x95\x34\xf7\x91\x3d\xbb\xb8\x78\xc7\xc6\x1f\x04\x84\xf8\xef\x95\xd7\x0b\xd9\x25\x51\x9a\xdc\x6c\xcf\x9b\x37\xcf\xef\xd9\x8c\x0d\x05\x87\x92\xe0\x70\x00\xf1\x6d\xe7\xc8\xeb\x86\x38\x2e\xb0\x21\x38\x1e\x47\x1c\x60\x93\x28\x91\xe6\x1a\xa2\x4f\x34\x1a\x1d\x0e\xa0\x2b\x10\xbf\x13\xff\xb4\x12\x4d\x07\x52\x64\x66\x2b\x1b\x22\xa0\x31\x79\x8f\x4a\x95\xbd\xc9\xa0\xbc\xca\x8d\xc4\x2a\xe3\xf3\xd2\x23\xd7\x04\xe2\xd1\x3a\x6b\x6c\xbd\x17\x0b\xab\x28\xe4\x62\xc7\xff\x83\x98\x3c\x1a\xb1\x64\x74\x61\x65\x63\x57\xba\xaa\xb4\x0f\xf1\xd6\xeb\x2d\xc1\xf5\x57\xd0\xac\x68\x07\xe2\x0e\xbd\x7a\x46\x4f\xa2\x2d\x04\xf8\x92\xd1\x4a\x87\x35\x84\x53\xfb\xb0\x57\xcc\x1b\xac\xf3\x05\xdb\xf3\xf3\x90\xf6\xd6\xe7\xd1\x77\x36\x44\x2e\x3e\x40\x27\xab\x8e\x30\x31\xc4\x20\xe6\xfc\x97\x64\xd4\x96\xc3\xb4\x37\x4d\xb7\xa7\x1f\xe2\x84\xe3\xf1\xfa\x52\x53\xdb\xfd\x80\x3e\xea\xcc\x9c\xc7\x56\xda\x50\x68\xa3\xf9\xae\x0d\x15\xc0\xc9\xa5\xce\xcc\x9e\xad\xd2\x10\x7a\xd8\x36\x20\x2d\x57\xba\x3e\x65\xc5\x04\x93\x62\xd4\x95\x58\xca\x15\xa9\x94\x49\x5f\x49\x9a\xc2\x78\x9c\x69\xce\xfd\x10\x3a\x70\x16\xf0\x01\x82\x0b\x35\x2f\x3c\x5b\xe9\x52\xb9\xc5\x39\xa8\xa7\x9b\x87\x3f\xc3\x61\xd2\xa5\x21\xe4\x15\xa2\xa1\xc6\xfa\xfd\x10\x74\x5f\xce\x86\xb2\x7b\x99\xbf\xf9\x8e\xde\x7b\x63\x2f\x34\x6d\xa2\x83\x61\xb7\x3a\xac\x6f\x4a\x71\xf2\xbf\x80\xa7\x27\x37\x4c\xa0\x4f\xf0\x96\x08\x86\x5f\x45\x57\x40\x9b\x1e\xfc\xd7\xf2\x71\xef\x08\xc6\x46\x73\xda\x5d\x44\xb6\xa1\x26\xcd\xd0\xb9\xdc\x3e\xdb\xd6\x08\x9b\x9d\x79\x3b\x14\xa6\x62\xd0\x82\xe2\xb3\xf5\x6b\x31\xe7\x48\xbe\x42\x49\x9d\x92\x42\x6b\x30\xb1\x5c\x0d\x9c\x7c\xba\x6f\xc7\x97\xbf\xf3\xfe\x23\xf8\x17\x00\x00\xff\xff\x9b\x8c\x9b\xac\x5e\x04\x00\x00")

func minimega_scriptTmplBytes() ([]byte, error) {
	return bindataRead(
		_minimega_scriptTmpl,
		"minimega_script.tmpl",
	)
}

func minimega_scriptTmpl() (*asset, error) {
	bytes, err := minimega_scriptTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "minimega_script.tmpl", size: 1118, mode: os.FileMode(420), modTime: time.Unix(1587055810, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"minimega_script.tmpl": minimega_scriptTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"minimega_script.tmpl": &bintree{minimega_scriptTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
