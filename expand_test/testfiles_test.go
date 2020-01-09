// helper functions to create, delete files during testing.
package expand_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// TestFilesNew creates a temporary directory for testing.
func testFilesNew() tft {
	p, err := ioutil.TempDir("", "tf")
	if err != nil {
		panic(err)
	}
	return tft(p)
}

type tft string

// Path returns the absolute path of 'path' in the test directory.
func (e tft) Path(path string) string {
	if path == "" {
		return ""
	}
	return filepath.Join(string(e), path)
}

// MustRemove removes the file at 'path' from the test directory.
func (e tft) MustRemove(path string) {
	ap := e.Path(path)
	err := os.Remove(ap)
	if err != nil {
		panic(err)
	}
}

// MustRemoveAll removes the test directory.
func (e tft) MustRemoveAll() {
	err := os.RemoveAll(string(e))
	if err != nil {
		panic(err)
	}
}

// MustCreate create a file at 'path' with content 'text' in the test directory.  
func (e tft) MustCreate(path, text string) {
	if path == "" {
		return
	}
	p := e.Path(path)
	d := filepath.Dir(p)
	err := os.MkdirAll(d, 0700)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(p, []byte(text), 0600)
	if err != nil {
		panic(err)
	}
}
