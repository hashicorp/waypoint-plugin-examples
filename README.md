# waypoint-plugin-examples
Example external plugins for Waypoint

# [golang](./golang)

Simple build plugin for Golang which runs go build for the source

```
plugin "golang" {
  type {
    build = true
  }
}


app "example" {

  build {
    use "golang" {
      output_name="server"
      source="./"
    }
  }
}
```