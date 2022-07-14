# sortcsv

Tool to sort csv files by named columns.

## NB 

Deprecated

## Example

The [example](example_sort) script shows how to use it to sort with multiple sub-sort columns.

```shell
$ ./example_sort

Sort by Last, First, Email
--
First,Last,Email
Mary,Bunce,mb@example.com
John,Doe,doej@example.com
Mary,O'Henry,mary@example.com
Mary,O'Henry,moh@example.com
Adam,Van der Oos,voos@example.com


Sort by First, Last, Email
--
First,Last,Email
Adam,Van der Oos,voos@example.com
John,Doe,doej@example.com
Mary,Bunce,mb@example.com
Mary,O'Henry,mary@example.com
Mary,O'Henry,moh@example.com
```

## Installation

Requires Go 1.8+ because it uses the new [`sort.Slice`](https://golang.org/pkg/sort/#Slice) feature.

`go get jw4.us/sortcsv`

## Usage

`sortcsv sort`

```
NAME:
   sortcsv sort - 

USAGE:
   sortcsv sort [command options] [arguments...]

OPTIONS:
   --infile value, --in value, -i value    Defaults to stdin (-) if not specified. (default: -)
   --outfile value, --out value, -o value  Defaults to stdout (-) if not specified. (default: -)
   --sortby value, -s value                Columns to sort by, repeat for sub-sort
   --tab, -t                               Input is tab delimited, rather than comma delimited (default: false)
   --help, -h                              show help (default: false)
``` 

If `--infile` or `--outfile` are not specified they default to `stdin` and `stdout` respectively.

The `--sortby` flag can be specified repeatedly to add sub-sorting.
   
