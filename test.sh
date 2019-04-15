#!/bin/sh

go get -u github.com/coc1961/go/trs

data=$(/bin/echo "3,6,7,8,8,10,13,15,16,20" | trs "," "\\n")
for a in $(seq 99); do res=`echo $data | go run percentil.go -f - -p $a`; echo "perc p$a=$res"; done