module github.com/hashicorp/waypoint-plugin-examples

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/hashicorp/waypoint v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.25.0
)

replace github.com/hashicorp/waypoint => ../waypoint

replace github.com/hashicorp/horizon => ../horizon

replace github.com/hashicorp/waypoint-hzn => ../waypoint-hzn
