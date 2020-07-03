// Code generated by go-bindata. DO NOT EDIT.
// sources:
// contracts/crypto.cdc (5.958kB)

package internal

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _contractsCryptoCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\x5b\x4f\x23\xb7\x17\x7f\xcf\xa7\x38\x7f\x9e\x12\xfd\x43\x60\x43\x97\x87\x48\xb3\x88\x02\x2d\x11\x6d\x85\xc8\x2e\x7d\x40\x88\x35\x33\x27\x13\x2b\x83\x27\xb2\x3d\x81\x14\xf1\xdd\xab\xb9\xfb\xd8\x1e\x08\x5b\xaa\xe6\x01\x26\x39\xf7\xdf\xb9\x78\x7c\x7a\xab\xec\x1e\x94\x96\x59\xa8\x81\x0b\x8d\x72\xce\x42\x84\x19\x8f\x05\xd3\x99\xc4\x6b\x94\x7c\xce\x51\x02\x3c\xf7\x7a\x00\x00\x39\xfb\x3c\x13\xb0\xce\x09\x9b\x7e\xf1\x5b\xfe\x51\xb5\xc4\x04\x6e\xa6\x42\xdf\x0e\x1b\x8a\x66\xf1\x04\x66\x5a\x72\x11\x0f\x09\x3b\x46\xa7\x4c\x33\x87\x7f\x95\xdd\x27\x3c\xbc\xc0\x8d\x43\x69\x6c\x1c\x27\x71\x2a\xb9\x5e\x3c\xb8\x8a\x17\x4c\x2d\x1c\x72\x41\x1d\x4c\xe0\xe7\x34\x4d\x7a\x2f\xbd\x22\xe6\x30\x15\x5a\xb2\x50\xc3\x89\xdc\xac\x74\x6a\xc6\x57\xc1\x31\x73\xcc\xc1\xb3\xe9\x25\x24\xa8\x41\xb0\x07\x6c\xcc\x34\x54\x2e\xb8\xee\x9b\xa4\x81\x21\x5a\x84\x82\xc9\x7c\x94\x33\x40\x50\xa8\x68\x88\x2f\xbd\xf2\x6f\xf1\x6f\x6f\x0f\xce\x4e\x4e\x67\xc7\x77\x97\xe3\xcf\x87\xc0\x15\x9c\x25\x09\x5f\x69\x1e\xc2\x49\x26\xd7\x08\xa7\x3c\xe6\x9a\x25\xad\xa7\xd0\xba\xda\x2f\x24\x07\x90\x0a\xd0\x0b\x84\x3f\xa6\xb3\xaf\x70\xb9\x9b\xeb\x09\x73\xd9\x9e\x19\x44\x6b\x64\xe2\x89\xda\xf2\x65\x86\xe1\x6a\xfc\xf9\x70\xf9\xe9\x9f\x39\xa4\x1a\x35\x5d\xfe\x34\x86\xba\x9d\x32\xb2\x75\x6e\x26\xfe\x3f\x48\xd4\xec\xfc\x78\x7c\x57\xa5\x69\x86\x61\x1e\x7b\xee\x12\x17\xb1\x81\xc1\x18\xfa\xb3\xf3\xe3\xdd\xf1\x00\x1e\xb9\x5e\x00\x83\xf1\xe7\xc3\xdd\x7b\xae\x21\xe2\x31\x2a\x4d\x40\xa8\x15\x4e\x68\x68\xa6\xc1\x83\x37\x0d\x1e\x94\x06\x0f\xb6\x34\x78\xd0\x69\xd0\x40\xfa\xb2\x6e\x50\x0f\xca\x76\xf3\x3a\x0c\xde\x1e\xee\xc8\x6e\x93\x20\x67\x24\x6c\xa9\xc6\x9b\xcc\x46\x19\x04\xad\xbb\x2e\x9b\x6b\x01\x02\x8f\x59\x7f\x3d\x18\x68\x5d\xe0\xe6\x37\xae\xf4\x99\xd0\xd2\x07\xd8\x12\x37\x53\x11\xe1\xd3\x04\xa6\x42\xbf\x06\xe7\xa5\xe3\x6a\xcd\x64\x8d\x3c\x9a\x3c\x9b\xf9\x11\x79\xbc\xd0\x13\xf8\xf6\x0b\x7f\x3a\xfc\xc9\x21\x73\x75\x85\xeb\x74\x89\x51\x35\x2c\x69\x1e\x08\x4a\xc4\xf3\x21\x21\xf9\xdc\xa6\x1c\xaf\xf9\x4c\x39\xa9\xc3\x94\x66\x7b\x5b\xff\xee\xcd\x7b\xed\x2f\x04\x8d\xeb\x3f\x5c\x1c\xc4\x7d\x08\x68\x38\x2e\x7b\x19\x03\x04\x55\x30\x2e\x43\x13\x08\x04\x6d\x50\xdb\x96\x56\x7d\x74\x15\x54\xc9\xd7\x45\x22\x51\x68\xc9\x51\x4d\xe0\xc6\x2c\xc0\x5b\x2b\xa1\x5e\xa0\x2a\x51\x08\xe0\xe6\xd6\xf0\xa1\x79\xdc\xdb\x83\xe3\x28\x52\xc0\x40\xe0\x63\x8e\x65\x39\x5b\xf2\xa1\x1e\xf3\x35\x0a\x3b\xca\xfa\x95\x81\x45\x11\xad\xa0\xbb\x7f\xb1\x50\xda\x5a\x98\xd8\x2d\x48\xa4\xcc\x1e\xcc\x1b\xdc\x00\x60\x94\xa0\x88\xf5\xc2\x61\xc7\x42\x4f\x40\xd4\xd2\xc0\x80\xb4\x47\xfd\x34\x74\x78\x8c\xf0\x57\xfe\xf0\x3d\x10\x2c\xba\x21\x30\x61\x28\xff\xbb\x74\xa3\x67\xe6\x2c\x51\x48\x18\x06\x9d\xd5\x30\x62\xab\x15\x8a\xa8\x5f\x04\x4f\xd9\x24\xea\x4c\x8a\x12\x96\x8e\x7a\xb9\x2a\x58\x54\x51\x23\x79\xc1\x30\x6d\x94\x0b\x2f\xc0\x01\x3e\x07\xae\x01\x9f\xb8\xd2\x6a\x44\x85\xcb\xd6\x58\xe2\x46\x01\x93\x08\x2c\x79\x64\x1b\x55\xd9\xc5\x68\x08\xf7\x59\xa1\x6f\x03\x0b\xb6\x46\xf8\xde\x84\xf8\x1d\xe6\x1c\x93\x08\x14\x6a\xd0\x29\x68\x99\xa1\x53\x95\x31\xea\x3e\x19\x65\x56\xc1\x1c\x59\x2d\xc2\xe7\x6d\xbd\x7c\xf1\x16\x8c\x25\x60\x40\x24\x78\x42\x48\x2f\x3d\x1f\x90\xa6\xca\x9b\xda\x56\x57\x27\xfe\xce\xe4\xf2\x35\x5c\x41\x96\x50\x94\x20\x45\x29\x2a\x10\xa9\x86\x08\x13\xd4\x08\xdc\xed\xd2\x92\xdf\x82\xe4\xe3\x30\xb0\xe2\x37\xbf\xe5\xbd\x15\x66\x52\xa2\xa8\x5a\x35\x78\x0b\x0a\xb0\x8a\xb4\x65\x79\x47\x7b\x9a\x26\x47\x5b\xf5\x2a\x91\xd8\xbe\x71\x89\xd8\x96\x5d\x4c\x64\xb6\x68\x69\x52\xe2\x40\x3a\xda\xa8\x9b\x3a\xd7\x5c\x5d\xb3\x84\x5b\x53\xb9\x79\xc7\x99\xa1\x6e\x8f\x8f\xe6\xe5\xea\x76\xe8\x70\x93\x8b\x9c\x39\x78\xf3\x43\xd9\x1e\xb8\x6b\x26\x61\x9d\x5b\xfd\xb3\x88\x46\xd5\x03\x1b\x02\xd8\x1f\xed\xbb\xc3\x59\x21\x8a\x8b\x22\x2b\x3c\xcc\xcf\xb3\xe7\xa9\xd0\xa5\xe6\x17\x08\xe0\xd9\x6a\xa1\x79\x2a\xdb\x00\x80\x0b\x12\x8d\xed\x0a\x54\xd7\x1a\xa1\x72\xe6\xba\x87\xca\xb6\xe1\xaa\xf4\xd2\x95\xe0\xf3\x56\xe9\xe8\xfd\x6d\x00\x6d\xa3\xbb\xf3\x17\xdc\xa1\x60\x3b\xc9\x95\xe1\xe5\x82\x95\xed\xcc\x12\x89\x2c\xda\xc0\x3d\xa2\x28\x10\xf3\xbb\x4d\xa0\xbc\x71\xa3\xb8\x85\xa3\xa3\xd2\xab\x8f\x73\xfc\x0a\xc3\x54\x46\x16\xba\x8f\x4c\x75\xb8\xb9\x85\x8f\x41\x59\xe5\x3e\x63\xbf\x62\x39\x00\x59\xa8\x33\x96\xe4\x06\x5d\xb6\xea\xc4\xb7\xa7\x8b\xc7\xd2\x56\xd5\x52\x26\xa0\x9a\xb2\x5e\xd8\x97\xb8\x31\x5e\xf0\x3e\xbe\x22\xd0\xac\xf8\xd7\xca\xf6\x7f\xe5\xaa\xa3\xbd\xda\xd4\xdb\x9d\x91\xbd\xcd\x21\x19\x69\x37\x3b\x2d\x46\xcd\x93\x3b\x8d\xa0\xde\xf9\x54\xd6\xa2\xf4\x81\x71\x31\xc3\x15\x93\x4c\xf3\x54\x7c\x65\xf1\x37\x85\xd2\x2f\x68\x8e\x93\xf6\xd9\xcf\x6b\x4c\xe4\x1c\xe2\xe6\xeb\x6b\x23\x19\x3a\xd6\x48\x54\x81\xcb\x51\xac\x00\xfc\xea\xe8\x94\xcf\x15\x91\x5f\x46\x64\x6b\x50\x7f\xec\x13\xb5\xfe\xbc\xb7\x0e\xcc\x49\x0a\x01\xfd\xfa\xff\x22\x2a\xcf\x95\xc3\xff\xda\x41\x64\xbf\x04\xf0\x69\xb4\xbf\xed\x0d\xa4\x5d\xf7\xbc\xf7\x82\x6b\xaf\x0d\xe9\x19\x55\xdc\x51\xe8\x3d\xd3\x91\xf8\xf1\xbb\x5e\xdb\x35\xc6\xc5\xbe\x23\xe4\xfa\x5a\xd5\x51\xcc\x74\xb1\xd4\x70\x3b\x8d\x36\x81\x67\x67\xb5\x5a\x99\x28\x62\xdd\x4e\x60\x60\x1e\x64\x34\x94\x66\x5f\x1b\xb8\xc6\xc9\x7b\xe3\x54\x70\xcd\x59\xc2\xff\x42\x08\x53\xa1\x34\x13\x5a\x59\x4a\x8d\xd5\x63\xe0\x59\xae\x54\x6b\xb3\x9d\x96\x6d\x67\xe0\x53\xd0\xee\x0b\xdf\xd6\xd2\xf0\xee\x0c\x2c\x67\x9a\xf5\x5a\x40\xef\x80\xb5\x7c\x4d\xb7\x7d\x68\xb6\x64\xdd\x72\x07\x95\x1c\x15\xec\xc8\x34\x04\xb0\x93\x29\x94\xbb\xd7\xfb\xa3\xfd\x9d\xaa\x46\x5e\xfe\x0e\x00\x00\xff\xff\xfc\xa3\x11\xb8\x46\x17\x00\x00"

func contractsCryptoCdcBytes() ([]byte, error) {
	return bindataRead(
		_contractsCryptoCdc,
		"contracts/crypto.cdc",
	)
}

func contractsCryptoCdc() (*asset, error) {
	bytes, err := contractsCryptoCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contracts/crypto.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0xa7, 0x1a, 0x9b, 0xa3, 0x5f, 0x7a, 0xaf, 0x19, 0xf1, 0xa4, 0xa6, 0x77, 0xff, 0x2f, 0xfb, 0xff, 0x95, 0x6a, 0xd0, 0x91, 0x73, 0x98, 0xe9, 0x28, 0xc4, 0xba, 0xc5, 0x93, 0x6, 0xcf, 0x3a}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"contracts/crypto.cdc": contractsCryptoCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"contracts": &bintree{nil, map[string]*bintree{
		"crypto.cdc": &bintree{contractsCryptoCdc, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
