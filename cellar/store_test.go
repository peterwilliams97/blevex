//  Copyright (c) 2016 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package cellar

import (
	"os"
	"testing"

	"github.com/blevesearch/bleve/index/store"
	"github.com/blevesearch/bleve/index/store/test"
)

func open(t *testing.T, mo store.MergeOperator) store.KVStore {
	rv, err := New(mo, map[string]interface{}{"path": "test"})
	if err != nil {
		t.Fatal(err)
	}
	return rv
}

func cleanup(t *testing.T, s store.KVStore) {
	err := s.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = os.RemoveAll("test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCellarKVCrud(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestKVCrud(t, s)
}

func TestCellarReaderIsolation(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestReaderIsolation(t, s)
}

func TestCellarReaderOwnsGetBytes(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestReaderOwnsGetBytes(t, s)
}

func TestCellarWriterOwnsBytes(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestWriterOwnsBytes(t, s)
}

func TestCellarPrefixIterator(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestPrefixIterator(t, s)
}

func TestCellarPrefixIteratorSeek(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestPrefixIteratorSeek(t, s)
}

func TestCellarRangeIterator(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestRangeIterator(t, s)
}

func TestCellarRangeIteratorSeek(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestRangeIteratorSeek(t, s)
}

func TestCellarMerge(t *testing.T) {
	s := open(t, &test.TestMergeCounter{})
	defer cleanup(t, s)
	test.CommonTestMerge(t, s)
}