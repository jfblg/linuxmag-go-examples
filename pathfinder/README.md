# Pathfinder

Linux Magazin 09/19

cdbm - is a utility which provides a user (after pressing "c") a list of his/her last visited directories.

Generate a new go module and start a build process there.
```
go mod init cdbm
go build
```

In order for utility to work it requires an "integration" with a shell

1. by each prompt a utility is called in order to add the current directory to the utility's DB
```
export PS1='$(cdbm -add)\h.\u:\W$ '
```

2. user needs to define a shell function "c" so that cdbm can present him a list of directories
```
function c() { dir=$(cdbm 3>&1 1>&2 2>&3); cd $dir; }
```
