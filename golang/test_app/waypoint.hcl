project = "guides"

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