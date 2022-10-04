package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2Instances(ctx context.Context, region string) {

	var resultCount int32 = 6

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))

	if err != nil {
		panic(fmt.Sprintf("failed loading config, %v", err))
	}

	ec2Client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInstancesInput{MaxResults: &resultCount}
	result, err := ec2Client.DescribeInstances(ctx, input)

	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		fmt.Println(err)
		return
	}

	for _, r := range result.Reservations {
		fmt.Println("Reservation ID: " + *r.ReservationId)
		fmt.Println("Instance IDs:")
		for _, i := range r.Instances {
			fmt.Println("   " + *i.InstanceId)
		}

		fmt.Println("")
	}

}
func main() {
	var (
		instanceId string
		err        error
	)
	ctx := context.Background()
	region := "ap-south-1"
	//ListEc2Instances(ctx, region)

	if instanceId, err = createEC2(ctx, region); err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	fmt.Printf("instance id: %s\n", instanceId)
}
