// +build ignore

package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sdboyer/gps"
	"github.com/sdboyer/gps/pkgtree"

)

// This is probably the simplest possible implementation of gps. It does the
// substantive work that `go get` does, except:
//  1. It drops the resulting tree into vendor instead of GOPATH
//  2. It prefers semver tags (if available) over branches
//  3. It removes any vendor directories nested within dependencies
//
//  This will compile and work...and then blow away any vendor directory present
//  in the cwd. Be careful!
func main() {
	// Assume the current directory is correctly placed on a GOPATH, and that it's the
	// root of the project.
	root, _ := os.Getwd()
	srcprefix := filepath.Join(build.Default.GOPATH, "src") + string(filepath.Separator)
	importroot := filepath.ToSlash(strings.TrimPrefix(root, srcprefix))
	/* m := gps.SimpleManifest{
		Deps: []gps.ProjectConstraint{
			{
				Ident: gps.ProjectIdentifier{
					ProjectRoot: gps.ProjectRoot("github.com/a8uhnf/go_stack"),
				},
				Constraint: gps.NewBranch("master"),
			},
		},
	} */
	// params.Manifest = m
	// Set up params, including tracing
	params := gps.SolveParameters{
		RootDir:         root,
		Trace:           false,
		TraceLogger:     log.New(os.Stdout, "", 0),
		ProjectAnalyzer: NaiveAnalyzer{},
		Manifest:        ManifestYaml{},
		RootPackageTree: pkgtree.PackageTree{
			ImportRoot: importroot,
			Packages: map[string]pkgtree.PackageOrErr{
				"Hello": pkgtree.PackageOrErr{
					P: pkgtree.Package{
						Name: "go_stack",
						
					},
				},
			},
		},
	}
	// Perform static analysis on the current project to find all of its imports.
	 // params.RootPackageTree, _ = pkgtree.ListPackages(root, importroot)

	// Set up a SourceManager. This manages interaction with sources (repositories).
	tempdir, _ := ioutil.TempDir("", "gps-repocache")
	sourcemgr, _ := gps.NewSourceManager(filepath.Join(tempdir))
	defer sourcemgr.Release()

	// Prep and run the solver
	fmt.Println("Got it never")
	solver, err := gps.Prepare(params, sourcemgr)
	fmt.Println("Got it", err)
	solution, err := solver.Solve()
	fmt.Println("Hello Error", err)
	if err == nil {
		// If no failure, blow away the vendor dir and write a new one out,
		// stripping nested vendor directories as we go.
		os.RemoveAll(filepath.Join(root, "vendor"))
		gps.WriteDepTree(filepath.Join(root, "vendor"), solution, sourcemgr, true)
	}
}

type NaiveAnalyzer struct{}

// DeriveManifestAndLock is called when the solver needs manifest/lock data
// for a particular dependency project (identified by the gps.ProjectRoot
// parameter) at a particular version. That version will be checked out in a
// directory rooted at path.
func (a NaiveAnalyzer) DeriveManifestAndLock(path string, n gps.ProjectRoot) (gps.Manifest, gps.Lock, error) {
	return nil, nil, nil
}

// Reports the name and version of the analyzer. This is used internally as part
// of gps' hashing memoization scheme.
func (a NaiveAnalyzer) Info() (name string, version int) {
	return "example-analyzer", 1
}

type ManifestYaml struct{}

func (a ManifestYaml) IgnoredPackages() map[string]bool {
	return nil
}

func (a ManifestYaml) RequiredPackages() map[string]bool {
	return nil
}

func (a ManifestYaml) Overrides() gps.ProjectConstraints {
	return gps.ProjectConstraints{
		"Hello": gps.ProjectProperties{
			Source:     "https://github.com/a8uhnf/go_stack",
			Constraint: gps.NewBranch("master"),
		},
	}
}

func (a ManifestYaml) DependencyConstraints() gps.ProjectConstraints {
	return gps.ProjectConstraints{
		"Hello": gps.ProjectProperties{
			Source:     "https://github.com/a8uhnf/go_stack",
			Constraint: gps.NewBranch("master"),
		},
	}
}

func (a ManifestYaml) TestDependencyConstraints() gps.ProjectConstraints {
	return nil
}
