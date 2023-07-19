package main

// The go.mod file is the root of dependency management in GoLang. All the modules which are needed
// or to be used in the project are maintained in go.mod file. For all the packages we are going
// to import/use in our project, it will create an entry of those modules in go.mod. Having a go mod file
// saves the efforts of running the go get command for each dependent module to run the project successfully.

// go.sum file is a file which maintains the checksum so when you run the project again
// it will not install all packages again. But use the cache which is stored inside
// $GOPATH/pkg/mod directory (module cache directory). go.sum is a generated file you don’t
// have to edit or modify this file.

// go mod tidy adds any missing module requirements necessary to build the current module’s packages
// and dependencies, if there are some not used dependencies go mod tidy will remove those from
// go.mod accordingly. It also adds any missing entries to go.sum and removes unnecessary entries.

// go mod vendor generates a vendor directory with the versions available. It copies all third-party
// dependencies to a vendor folder in your project root.
