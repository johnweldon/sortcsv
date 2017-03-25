#!/bin/bash

#
# If sortcsv is not in your path, use Go (1.8+) to download and install it
#
command -v sortcsv >/dev/null 2>&1
if [ "$?" -ne "0" ]; then
  go get -v -u github.com/johnweldon/sortcsv
  go build -i github.com/johnweldon/sortcsv
fi

#
# Bail immediately on error
#
set -e

#
# These examples sort the two contacts csv files from the Basic LinkedIn export archive
#

CONTACTS="Contacts"
sortcsv sort -i ${CONTACTS}.csv -s "Last Name" -s "First Name" -s "Email" -s "Phone Numbers" -s "Company" -s "Tags" -o _${CONTACTS}.csv
mv _${CONTACTS}.csv ${CONTACTS}.csv

CONNECTIONS="Connections"
sortcsv sort -i ${CONNECTIONS}.csv -s "LastName" -s "FirstName" -s "EmailAddress" -s "Company" -s "Tags" -o _${CONNECTIONS}.csv
mv _${CONNECTIONS}.csv ${CONNECTIONS}.csv
