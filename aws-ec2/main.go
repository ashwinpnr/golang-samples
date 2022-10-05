package main

import (
	"context"

	"github.com/ashwinpnr/golang-samples/aws-ec2/ec2module"
)

func main() {

	ctx := context.Background()
	region := "ap-south-1"
	ec2module.ListEc2Instances(ctx, region)
	/*
		var (
			instanceId string
			err        error
		)
			if instanceId, err = ec2.CreateEC2(ctx, region); err != nil {
				fmt.Printf("Error: %s", err)
				os.Exit(1)
			}

			fmt.Printf("instance id: %s\n", instanceId)
	*/
}
