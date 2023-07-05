module github.com/hashicorp/waypoint-plugin-examples/plugins/filepath

go 1.14

require (
	github.com/hashicorp/go-hclog v0.14.1
	github.com/hashicorp/waypoint-plugin-examples/plugins/gobuilder_final v0.0.0-00010101000000-000000000000
	github.com/hashicorp/waypoint-plugin-sdk v0.0.0-20211014201256-80d5426fa6e4
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

replace github.com/hashicorp/waypoint-plugin-examples/plugins/gobuilder_final => ../gobuilder_final
