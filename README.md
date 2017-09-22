# Don't Repeat Yourself
Prefix a command to avoid repetition.

## Installation
```
go install -u github.com/inkel/prefix
```

Or download the latest [release](https://github.com/inkel/prefix/releases).

## Usage

```
$ prefix git
git > init
Initialized empty Git repository in /home/inkel/dev/go/src/github.com/inkel/prefix/.git/
git > add main.go main_test.go
git > s
## Initial commit on master
A  main.go
A  main_test.go
?? prefix
git > commit -m "initial commit"
[master (root-commit) 6727b37] initial commit
 2 files changed, 138 insertions(+)
 create mode 100644 main.go
 create mode 100644 main_test.go
git > remote add origin git@github.com:inkel/prefix.git
git > push -u origin master
Counting objects: 4, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (4/4), done.
Writing objects: 100% (4/4), 1.99 KiB | 0 bytes/s, done.
Total 4 (delta 0), reused 0 (delta 0)
To git@github.com:inkel/prefix.git
 * [new branch]      master -> master
Branch master set up to track remote branch master from origin.
git > ^D
```

## License
See [LICENSE](LICENSE).
