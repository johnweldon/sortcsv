// sortcsv allows sorting csv, and tsv, files by column name
//
//    NAME:
//       sortcsv sort -
//
//    USAGE:
//       sortcsv sort [command options] [arguments...]
//
//    OPTIONS:
//       --infile value, --in value, -i value    Defaults to stdin (-) if not specified. (default: -)
//       --outfile value, --out value, -o value  Defaults to stdout (-) if not specified. (default: -)
//       --sortby value, -s value                Columns to sort by, repeat for sub-sort
//       --tab, -t                               Input is tab delimited, rather than comma delimited (default: false)
//       --help, -h                              show help (default: false)
package main
