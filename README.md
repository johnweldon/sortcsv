# sortcsv

Tool to sort csv files by named columns.

## Example

The [example](example_sort_script.sh) script shows how to use it to sort the LinkedIn Basic data archive Contacts.csv and Conections.csv

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
   
