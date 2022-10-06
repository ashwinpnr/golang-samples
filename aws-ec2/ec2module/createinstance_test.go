package ec2module

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type EC2RunInstanceAPIImpl struct {
}

func (e *EC2DescribeInstancesAPIImpl) RunInstances(ctx context.Context, params *ec2.RunInstancesInput, optFns ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error) {
	instances := []types.Instance{
		{InstanceId: aws.String("test-instance-id")},
	}

	output := &ec2.RunInstancesOutput{
		Instances: instances,
	}

	return output, nil
}

func TestRunInstance(t *testing.T) {
	thisTime := time.Now()
	nowString := thisTime.Format("2006-01-02 15:04:05 Monday")
	t.Log("Starting unit test at " + nowString)
	client := EC2DescribeInstancesAPIImpl{}
	ctx := context.Background()
	runInstance, err := CreateEc2Instance(ctx, &ec2.RunInstancesInput{}, &client)
	if err != nil {
		t.Errorf("RunInstance error: %s", err)
	}

	if len(runInstance.Instances) == 0 {
		t.Errorf("RunInstance has empty length (%d)", len(runInstance.Instances))
	}

	t.Logf("Instance Id %s", *runInstance.Instances[0].InstanceId)

}
