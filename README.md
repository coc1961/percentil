# percentil

## Cálculo de Percentil

### Instalar

```sh

go get -u github.com/coc1961/percentil

```

### Ejemplo de uso

> Ejemplo leyendo de stdin, percentil 20
```sh

 cat testdata/testDatosAgrupados1.csv | percentil -f - -p 20

 Output:

 308.44
 
```

> Ejemplo leyendo desde archivo, percentil 20
```sh

 percentil -f testdata/testDatosAgrupados1.csv -p 20

 Output:

 308.44
 
```

> Ejemplos de tablas de datos (agrupado y no agrupados) ver carpeta testdata

```sh

ll testdata/

total 20
drwxrwxr-x 2 user user 4096 Apr 12 20:47 ./
drwxrwxr-x 5 user user 4096 Apr 11 21:34 ../
-rw-rw-r-- 1 user user   43 Apr 12 20:48 testDatosAgrupados1.csv
-rw-rw-r-- 1 user user   66 Apr 12 20:26 testDatosAgrupados.csv
-rw-rw-r-- 1 user user   42 Apr 12 20:42 testDatosNoAgrupados.csv


cat testdata/testDatosAgrupados1.csv
200,85
300,90
400,120
500,70
600,62
700,36

cat testdata/testDatosNoAgrupados.csv 
0
4
1
0
0
7
2
1
4
0
3
9
2
0
0
4
8
1
0
9
4

```