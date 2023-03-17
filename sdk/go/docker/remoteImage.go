// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// <!-- Bug: Type and Name are switched -->
// Pulls a Docker image to a given Docker host from a Docker Registry.
//
//	This resource will *not* pull new layers of the image automatically unless used in conjunction with RegistryImage data source to update the `pullTriggers` field.
//
// ## Example Usage
// ### Basic
//
// Finds and downloads the latest `ubuntu:precise` image but does not check
// for further updates of the image
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := docker.NewRemoteImage(ctx, "ubuntu", &docker.RemoteImageArgs{
//				Name: pulumi.String("ubuntu:precise"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Dynamic updates
//
// To be able to update an image dynamically when the `sha256` sum changes,
// you need to use it in combination with `RegistryImage` as follows:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ubuntuRegistryImage, err := docker.LookupRegistryImage(ctx, &docker.LookupRegistryImageArgs{
//				Name: "ubuntu:precise",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = docker.NewRemoteImage(ctx, "ubuntuRemoteImage", &docker.RemoteImageArgs{
//				Name: *pulumi.String(ubuntuRegistryImage.Name),
//				PullTriggers: pulumi.StringArray{
//					*pulumi.String(ubuntuRegistryImage.Sha256Digest),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type RemoteImage struct {
	pulumi.CustomResourceState

	// Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
	Build RemoteImageBuildPtrOutput `pulumi:"build"`
	// If true, then the image is removed forcibly when the resource is destroyed.
	ForceRemove pulumi.BoolPtrOutput `pulumi:"forceRemove"`
	// The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
	KeepLocally pulumi.BoolPtrOutput `pulumi:"keepLocally"`
	// The name of the Docker image, including any tags or SHA256 repo digests.
	Name pulumi.StringOutput `pulumi:"name"`
	// The platform to use when pulling the image. Defaults to the platform of the current machine.
	Platform pulumi.StringPtrOutput `pulumi:"platform"`
	// List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
	PullTriggers pulumi.StringArrayOutput `pulumi:"pullTriggers"`
	// The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`.
	RepoDigest pulumi.StringOutput `pulumi:"repoDigest"`
	// A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
	Triggers pulumi.MapOutput `pulumi:"triggers"`
}

// NewRemoteImage registers a new resource with the given unique name, arguments, and options.
func NewRemoteImage(ctx *pulumi.Context,
	name string, args *RemoteImageArgs, opts ...pulumi.ResourceOption) (*RemoteImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource RemoteImage
	err := ctx.RegisterResource("docker:index/remoteImage:RemoteImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemoteImage gets an existing RemoteImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemoteImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemoteImageState, opts ...pulumi.ResourceOption) (*RemoteImage, error) {
	var resource RemoteImage
	err := ctx.ReadResource("docker:index/remoteImage:RemoteImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemoteImage resources.
type remoteImageState struct {
	// Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
	Build *RemoteImageBuild `pulumi:"build"`
	// If true, then the image is removed forcibly when the resource is destroyed.
	ForceRemove *bool `pulumi:"forceRemove"`
	// The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
	ImageId *string `pulumi:"imageId"`
	// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
	KeepLocally *bool `pulumi:"keepLocally"`
	// The name of the Docker image, including any tags or SHA256 repo digests.
	Name *string `pulumi:"name"`
	// The platform to use when pulling the image. Defaults to the platform of the current machine.
	Platform *string `pulumi:"platform"`
	// List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
	PullTriggers []string `pulumi:"pullTriggers"`
	// The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`.
	RepoDigest *string `pulumi:"repoDigest"`
	// A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
	Triggers map[string]interface{} `pulumi:"triggers"`
}

type RemoteImageState struct {
	// Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
	Build RemoteImageBuildPtrInput
	// If true, then the image is removed forcibly when the resource is destroyed.
	ForceRemove pulumi.BoolPtrInput
	// The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
	ImageId pulumi.StringPtrInput
	// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
	KeepLocally pulumi.BoolPtrInput
	// The name of the Docker image, including any tags or SHA256 repo digests.
	Name pulumi.StringPtrInput
	// The platform to use when pulling the image. Defaults to the platform of the current machine.
	Platform pulumi.StringPtrInput
	// List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
	PullTriggers pulumi.StringArrayInput
	// The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`.
	RepoDigest pulumi.StringPtrInput
	// A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
	Triggers pulumi.MapInput
}

func (RemoteImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteImageState)(nil)).Elem()
}

type remoteImageArgs struct {
	// Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
	Build *RemoteImageBuild `pulumi:"build"`
	// If true, then the image is removed forcibly when the resource is destroyed.
	ForceRemove *bool `pulumi:"forceRemove"`
	// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
	KeepLocally *bool `pulumi:"keepLocally"`
	// The name of the Docker image, including any tags or SHA256 repo digests.
	Name string `pulumi:"name"`
	// The platform to use when pulling the image. Defaults to the platform of the current machine.
	Platform *string `pulumi:"platform"`
	// List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
	PullTriggers []string `pulumi:"pullTriggers"`
	// A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
	Triggers map[string]interface{} `pulumi:"triggers"`
}

// The set of arguments for constructing a RemoteImage resource.
type RemoteImageArgs struct {
	// Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
	Build RemoteImageBuildPtrInput
	// If true, then the image is removed forcibly when the resource is destroyed.
	ForceRemove pulumi.BoolPtrInput
	// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
	KeepLocally pulumi.BoolPtrInput
	// The name of the Docker image, including any tags or SHA256 repo digests.
	Name pulumi.StringInput
	// The platform to use when pulling the image. Defaults to the platform of the current machine.
	Platform pulumi.StringPtrInput
	// List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
	PullTriggers pulumi.StringArrayInput
	// A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
	Triggers pulumi.MapInput
}

func (RemoteImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteImageArgs)(nil)).Elem()
}

