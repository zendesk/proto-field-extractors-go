module github.com/zendesk/field-extractors-go

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	github.com/zendesk/field-extractors-go/protobuf v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.27.1
)

replace github.com/zendesk/field-extractors-go/protobuf => ./generated/github.com/zendesk/field-extractors-go/protobuf
