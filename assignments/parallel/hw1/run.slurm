#!/bin/bash
#SBATCH -N 1
#SBATCH -n 1
#SBATCH -t 01:00:00
#SBATCH -o output_filename
#SBATCH -p training
#SBATCH -A parallelcomputing


gcc -O3 -lm -o speedup speedup.c ; ./speedup large_input.txt 0.5
