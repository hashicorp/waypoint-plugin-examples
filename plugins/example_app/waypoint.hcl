project = "custom-waypoint-plugins"

app "example" {

  build {
    use "gobuilder" {
      output_name = "app"
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
