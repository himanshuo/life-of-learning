/** Generates input for the original.c or original.f90 programs
 * Written by Aaron Bloomfield, 2007
 *
 * The comments that follow are from the HPC Bootcamp hw1
 * instructions.
 *
 * To generate an input file, you will need to specify one or two
 * parameters.  The first parameter is how many coordinates to
 * generate.  The second parameter is an optional random seed.
 *
 * To generate a file containing 100 coordinates, enter the command:
 *
 *	./generate_input 100 > input.txt
 *
 * This will create an input.txt file that contains 100 coordinates in
 * the input file format needed for original.c/original.f.  You can
 * view this file through Windows Explorer, although Notepad will
 * screw up the formatting, so use Wordpad or Word.  Each time you run
 * the generate_input program, it will generate a different set of 100
 * coordinates.  If you want to generate the same set of coordinates
 * each time (useful for comparing your program to somebody else's),
 * you should enter a random seed, which can be any integer value.
 * For example:
 *
 *	./generate_input 100 5 > input.txt
 *
 * This will generate the exact same set of coordinates each time the
 * program runs.  Note that running the generate_input program with
 * the same seed but on different computers (dogwood and your
 * Linux/Windows machine, for example) will not necessarily create the
 * same input file.
 *
 */

#include <stdio.h>
#include <stdlib.h>

int main (int argc, char **argv) {
    int ret, n, seed = time(NULL), i;

    /* parse parameters */
    if ( (argc < 2) || (argc > 3) ) {
        printf ("Usage: %s <size> [<seed>]\n", argv[0]);
        exit(1);
    }

    /* read in first parameter */
    ret = sscanf (argv[1], "%d", &n);
    if ( ret != 1 ) {
        printf ("Error: invalid size format: '%s'\n", argv[1]);
        exit(2);
    }

    /* read in second parameter */
    if ( argc == 3 ) {
        ret = sscanf (argv[2], "%d", &seed);
        if ( ret != 1 ) {
            printf ("Error: invalid seed format: '%s'\n", argv[2]);
            exit(3);
        }
    }

    printf ("%d\n", n);

    srand(seed);

    for ( i = 0; i < n; i++ )
        // the fortran version will not work unless these are printed out as %9.3f
        printf ("%9.3f %9.3f %9.3f %9.3f\n", rand()/(double)RAND_MAX, rand()/(double)RAND_MAX, rand()/(double)RAND_MAX, rand()/(double)RAND_MAX);

    return 0;
}
