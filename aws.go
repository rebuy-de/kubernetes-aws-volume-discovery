package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func awsGetVolumes() ([]*ec2.Volume, error) {
	config := aws.NewConfig().WithRegion("eu-central-1")
	svc := ec2.New(session.New(), config)

	resp, err := svc.DescribeVolumes(nil)
	if err != nil {
		return nil, err
	}

	return resp.Volumes, nil
}
