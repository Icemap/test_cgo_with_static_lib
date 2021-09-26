#!/bin/bash

rm -f libtest_calc.a test_calc.o
gcc -c test_calc.c
ar rc libtest_calc.a test_calc.o