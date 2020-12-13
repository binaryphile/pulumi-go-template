package main

import (
	do "github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a DigitalOcean resource (SpacesBucket)
		bucket, err := do.NewSpacesBucket(ctx, "hello-world", &do.SpacesBucketArgs{
			Acl:    pulumi.String("public-read"),
			Region: pulumi.String("nyc3"),
		})
		if err != nil {
			return err
		}

		_, err = do.NewSpacesBucketObject(ctx, "index.html", &do.SpacesBucketObjectArgs{
			Bucket:      bucket.Name,
			Key:         pulumi.String("index.html"),
			Acl:         pulumi.String("public-read"),
			Content:     pulumi.String("<html>Hello, World!</html>"),
			ContentType: pulumi.String("text/html"),
			Region:      pulumi.String("nyc3"),
		},
			pulumi.DependsOn([]pulumi.Resource{bucket}),
		)
		if err != nil {
			return err
		}

		ctx.Export("bucketDomainName", bucket.BucketDomainName)

		return nil
	})
}
