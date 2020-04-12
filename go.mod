module github.com/zephinzer/eg-cadence-jobportal

go 1.14

require (
	github.com/uber-go/tally v3.3.15+incompatible
	go.uber.org/cadence v0.11.2
	go.uber.org/yarpc v1.44.0
	go.uber.org/zap v1.14.1
)

replace github.com/apache/thrift => github.com/apache/thrift v0.0.0-20190309152529-a9b748bb0e02
