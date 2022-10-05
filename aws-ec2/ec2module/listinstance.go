package ec2module

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2DescribeInstancesAPI interface {
	DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

func GetInstance(ctx context.Context, input *ec2.DescribeInstancesInput, client EC2DescribeInstancesAPI) (*ec2.DescribeInstancesOutput, error) {
	return client.DescribeInstances(ctx, input)
}

func ListEc2Instances(ctx context.Context, region string) {

	var resultCount int32 = 6 // Maximum number of result is set as 6

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))

	if err != nil {
		panic(fmt.Sprintf("failed loading config, %v", err))
	}

	ec2Client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInstancesInput{MaxResults: &resultCount}
	result, err := GetInstance(ctx, input, ec2Client)

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
