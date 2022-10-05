package ec2module

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type EC2RunInstanceAPI interface {
	RunInstances(ctx context.Context, params *ec2.RunInstancesInput, optFns ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error)
}

func CreateEc2Instance(ctx context.Context, params *ec2.RunInstancesInput, client EC2RunInstanceAPI) (*ec2.RunInstancesOutput, error) {
	return client.RunInstances(ctx, params)
}

func CreateEC2(ctx context.Context, region string) (string, error) {
	name := "Name"
	value := "golang-instance"
	secretName := "go-aws-ec2-pem"
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return "", fmt.Errorf("LoadDefaultConfig error: %s", err)
	}

	ec2Client := ec2.NewFromConfig(cfg)

	existingKeyPairs, err := ec2Client.DescribeKeyPairs(ctx, &ec2.DescribeKeyPairsInput{
		KeyNames: []string{"go-aws-ec2"},
	})
	if err != nil && !strings.Contains(err.Error(), "InvalidKeyPair.NotFound") {
		return "", fmt.Errorf("DescribeKeyPairs error: %s", err)
	}

	if existingKeyPairs == nil || len(existingKeyPairs.KeyPairs) == 0 {
		keyPair, err := ec2Client.CreateKeyPair(ctx, &ec2.CreateKeyPairInput{
			KeyName: aws.String("go-aws-ec2"),
		})
		if err != nil {
			return "", fmt.Errorf("CreateKeyPair error: %s", err)
		}

		err = os.WriteFile(secretName+".pem", []byte(*keyPair.KeyMaterial), 0600)
		if err != nil {
			return "", fmt.Errorf("WriteFile (keypair) error: %s", err)
		}

		// Push PEM file contents into AWS Secret manager
		secretClient := secretsmanager.NewFromConfig(cfg)
		secretDescription := "Contain contents of file " + secretName + ".pem"
		secretClient.CreateSecret(ctx, &secretsmanager.CreateSecretInput{Name: &secretName,
			Description:  &secretDescription,
			SecretString: keyPair.KeyMaterial})

	}

	describeImages, err := ec2Client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("name"),
				Values: []string{"ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"},
			},
			{
				Name:   aws.String("virtualization-type"),
				Values: []string{"hvm"},
			},
		},
		Owners: []string{"099720109477"},
	})
	if err != nil {
		return "", fmt.Errorf("DescribeImages error: %s", err)
	}
	if len(describeImages.Images) == 0 {
		return "", fmt.Errorf("describeImages has empty length (%d)", len(describeImages.Images))
	}

	runInstance, err := CreateEc2Instance(ctx, &ec2.RunInstancesInput{
		ImageId:      describeImages.Images[0].ImageId,
		InstanceType: types.InstanceTypeT3Micro,
		KeyName:      aws.String("go-aws-ec2"),
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
		TagSpecifications: []types.TagSpecification{{ResourceType: types.ResourceTypeInstance,
			Tags: []types.Tag{
				{
					Key:   &name,
					Value: &value,
				},
			},
		},
		},
	}, ec2Client)

	if err != nil {
		return "", fmt.Errorf("RunInstance error: %s", err)
	}

	if len(runInstance.Instances) == 0 {
		return "", fmt.Errorf("RunInstance has empty length (%d)", len(runInstance.Instances))
	}

	return *runInstance.Instances[0].InstanceId, nil
}
