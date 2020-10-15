module github.com/hashicorp/waypoint-plugin-examples/plugins/filepath

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/hashicorp/go-hclog v0.14.1
	github.com/hashicorp/waypoint-plugin-examples/plugins/gobuilder_final v0.0.0-00010101000000-000000000000
	github.com/hashicorp/waypoint-plugin-sdk v0.0.0-20201015055750-7833de86f9fa
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/hashicorp/waypoint-plugin-examples/plugins/gobuilder_final => ../gobuilder_final
