# 01 Environment Preparation

##Install Go
Download [Go binary distributions](https://golang.org/dl/), currently using 1.10.3 version.

### Set up *GOPATH* environment variable
The *GOPATH* environment variable specifies the location of your [workspace](https://golang.org/doc/code.html#Workspaces). It defaults to a directory named go inside your home directory, so $HOME/go on Unix. Add the workspace's bin subdirectory to your PATH:

export PATH=$PATH:$(go env GOPATH)/bin
If you have not set GOPATH, you can substitute $HOME/go in those commands or else run:

export GOPATH=$(go env GOPATH) 
Set up directory structure
Inside the $GOPATH directory, you have to create 3 folders (if they are not there already):
```
$ cd $GOPATH
$ mkdir src
$ mkdir pkg
$ mkdir bin
```
The purpose of each directory can be seen from its name:

/bin - Is where all the executable binaries created by compiling your code.
/pkg - Contains package objects made by libraries (which you don’t have to worry about now).
/src - Is where all your Go source code goes.