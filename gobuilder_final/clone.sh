#! /bin/bash

function help() {
  echo "usage: clone.sh [location] [package]"
  echo ""
  echo "e.g."
  echo "./clone.sh ../testing github.com/hashicorp/waypoint-plugin-examples/testing"
}

function process() {
  default="github.com/hashicorp/waypoint-plugin-examples/gobuilder"
  new="$2"

  # replace all the instance of the template
  find $1 -type f -exec sed -i.bak "s|${default}|${new}|g" {} \;

  # Remove backup files
  find $1 -name *.bak -exec rm -rf {} \;

  # Remove go.sum
  rm -f $1/go.sum
}

for arg in "$@"
do
    if [ "$arg" == "--help" ] || [ "$arg" == "-h" ]; then
      help
    fi
done

if [ "$1" == "" ]; then
  echo "Error: Please specify the location to clone this template to"
  echo ""
  help
  exit 1
fi

if [ "$2" == "" ]; then
  echo "Error: Please specify the package for the new plugin"
  echo ""
  help
  exit 1
fi

cp -R . $1
process $1 $2

echo "Created new pluging in $1"
echo "You can build this plugin by running the following command"
echo ""
echo "cd $1 && make"