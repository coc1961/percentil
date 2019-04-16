#!/bin/sh

go get -u github.com/coc1961/go/trs

printPersentil(){
    data=$(/bin/echo $1 | trs ";" "\\n" )
    for a in $(seq 99)
    do 
        res=`echo $data | go run percentil.go -f - -p $a`; echo "perc p$a=$res"
    done
}


#printPersentil "3;6;7;8;8;10;13;15;16;20"

if [  -n "$1"  ]; 
then 
    printPersentil $1
else 
    echo "Ejecutar: ---> $0 \"numeros....\""
    echo "          |"
    echo "          ---> Ejemplo $0 \"3;4;5;6;7;8;9\""
fi


# Ejemplo con datos agrupados
# ./test.sh "200,85;300,90;400,120;500,70;600,62;700,36"

# Ejemplo con datos no agrupados
# ./test.sh "3;6;7;8;8;10;13;15;16;20"
