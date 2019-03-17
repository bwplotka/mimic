package gocodeit

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v2"

	"github.com/pmezard/go-difflib/difflib"
)

type Files struct {
	path []string

	m map[string]string
}

func (f *Files) Add(fileName string, r io.Reader) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		Panicf("failed to output: %s", err)
	}

	output := filepath.Join(append(f.path, fileName)...)

	// Check whether we have already written something into this file.
	if _, ok := f.m[output]; ok {
		Panicf("filename clash: %s", output)
	}
	f.m[output] = string(b)
}

func (f *Files) sortedPaths() []string {
	var paths []string
	for k := range f.m {
		paths = append(paths, k)
	}
	sort.Strings(paths)
	return paths
}

func (f *Files) write(outputDir string) {
	for file, contents := range f.m {
		out := filepath.Join(outputDir, file)
		if err := os.MkdirAll(path.Dir(out), 0755); err != nil {
			Panicf("failed to create directory: %v", err)
		}

		if err := ioutil.WriteFile(out, []byte(contents), 0755); err != nil {
			Panicf("failed to write file: %v", err)
		}
	}
}

func (f *Files) diff(f2 *Files) string {
	var out string

	for _, p := range f.sortedPaths() {
		if f.m[p] != f2.m[p] {
			o, err := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
				A:        difflib.SplitLines(f.m[p]),
				B:        difflib.SplitLines(f2.m[p]),
				FromFile: p,
				ToFile:   p,
				Context:  3,
			})
			if err != nil {
				Panicf("diffing via difflib failed: %v", err)
			}
			out += o
		}
	}
	for _, p := range f2.sortedPaths() {
		if _, ok := f.m[p]; !ok {
			o, err := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
				A:        difflib.SplitLines(f2.m[p]),
				B:        []string{},
				FromFile: p,
				ToFile:   p,
			})
			if err != nil {
				Panicf("diffing via difflib failed: %v", err)
			}
			out += o
		}
	}
	return out
}

func UnmarshalSecretFile(file string, in interface{}) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		Panicf("read file from: %v", err)
	}

	if err := yaml.Unmarshal(b, in); err != nil {
		Panicf("failed to unmarshal file from: %v", err)
	}
}
