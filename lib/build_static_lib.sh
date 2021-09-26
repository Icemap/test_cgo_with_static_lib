#!/bin/bash

gcc -c test_calc.cc
ar rc libtest_calc.a test_calc.o