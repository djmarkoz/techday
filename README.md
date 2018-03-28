# Golang Techday

----

This git repository is a sample CLI application build with [Cobra] written in Go. 

It contains the examples from the presentation. 

Try adding new commands to the CLI.

```
cobra add doSomething
```
 

## Installing Go

```
brew install go
brew install dep

export PATH=$PATH:$(go env GOPATH)/bin
export GOPATH=$(go env GOPATH)
```

## Install Cobra
``` 
go get -u github.com/spf13/cobra/cobra
```

## Getting the Project
```
go get -u -v -d github.com/djmarkoz/techday 
cd $GOPATH/src/github.com/djmarkoz/techday
make
```

## Setting up Intellij for Go
1. **Install plugin: `Go`** _This plugin extends IntelliJ platform with Go-specific coding assistance and tool integrations, and has everything you could find in [GoLand]._
2. **Install plugin: `File Watchers`** _This plugin will allow you to format your Go code and organize/fix your imports. **Must have**_
2. **Restart Intellij**
3. **`File` > `New` > `Project...`**
4. **Choose `Go` on the left**
5. **Select your Go version and press `Next`**
6. **Set the `Project location`;** probably `~/go/src/github.com/djmarkoz/techday` 
7. **Press `Finish`**
8. **Goto preferences;** `cmd+,`
9. **`Tools` > `File Watchers` > press `+` > select `goimports`** > press `Ok`
 
----

## Just Build and Run the project

If you want to build right away there are two options:

##### You have a working [Go environment].

```
$ go get -d github.com/djmarkoz/techday
$ cd $GOPATH/src/github.com/djmarkoz/techday
$ make
$ ./bin/techday-cli
```

##### You have a working [Docker environment].

```
$ git clone https://github.com/djmarkoz/techday
$ cd techday
$ make docker-run
```

----

## Package design

This `techday` project currently has one application called `techday-cli`.

#### cmd/

All the different applications this project has. Like `techday-cli`, or perhaps in the future `techday-server`. These packages should only contain packages which are specific to one application. Packages in `cmd/` can have dependencies to `internal/`, `pgk/` and remote packages.

#### internal/

Everything only for this project. Likely to be shared by multiple applications (cmds). The compiler will project these packages, they can not be reused by other projects. Packages in `internal/` should not have dependencies to `cmd/` packages. But probably will have dependencies to packages in `pkg/` or even remote packages. 

#### pkg/ 

Some packages are likely to be reused by other projects and are like libraries. These should be placed in `pkg/`.  Packages in `pkg/` should not have any dependencies on `internal/` or `cmd/` packages. It's best to not even have dependencies on any remote package.    

----

[Cobra]: https://github.com/spf13/cobra
[Docker environment]: https://docs.docker.com/engine
[Go environment]: https://golang.org/doc/install
[GoLand]: https://www.jetbrains.com/go/
