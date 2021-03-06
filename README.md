## Installation

`go get github.com/Gira-X/simple-ls`

This generates the command `simple-ls` in `GOPATH/bin` or `GOBIN`.

## Options

It only supports the command argument `hidden` which only lists hidden directories and files.

## Fish shell functions

I rarely use the default `ls` and I am instead using those functions in my 
[fish shell](https://fishshell.com/)
in `config.fish` with `simple-ls` renamed to `l`:

```bash
function l
	simple-ls
end

function lh
	simple-ls hidden
end
```

## Output

Directories are printed in a blue foreground.

[![asciicast](https://asciinema.org/a/360030.svg)](https://asciinema.org/a/360030)

```bash
> simple-ls
/Users/Gira/simple-ls (2 hidden dirs)

LICENSE  README.md  extract.go  go.mod  go.sum  main.go
```

Note that `.DS_Store` is not considered in the `... hidden files` count to avoid printing
`(1 hidden file)` where `.DS_Store` is the only hidden file, and generally the file system
is littered with those files anyway.

```bash
> simple-ls hidden
/Users/Gira/simple-ls (2 hidden dirs)

.git  .idea

.DS_Store
```
