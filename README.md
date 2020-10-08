# waypoint-plugin-examples
Example external plugins for Waypoint

# [golang](./golang)

Simple build plugin for Golang which runs go build for the source

```
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
