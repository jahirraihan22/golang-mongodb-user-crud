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

vagrant init "ubuntu/jammy64" --output ./test/$Vagrantfile

if [ $? -ne 0 ]; then
  echo -e "Something went wrong, exiting............."
  exit 1
fi

vagrant up ./test/$Vagrantfile