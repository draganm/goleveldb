// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//go:build js

package storage

import "errors"

var errNotSupported = errors.New("leveldb/storage: not supported on js/wasm")

func newFileLock(path string, readOnly bool) (fl fileLock, err error) {
	return nil, errNotSupported
}

func rename(oldpath, newpath string) error {
	return errNotSupported
}

func isErrInvalid(err error) bool {
	return false
}

func syncDir(name string) error {
	return errNotSupported
}
