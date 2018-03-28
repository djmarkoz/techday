# Golang Techday

[Presentation slides]

----

This git repository is a sample CLI application build with [Cobra] written in Go. 
It also contains the [examples](docs/presentation/src) from the presentation.
 
## Installing Go

```
brew install go
brew install dep

# Add these to your .profile
export PATH=$PATH:$(go env GOPATH)/bin
export GOPATH=$(go env GOPATH)
```

## Install Cobra
``` 
go get -u github.com/spf13/cobra/cobra
```

## Getting and building this project
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

## Try adding new commands to the sample CLI. 😎

```
cd cmd/techday-cli
cobra add newCommand
```

## Run the CLI

If you want to build right away there are two options:

##### From within Intellij

1. Open `cmd/techday-cli/main.go`
2. Press the play arrow before the `func main()` method 
3. Press `Run`

##### From the shell

Either one of:

```
$ ./bin/techday-cli
$ make run
$ make install && techday-cli #globally installed
```

##### In Docker (does not require Go, Dep or Cobra installed!)

```
$ make docker
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

## Makefile

```
Options:

make
    builds and tests the project
    
make build
    builds the project
    
make clean
    removes ./bin
        
make cleanall
    removes ./bin and ./vender    

make docker
    builds a Docker image with the CLI (does not require Go, Dep or Cobra installed!)
    
make docker-run
    runs the CLI Docker image
        
make install
    installs binaries globally on $GOPATH/bin
    
make test
    tests the project
    
make uninstall
    removes binaries from $GOPATH/bin
    
make vendor
    downloads the vendor dependencies
```
----

[Cobra]: https://github.com/spf13/cobra
[Docker environment]: https://docs.docker.com/engine
[Go environment]: https://golang.org/doc/install
[GoLand]: https://www.jetbrains.com/go/
[Presentation slides]: https://go-talks.appspot.com/github.com/djmarkoz/techday/docs/presentation/techday.slide