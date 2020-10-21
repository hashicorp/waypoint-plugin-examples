#! /bin/bash

function help() {
  echo "usage: clone.sh [name] [location] [package]"
  echo ""
  echo "This script requires three parameters:"

  echo ""
  echo "[name]"
  echo "The name of your plugin e.g. myplugin"

  echo ""
  echo "[location]"
  echo "The location to store the cloned template"
 
  echo ""
  echo "[package]"
  echo "The Go package name for the cloned template"
  echo "github.com/hashicorp/waypoint-plugin-examples/myplugin"
  
  echo ""
  echo "e.g."
  echo "./clone.sh myplugin ../myplugin github.com/hashicorp/waypoint-plugin-examples/myplugin"
}

function process() {
  # Rename the github pacakges
  default="github.com/hashicorp/waypoint-plugin-examples/template"
  new="$3"

  # replace all the instance of the template
  find $2 -type f -exec sed -i.bak "s|${default}|${new}|g" {} \;

  # Remove backup files
  find $2 -name *.bak -exec rm -rf {} \;

  # Process the github actions workflow and makefile 
  default="template"
  new="$1"
  
  # replace all the instance of the template
  sed -i.bak "s|${default}|${new}|g" $2/Makefile
  sed -i.bak "s|${default}|${new}|g" $2/.github/workflows/build-plugin.yml

  # Remove backup files
  find $2 -name *.bak -exec rm -rf {} \;

  # Remove go.sum
  rm -f $1/go.sum
}

for arg in "$@"
do
    if [ "$arg" == "--help" ] || [ "$arg" == "-h" ]; then
      help
      exit 0
    fi
done

if [ "$1" == "" ]; then
  echo "Error: Please specify the name of your plugin"
  echo ""
  help
  exit 1
fi

if [ "$2" == "" ]; then
  echo "Error: Please specify the location to clone this template to"
  echo ""
  help
  exit 1
fi

if [ "$3" == "" ]; then
  echo "Error: Please specify the package for the new plugin"
  echo ""
  help
  exit 1
fi

cp -R . $2
process $1 $2 $3

echo "Created new plugin in $1"
echo "You can build this plugin by running the following command"
echo ""
echo "cd $2 && make"
