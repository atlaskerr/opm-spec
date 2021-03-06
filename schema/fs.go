// Copyright © 2019 Open Package Management Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "esc -private -modtime=1546544639 -pkg=schema -include=.*\.schema.json$ -ignore=test-fixtures/* ."; DO NOT EDIT.

package schema

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return []os.FileInfo(fis[0:limit]), nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/namespace/create-namespace.schema.json": {
		name:    "create-namespace.schema.json",
		local:   "namespace/create-namespace.schema.json",
		size:    621,
		modtime: 1546544639,
		compressed: `
H4sIAAAAAAAC/7xSS07DMBDd5xSWYZnEiUgTyJYtQlxhmg6pS/zReJBAVe6O8mtL03V39pv3kz3HSAgh
H/VO1kLumX2olXIerYfmC1oMqXYqNHs0oCwYDB4aPJ+ShhAYZTzZTMTFqlbqEJxNJjR11KodwScnWTVb
PsxCT84jscYgazFWGtAOtthdIiMVmJHsxy3FTEmPedyv4GHCvx6HdoFJ21b+m/fnWx9Ha43bHrDhRXOi
SELv3rTRfFXUwI8230bWIn8uirIqiqx6qrKXzSYv8/IywWg7M7Nbwdoytkjr5MCOoMW7hkenBpI1dyPn
ddwB8b4sxfyp1w/XR38BAAD//69tkQVtAgAA
`,
	},

	"/namespace/list-namespaces.schema.json": {
		name:    "list-namespaces.schema.json",
		local:   "namespace/list-namespaces.schema.json",
		size:    1380,
		modtime: 1546544639,
		compressed: `
H4sIAAAAAAAC/9SUy07yQBTH9zxFM9+3LEyJhWp3hBBDgiwMujEuhnIsBzuXzBwSieHdTS9cgi2ycIHb
87+cX9qZ+Wx5nsf+44LFHlsSGRdzrg0oI5J3kYLroOYuWYIUXAkJzogE2hk6Yn4ZLcVdPOZ85bRql9OO
tilfWPFG7SCqav5VQWO1AUsIjsVegZFP9zuOp7mABPJk1thykDMxh6xWylXaGMjB9XwFCbFTy9b/1pfj
/dTmyKJKL2mzYPRQrxU1VUrxgXItWex1b8OwH4VhEN1EwV2v1+13+35tBFUVCfxzlKgIUrCXYk5Q4pVj
OhK0bvzXoIptLzVaLg+Gs/HziPkN8mz0+DCeDmbj6T2rsbz6v3MgHGkrUvgTH7sgfXKwuCbQ1hns5ut+
bNy7hLViszMVxaWNEVJWOCboyJseHqxKPtmybX0FAAD//5/7w6hkBQAA
`,
	},

	"/namespace/namespace.schema.json": {
		name:    "namespace.schema.json",
		local:   "namespace/namespace.schema.json",
		size:    1001,
		modtime: 1546544639,
		compressed: `
H4sIAAAAAAAC/8SSTU/yQBSF9/0Vk3nfZWHaWKh2RwgxJEqMQTfGxVCuZZD5yMwl0ZD+d9NSoF8bN7js
mfPcc3JvDx4hhP4Xa5oQukE0LmFMG1CGp588AzcUmrl0A5IzxSU4w1Og/pE66icyYWzrtBoc1aG2GVtb
/oGDIK4m/KtAY7UBiwIcTUjZoFB3fAW7ulJaOSJY9dRHVJbhIfTzjly84LeBop1DK1RGG+/55Sv3vS6j
V1tI8cScLbRYQqtjf8wFsWD0VO8VtjjJv4TcS5qQ8DaKxnEUBfFNHNyNRuE4HNdLSaEqZ9DXVSiEDGx/
8oOQ4vrJDjnu28cEVU56a95vMl3OX2fUb6rL2fPjfDFZzhf3tcO9+7/Yu0NteQZ/tYAy/MXB+lrZ3rlA
5y/OvZ8AAAD//yXw2tXpAwAA
`,
	},

	"/": {
		name:  "/",
		local: `.`,
		isDir: true,
	},

	"/namespace": {
		name:  "namespace",
		local: `namespace`,
		isDir: true,
	},

	"/repository": {
		name:  "repository",
		local: `repository`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	".": {},

	"namespace": {
		_escData["/namespace/create-namespace.schema.json"],
		_escData["/namespace/list-namespaces.schema.json"],
		_escData["/namespace/namespace.schema.json"],
	},

	"repository": {},
}
