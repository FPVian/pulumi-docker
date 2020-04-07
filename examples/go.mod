module github.com/pulumi/pulumi-docker/examples

go 1.13

require (
	github.com/pulumi/pulumi-docker/sdk v0.0.0-00010101000000-000000000000
	github.com/pulumi/pulumi/pkg v1.13.1
	github.com/pulumi/pulumi/sdk v1.14.0
	github.com/stretchr/testify v1.5.1
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/pulumi/pulumi-docker/sdk => ../sdk
)
