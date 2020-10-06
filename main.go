package main

import (
	"fmt"
	"log"

	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	resourceGroupName := "octopuspulumitest"
	location := "eastus"
	ctx := &pulumi.Context{}

	newResourceGroup(ctx, resourceGroupName, location)

}

func newResourceGroup(ctx *pulumi.Context, resourceGroupName string, location string) {
	if ctx == nil {
		log.Println("Pulumi CTX is not working as expected... please check issues on the SDK: github.com/pulumi/pulumi/sdk/v2/go/pulumi")
	} else {
		pulumi.Run(func(ctx *pulumi.Context) error {
			resourceGroup, err := core.NewResourceGroup(ctx, resourceGroupName, &core.ResourceGroupArgs{Location: pulumi.String(location)})
			if err != nil {
				log.Println(err)
			}

			fmt.Println(resourceGroup)
			return nil
		})
	}
}
