#!/bin/bash

for kind in map slice str; do 
    cat out/${kind}-bench.txt | 
	awk "/BenchmarkGC/ {count++; total+=\$3} END {printf \"%-5s: %12.2f\n\", \"${kind}\", total/count}";
done 