type RemoteImageInput interface {
	pulumi.Input

	ToRemoteImageOutput() RemoteImageOutput
	ToRemoteImageOutputWithContext(ctx context.Context) RemoteImageOutput
}

func (*RemoteImage) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteImage)(nil)).Elem()
}

func (i *RemoteImage) ToRemoteImageOutput() RemoteImageOutput {
	return i.ToRemoteImageOutputWithContext(context.Background())
}

func (i *RemoteImage) ToRemoteImageOutputWithContext(ctx context.Context) RemoteImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteImageOutput)
}

// RemoteImageArrayInput is an input type that accepts RemoteImageArray and RemoteImageArrayOutput values.
// You can construct a concrete instance of `RemoteImageArrayInput` via:
//
//	RemoteImageArray{ RemoteImageArgs{...} }
type RemoteImageArrayInput interface {
	pulumi.Input

	ToRemoteImageArrayOutput() RemoteImageArrayOutput
	ToRemoteImageArrayOutputWithContext(context.Context) RemoteImageArrayOutput
}

type RemoteImageArray []RemoteImageInput

func (RemoteImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemoteImage)(nil)).Elem()
}

func (i RemoteImageArray) ToRemoteImageArrayOutput() RemoteImageArrayOutput {
	return i.ToRemoteImageArrayOutputWithContext(context.Background())
}

func (i RemoteImageArray) ToRemoteImageArrayOutputWithContext(ctx context.Context) RemoteImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteImageArrayOutput)
}

// RemoteImageMapInput is an input type that accepts RemoteImageMap and RemoteImageMapOutput values.
// You can construct a concrete instance of `RemoteImageMapInput` via:
//
//	RemoteImageMap{ "key": RemoteImageArgs{...} }
type RemoteImageMapInput interface {
	pulumi.Input

	ToRemoteImageMapOutput() RemoteImageMapOutput
	ToRemoteImageMapOutputWithContext(context.Context) RemoteImageMapOutput
}

type RemoteImageMap map[string]RemoteImageInput

func (RemoteImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemoteImage)(nil)).Elem()
}

