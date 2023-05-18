#!/bin/bash

echo -e "creating vm using vagrant\n"

echo -e "--------------------------------------------"

echo -e "checking dependency"

echo -e "$(vagrant -v)\n"


if [ $? -ne 0 ]; then
  echo -e "vagrant is not installed\n"
  echo -e "installing vagrant\n"

# TODO installing command
fi

Vagrantfile="Vagrantfile"

if [ -f "./test/$Vagrantfile" ]; then
  echo "a vm already exist in this directory"
  exit 1
fi

vagrant init "ubuntu/jammy64" --output ./test/$Vagrantfile

if [ $? -ne 0 ]; then
  echo -e "Something went wrong, exiting............."
  exit 1
fi
# vagrant up new-test-vm

# TODO
# ______________________________________

#1. local memory use
#2. handling local file system
#3. ssh