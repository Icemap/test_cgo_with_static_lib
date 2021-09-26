#!/bin/bash

## rebuild static lib
cd gowapper/lib
sh build_static_lib.sh

## bazel build with static library
cd ../../main
bazel build ...
