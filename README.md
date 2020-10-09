# waypoint-plugin-examples
Example external plugins for Waypoint

## [Plugin Template](./template)

You can use this default template as the basis for your new plugin, all the main Waypoint interfaces are implemented in this sample

## [golang](./golang)

Simple build plugin for Golang which runs go build for the source

## [filepath](./filepath)

Plugin which implements Registry, Platform, and ReleaseManager which interacts with the local file system

## Running the Sample App

1. Build the sample plugins

This command will build the plugins and move them to the ./example_app/.waypoint/plugins folder

```shell
make

make -C golang
make[1]: Entering directory '/home/nicj/go/src/github.com/hashicorp/waypoint-plugin-examples/golang'
go generate
CGO_ENABLED=0 go build -o ./bin/waypoint-plugin-golang ./*.go
make[1]: Leaving directory '/home/nicj/go/src/github.com/hashicorp/waypoint-plugin-examples/golang'
make -C filepath
make[1]: Entering directory '/home/nicj/go/src/github.com/hashicorp/waypoint-plugin-examples/filepath'
go generate
CGO_ENABLED=0 go build -o ./bin/waypoint-plugin-filepath ./*.go
make[1]: Leaving directory '/home/nicj/go/src/github.com/hashicorp/waypoint-plugin-examples/filepath'
mkdir -p ./example_app/.waypoint/plugins
cp ./golang/bin/* ./example_app/.waypoint/plugins
cp ./filepath/bin/* ./example_app/.waypoint/plugins
```

The ./example_app folder has a simeple Golang application and the following Waypoint file


```javascript
project = "guides"

app "example" {

  build {
    use "golang" {
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