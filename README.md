# pptgrep

This tool can search `ppt` files which have an input word.

## Usage

```
NAME:
   pptgrep - Searching ppt files which have an input word

USAGE:
   pptgrep [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --word value, -w value  Searching files which has the word
   --path value, -p value  Starting point (default: ".")
   --help, -h              show help
   --version, -v           print the version
```

## Example

Input only a word

```
$ ./pptgrep -w GOOGLE
/home/k2la/test.pptx
```

Input a word and a starting point

```
$ ./pptgrep -w APPLE -p ~
/home/k2la/test1.pptx
/home/k2la/test2.pptx
/home/k2la/test3.pptx
```
