#!/bin/bash

## delete build file
rm -f libtest_calc.a
rm -f test_calc.o

## rebuild
gcc -c test_calc.c
ar rc libtest_calc.a test_calc.o