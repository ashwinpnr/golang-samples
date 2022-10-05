package ec2module

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type EC2DescribeInstancesAPIImpl struct {
}

func (e *EC2DescribeInstancesAPIImpl) DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {

	reservations := []types.Reservation{
		{
			ReservationId: aws.String("aws-test-reservationID"),
			Instances: []types.Instance{
				{InstanceId: aws.String("aws-test-instanceID1")},
				{InstanceId: aws.String("aws-test-instanceID2")},
			},
		},
	}

	output := &ec2.DescribeInstancesOutput{
		Reservations: reservations,
	}

	return output, nil
}

func TestDescribeInstance(t *testing.T) {
	thisTime := time.Now()
	nowString := thisTime.Format("2006-01-02 15:04:05 Monday")
	t.Log("Starting unit test at " + nowString)

	ctx := context.Background()
	input := &ec2.DescribeInstancesInput{}
	client := EC2DescribeInstancesAPIImpl{}

	result, err := GetInstance(ctx, input, &client)

	if err != nil {
		t.Error("Got an error retrieving information about your Amazon EC2 instances:")
		t.Error(err)
		return
	}

	for _, r := range result.Reservations {
		t.Log("Reservation ID: " + *r.ReservationId)
		t.Log("Instance IDs:")
		for _, i := range r.Instances {
			t.Log("   " + *i.InstanceId)
		}

		t.Log("")
	}
}
