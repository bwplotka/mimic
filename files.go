package mimic

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/pkg/errors"
	"github.com/pmezard/go-difflib/difflib"
)

// FilePool is a struct for storing and managing files to be generated as part of generation.
type FilePool struct {
	Logger log.Logger

	path []string

	m map[string]string
}

// Add adds a file to the file pool at the current path. The file is identified by filename.
// Content of the file is passed via an io.Reader.
//
// If the file with the given name has already been added at this path the code will `panic`.
// NOTE: See mimic/encoding for different marshallers to use as io.Reader.
func (f *FilePool) Add(fileName string, r io.Reader) {
	if filepath.Base(fileName) != fileName {
		Panicf("")
	}

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

func (f *FilePool) sortedPaths() []string {
	var paths []string
	for k := range f.m {
		paths = append(paths, k)
	}
	sort.Strings(paths)
	return paths
}

func (f *FilePool) write(outputDir string) {
	for file, contents := range f.m {
		out := filepath.Join(outputDir, file)
		if err := os.MkdirAll(filepath.Dir(out), 0755); err != nil {
			PanicErr(errors.Wrapf(err, "create directory %s", filepath.Dir(out)))
		}

		// TODO(bwplotka): Diff the things if something is already here.

		_ = level.Debug(f.Logger).Log("msg", "writing file", "file", out)
		if err := ioutil.WriteFile(out, []byte(contents), 0755); err != nil {
			PanicErr(errors.Wrapf(err, "write file to %s", out))
		}
	}
}

func (f *FilePool) diff(f2 *FilePool) string {
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
				PanicErr(errors.Wrap(err, "diffing via difflib"))
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
				PanicErr(errors.Wrap(err, "diffing via difflib"))
			}
			out += o
		}
	}
	return out
}
