# service-profile-api

A repository for the  profile service api being developed
for ant investors

### How do I update the definitions? ###

* The api definition is defined in the proto file profile.proto
* To update the proto service you need to run the command :


`protoc \
    --proto_path=../common/service \
    --go_out=./ --validate_out=lang=go:. \
    common/validate.proto common/common.proto profile/v1/profile.proto`
`protoc --proto_path=../common/service common/common.proto profile/v1/profile.proto --go-grpc_out=./ `

  with that in place update the implementation appropriately
