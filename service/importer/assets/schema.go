// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// service/importer/assets/schema.json
package assets

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

var _serviceImporterAssetsSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\x4d\x6f\x9c\x3c\x10\xbe\xf3\x2b\x2c\x27\xb7\xf7\xdd\x6e\x7a\xa9\xd4\xbd\xb5\xb7\x4a\x55\x13\x29\xb7\x46\xb4\x72\x60\xd8\x38\xc5\x1f\xf1\x47\xa4\x6d\xc4\x7f\xaf\x30\x81\x05\x6c\x03\x49\xb6\x51\xd5\x2d\x27\xf0\x78\xc6\xf3\x3c\x33\x7e\x6c\x1e\x12\x84\xf0\xa9\xce\x6e\x80\x11\xbc\x41\xf8\xc6\x18\xb9\x59\xaf\x6f\xb5\xe0\xab\x66\xf4\x8d\x50\xdb\x75\xae\x48\x61\x56\x67\xef\xd6\xcd\xd8\x09\xfe\xbf\xf6\x33\x3b\x09\xb5\x93\xb8\xbe\x85\xcc\x34\x63\x0a\xee\x2c\x55\x90\xe3\x0d\xba\x4a\x10\x42\x08\x73\xc2\x00\x27\x08\xa5\xce\x4e\xf2\x9c\x1a\x2a\x38\x29\x2f\x94\x90\xa0\x0c\x05\x8d\x37\xa8\x20\xa5\x06\x37\x41\xf6\x87\x1f\x7a\x21\xda\xaf\xde\xc2\xda\x28\xca\xb7\x6e\x61\x37\xce\x28\xff\x0c\x7c\x6b\x6e\xf0\x06\xbd\x75\x83\x55\x63\xc3\xa4\xa4\x44\x4f\x84\x18\x4c\xce\x41\x67\x8a\xca\x3a\xcd\xa5\x2e\x0a\xa4\xd0\xd4\x08\xb5\x5b\xea\x01\xf7\xc0\xcd\x20\x25\xc1\xe1\xbc\xe8\x78\xab\x9f\x87\x7d\x08\x6e\xcb\x12\xb7\xce\xce\xd6\xbd\x85\x2b\xd1\xd9\xe6\x18\xef\x26\x4a\x62\x0c\x28\x7e\xe1\x17\xa0\x9b\xf2\xed\xea\x6c\xf5\x9e\xac\x7e\x7e\x58\x7d\xfd\xbe\x4a\xff\x3b\xf5\x66\xd4\xdd\xa4\xa0\x46\x81\x4f\xd6\x39\x14\x94\xbb\xb5\xf5\xda\xc1\xc5\x83\xb9\x55\x12\x7a\x6f\xdf\xd2\x01\x5b\x86\xe8\x1f\xc7\x43\x56\x8d\xf6\xf9\x5c\xe5\x20\x81\xe7\xc0\xb3\x61\x56\x31\xdc\x8b\x30\xcf\xe0\x9d\xc3\x1a\xc1\xd9\x65\xba\x1b\x56\xc1\x13\x91\xce\x42\x19\xd9\x42\x9f\x9a\xd4\x23\xa3\x1a\x90\x91\x09\x5e\xd0\xad\x55\x64\xbc\x9b\x67\x33\x4a\xda\x60\x2e\x14\xee\xcd\xda\x0b\x93\x24\x8a\x30\x30\xa0\xfe\xc6\xde\x9c\xc8\xe5\x49\xf9\x4c\x97\x74\xbf\xd4\x68\x38\xf5\x62\x04\x8e\x86\x81\x7d\x74\x4c\x04\xc1\xf4\xa5\xb8\xff\x54\xe3\xd5\xa2\xe7\xc0\x8b\xc3\x0a\xd9\xb0\x36\x17\xf3\x5a\x88\x12\x08\x5f\x16\xf4\xd1\x69\x49\x92\xbe\x77\x7d\x1c\x71\xcb\x82\xb5\x71\xd6\xcb\xb8\x27\x42\xf8\x8b\x65\xd7\xa0\x62\xd6\x8f\x8f\x30\x22\xe6\xf3\xa6\xbb\x02\x46\xaf\x01\x50\x53\x94\x82\xd8\xd2\xd4\x78\x2e\x63\xa4\x27\x53\xdf\x4f\x3d\x76\x0e\x26\xa1\xc1\x2d\x80\x85\x35\xd2\x1a\xdd\xa2\x48\xf7\x8a\x1b\x91\xda\x40\x97\x47\xdb\xb0\xd7\x29\x13\xdd\xbc\xc8\x9d\x72\x97\xe6\x22\x5d\xef\xc9\x62\x30\x96\x86\x4c\xc1\xa1\x82\xb5\x04\x46\x30\x1d\x99\x92\xe6\xc4\x90\xa3\x51\x52\x07\x36\x12\x6f\x79\x33\x75\x2b\xbc\x40\x36\x2a\xff\x6e\xff\x9b\x75\xa3\x5f\xe8\x3f\x56\x34\x02\x05\x7a\xca\x2e\x0f\x92\xdb\xbb\xa2\x1d\xec\x72\x1b\x61\xad\xb9\x6d\x4e\xe3\x1e\xe8\x85\xff\x03\x3a\x66\x24\x13\x8c\x11\x9e\x3f\x87\x4c\xa2\xb6\x9e\xca\xf9\x57\x4d\x34\x73\xdd\x44\xa3\x2b\xe7\x60\x79\xa2\x14\xd9\xf9\xfa\x44\x0d\xb0\x88\x52\x4c\x6f\xdb\xf8\x16\x4a\x83\x10\xef\x45\x69\x99\x27\x4a\xaf\x84\xd2\x72\x7a\x67\xe1\xd3\x23\x56\xa3\xac\xaf\xbd\xaf\x4d\x44\xa1\x04\xfb\x47\x86\x14\xca\x3f\xdf\x8f\x84\x86\xa1\x04\xd6\x7f\xa3\x49\xf5\x2b\x00\x00\xff\xff\x2f\x0e\x3a\x0b\xb5\x13\x00\x00")

func serviceImporterAssetsSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_serviceImporterAssetsSchemaJson,
		"service/importer/assets/schema.json",
	)
}

func serviceImporterAssetsSchemaJson() (*asset, error) {
	bytes, err := serviceImporterAssetsSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "service/importer/assets/schema.json", size: 5045, mode: os.FileMode(420), modTime: time.Unix(1543590884, 0)}
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
	"service/importer/assets/schema.json": serviceImporterAssetsSchemaJson,
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
	"service": &bintree{nil, map[string]*bintree{
		"importer": &bintree{nil, map[string]*bintree{
			"assets": &bintree{nil, map[string]*bintree{
				"schema.json": &bintree{serviceImporterAssetsSchemaJson, map[string]*bintree{}},
			}},
		}},
	}},
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
