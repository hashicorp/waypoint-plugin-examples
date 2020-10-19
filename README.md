# waypoint-plugin-examples
Example external plugins for Waypoint

For a full guide on building plugins with Waypoint please see the
[Extending Waypoint](https://www.waypointproject.io/docs/extending-waypoint) documentation

## [Plugin Template](./template)

This template can be used to scaffold new Waypoint plugins. All the main Waypoint components and interfaces are
implemented in this sample.

## [Go Builder Plugin](./plugins/gobuilder_final)

This plugin is the final example code from the plugin development guide in the Waypoint documentation.

## [Filepath Plugin](./plugins/filepath)

Plugin which implements Registry, Platform, and ReleaseManager which interacts with the local file system

## Running the Sample App

1. Build the sample plugins

This command will build the plugins and move them to the ./plugins/example_app/.waypoint/plugins folder

```shell
cd plugins
make

### Build Go Builder Plugin
make -C gobuilder_final
make[1]: Entering directory '/home/nicj/go/src/github.com/hashicorp/waypoint-plugin-examples/plugins/gobuilder_final'

Build Protos
protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./builder/output.proto

Compile Plugin
go build -o ./bin/waypoint-plugin-gobuilder ./main.go 
make[1]: Leaving directory '/home/nicj/go/src/github.com/hashicorp/waypoint-plugin-examples/plugins/gobuilder_final'


### Build Filepath Builder Plugin
make -C filepath
make[1]: Entering directory '/home/nicj/go/src/github.com/hashicorp/waypoint-plugin-examples/plugins/filepath'
protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./registry/output.proto
protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./platform/output.proto
protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./release/output.proto
make[1]: Leaving directory '/home/nicj/go/src/github.com/hashicorp/waypoint-plugin-examples/plugins/filepath'


### Install Plugins
mkdir -p ./example_app/.waypoint/plugins
cp ./gobuilder_final/bin/* ./example_app/.waypoint/plugins
cp ./filepath/bin/* ./example_app/.waypoint/plugins
```

The ./plugins/example_app folder has a simple Golang application and the following Waypoint file


```javascript
project = "guides"

app "example" {

  build {
    use "gobuilder" {
      output_name = "server"
      source = "./"
    }

    registry {
      use "filepath" {
        directory = "./bin"
      }
    }
  }

  deploy {
    use "filepath" {
      directory = "./deployments"
    }
  }

  release {
    use "filepath" {}
  }
}
```

You can build the application using the following command:

```shell
cd ./example_app

waypoint init
✓ Configuration file appears valid
✓ Local mode initialized successfully
✓ Project "guides" and all apps are registered with the server.
✓ Plugins loaded and configured successfully
✓ Authentication requirements appear satisfied.

Project initialized!

You may now call 'waypoint up' to deploy your project or
commands such as 'waypoint build' to perform steps individually.
```

And then:

```shell
waypoint init
waypoint up

» Building...
✓ Application built successfully server
✓ Application binary pushed to registry

» Deploying...

» Releasing...
✓ Created release deployments/release

The deploy was successful! A Waypoint deployment URL is shown below. This
can be used internally to check your deployment and is not meant for external
traffic. You can manage this hostname using "waypoint hostname."

   Release URL: deployments/release
```