func (i RemoteImageMap) ToRemoteImageMapOutput() RemoteImageMapOutput {
	return i.ToRemoteImageMapOutputWithContext(context.Background())
}

func (i RemoteImageMap) ToRemoteImageMapOutputWithContext(ctx context.Context) RemoteImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteImageMapOutput)
}

type RemoteImageOutput struct{ *pulumi.OutputState }

func (RemoteImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteImage)(nil)).Elem()
}

func (o RemoteImageOutput) ToRemoteImageOutput() RemoteImageOutput {
	return o
}

func (o RemoteImageOutput) ToRemoteImageOutputWithContext(ctx context.Context) RemoteImageOutput {
	return o
}

// Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
func (o RemoteImageOutput) Build() RemoteImageBuildPtrOutput {
	return o.ApplyT(func(v *RemoteImage) RemoteImageBuildPtrOutput { return v.Build }).(RemoteImageBuildPtrOutput)
}

// If true, then the image is removed forcibly when the resource is destroyed.
func (o RemoteImageOutput) ForceRemove() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RemoteImage) pulumi.BoolPtrOutput { return v.ForceRemove }).(pulumi.BoolPtrOutput)
}

// The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
func (o RemoteImageOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteImage) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
func (o RemoteImageOutput) KeepLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RemoteImage) pulumi.BoolPtrOutput { return v.KeepLocally }).(pulumi.BoolPtrOutput)
}

// The name of the Docker image, including any tags or SHA256 repo digests.
func (o RemoteImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The platform to use when pulling the image. Defaults to the platform of the current machine.
func (o RemoteImageOutput) Platform() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemoteImage) pulumi.StringPtrOutput { return v.Platform }).(pulumi.StringPtrOutput)
}

// List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
func (o RemoteImageOutput) PullTriggers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RemoteImage) pulumi.StringArrayOutput { return v.PullTriggers }).(pulumi.StringArrayOutput)
}

// The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`.
func (o RemoteImageOutput) RepoDigest() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteImage) pulumi.StringOutput { return v.RepoDigest }).(pulumi.StringOutput)
}

// A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
func (o RemoteImageOutput) Triggers() pulumi.MapOutput {
	return o.ApplyT(func(v *RemoteImage) pulumi.MapOutput { return v.Triggers }).(pulumi.MapOutput)
}

type RemoteImageArrayOutput struct{ *pulumi.OutputState }

func (RemoteImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemoteImage)(nil)).Elem()
}

func (o RemoteImageArrayOutput) ToRemoteImageArrayOutput() RemoteImageArrayOutput {
	return o
}

func (o RemoteImageArrayOutput) ToRemoteImageArrayOutputWithContext(ctx context.Context) RemoteImageArrayOutput {
	return o
}

func (o RemoteImageArrayOutput) Index(i pulumi.IntInput) RemoteImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RemoteImage {
		return vs[0].([]*RemoteImage)[vs[1].(int)]
	}).(RemoteImageOutput)
}

type RemoteImageMapOutput struct{ *pulumi.OutputState }

func (RemoteImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemoteImage)(nil)).Elem()
}

func (o RemoteImageMapOutput) ToRemoteImageMapOutput() RemoteImageMapOutput {
	return o
}

func (o RemoteImageMapOutput) ToRemoteImageMapOutputWithContext(ctx context.Context) RemoteImageMapOutput {
	return o
}

func (o RemoteImageMapOutput) MapIndex(k pulumi.StringInput) RemoteImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RemoteImage {
		return vs[0].(map[string]*RemoteImage)[vs[1].(string)]
	}).(RemoteImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteImageInput)(nil)).Elem(), &RemoteImage{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteImageArrayInput)(nil)).Elem(), RemoteImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteImageMapInput)(nil)).Elem(), RemoteImageMap{})
	pulumi.RegisterOutputType(RemoteImageOutput{})
	pulumi.RegisterOutputType(RemoteImageArrayOutput{})
	pulumi.RegisterOutputType(RemoteImageMapOutput{})
}
