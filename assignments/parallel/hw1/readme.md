Testing Results
====================
### CORRECT SAMPLE RESULT:
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

### Optimization Table
| Optimization Number  | Optimization Description |
|---|---|
| 0  | -O2  |
| 1  | -O3  |
|   |   |
|   |   |
|   |   |

### Test Results
| Test Name  | Natoms | Optimizations  | Time to read coord file (s)  | Time to calculate E (s)  | Total Execution Time (s)  |
|---|---|---|---|---|---|
| original   | 100  | -  | 0.0002  | 0.0004  | 0.0006  |
| O2  | 100  |  0 | 0.0002  | 0.0002  | 0.0004  |
| O3  | 100  |  1 | 0.0002  | 0.0002  | 0.0004  |
|   |   |   |   |   |   |
|   |   |   |   |   |   |




### Potential Speedup techniques (todo)
1. convert pow(stuff, 2) to stuff * stuff
2. convert rij = sqrt(vec2);; if(rij<cut) to if(vec2*vec2 < cut*cut)
3. division to multiplication
4. move expressions out of for loops
5. see if you can remove exp
6. remove coords[0][i-1] out of second for loop
7. remove q[i-1] out of loop
8. a=3.2 then 1.0/a to a=1.0/3.2
9.
