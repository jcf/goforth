#!/usr/bin/env zsh
root=`pwd`

for d in *(/); do
  echo "Building $d..."
  cd $d
  go build
  cd $root
done
