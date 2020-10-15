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