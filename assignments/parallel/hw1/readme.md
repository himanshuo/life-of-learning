Testing Results
====================
### CORRECT SAMPLE RESULT:
##### gcc -o original original.c -lm; ./original input.txt 0.5
Value of system clock at start = 539

Coordinates will be read from file: input.txt
Natom = 100
cut =     0.5000
Value of system clock after coord read = 757
Value of system clock after coord read and E calc = 1144
                         Final Results
                         -------------
                   Num Pairs = 1468
                     Total E = 6066.9392663564
     Time to read coord file =         0.0002 Seconds
         Time to calculate E =         0.0004 Seconds
        Total Execution Time =         0.0006 Seconds
##### gcc -o original original.c -lm; ./original large_input.txt 0.5
Value of system clock at start = 521
Coordinates will be read from file: large_input.txt
Natom = 20000
cut =     0.5000
Value of system clock after coord read = 19810
Value of system clock after coord read and E calc = 4392493
                         Final Results
                         -------------
                   Num Pairs = 55127556
                     Total E = 231391717.1474710703
     Time to read coord file =         0.0193 Seconds
         Time to calculate E =         4.3727 Seconds
        Total Execution Time =         4.3920 Seconds

### Optimization Table
| Optimization Number  | Optimization Description |
|---|---|
| 0  | -O2  |
| 1  | -O3  |
| 2  | convert exponentials (pow) into simple multiplications  |
| 3  | reduce e^(rij * qi) * e^(rij * qj) to e^(rij*(qi + qj))  |
| 4  | division to multiplication for 1.0/a  |
| 5  | move lookups out of for loops (coords) |
| 6  | move lookups out of for loops (q) |
| 7  |   |
| 8  |   |

### Test Results
| Test Name  | Natoms | Optimizations  | Time to read coord file (s)  | Time to calculate E (s)  | Total Execution Time (s)  |
|---|---|---|---|---|---|
| original   | 100  | -  | 0.0002  | 0.0004  | 0.0006  |
| O2  | 100  |  0 | 0.0002  | 0.0002  | 0.0004  |
| O3  | 100  |  1 | 0.0002  | 0.0002  | 0.0004  |
| pow  | 100  | 2  | 0.0002  | 0.0004  | 0.0006  |
| simplify equations  | 100  | 2,3  | 0.0002  | 0.0002  | 0.0005  |
| boo division  | 100  | 2,3,4  | 0.0002  | 0.0002  | 0.0005 |
| out of loop coords  | 100  | 2,3,4,5  | 0.0002  | 0.0002  | 0.0004 |
| out of loop q | 100  | 2,3,4,5,6  | 0.0002  | 0.0002  | 0.0004 |
| O3 again | 100  | 1,2,3,4,5,6  | 0.0002  | 0.0001  | 0.0004 |
| larger sample | 20000  | 1,2,3,4,5,6  | 0.0188  | 4.0596  | 4.0784 |
| original large | 20000  | -  | 0.0188  | 4.0596  | 4.0784 |


### Potential Speedup techniques (todo)
2. convert rij = sqrt(vec2);; if(rij<cut) to if(vec2*vec2 < cut*cut)
3. division to multiplication (for /rij)
5. see if you can remove exp
9. remove unneccessary iterations: for(j: 1->natom) if(j<i) -> for(j=1; j<natom && j<i; ++j)
