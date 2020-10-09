project = "guides"

app "example" {

  build {
    use "gobuilder" {
      output_name = "app"
      source = "./"
    }
  }

  deploy {
    use "gobuilder" {}
  }
}