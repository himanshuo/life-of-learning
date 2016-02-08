Testing Results
====================
### CORRECT SAMPLE RESULT ON SLURM:
##### gcc -o original original.c -lm; ./original large_input.txt 0.5
Value of system clock at start = 0
Coordinates will be read from file: large_input.txt
Natom = 20000
cut =     0.5000
Value of system clock after coord read = 30000
Value of system clock after coord read and E calc = 11430000
                         Final Results
                         -------------
                   Num Pairs = 55127556
                     Total E = 231391717.1474710703
     Time to read coord file =         0.0300 Seconds
         Time to calculate E =        11.4000 Seconds
        Total Execution Time =        11.4300 Seconds




### Optimization Table
| Optimization Number  | Optimization Description |
|---|---|
| 0  | -O2  |
| 1  | -O3  |
| 2  | remove the if statement by making it a part of the for guard condition
| 3  | use t += c instead of t=t+c |
| 4  | move lookups out of for loops (coords) |
| 5  | move lookups out of for loops (q) |
| 6  | reduce e^(rij*qi)*e^(rij*qj) to e^(rij*qi+rij*qj))   
| 7  | division to multiplication for 1.0/a


### Test Results
| Test Name  | Natoms | Optimizations  | Time to read coord file (s)  | Time to calculate E (s)  | Total Execution Time (s)  |
|---|---|---|---|---|---|
| original   | 20000  | -  | 0.0300  | 17.6400  | 17.6700  |
| O2  | 20000  |  0 | 0.0100  | 6.2300  | 6.2400  |
| O3  | 20000  |  1 | 0.0100  | 6.2100  | 6.2200  |
| no_if  | 20000  | 1,2  | 0.0100  | 6.0300  | 6.0400  |
| plus_eq  | 20000  | 1,2,3  | 0.0300  | 10.3000  | 10.3300  |
| less_lookups(coords)  | 20000  | 1,2,4  | 0.0100  | 6.0000  | 6.0100  |
| less_lookups(q)  | 20000  | 1,2,4,5  | 0.0100  | 5.9500  | 5.9600  |
| simplify  | 20000  | 1,2,4,5,6  | 0.0100  | 4.3300  | 4.3400  |


### Potential Speedup techniques (todo)
2. convert rij = sqrt(vec2);; if(rij<cut) to if(vec2*vec2 < cut*cut)
3. division to multiplication (for /rij)
5. see if you can remove exp
9. remove unneccessary iterations: for(j: 1->natom) if(j<i) -> for(j=1; j<natom && j<i; ++j)
