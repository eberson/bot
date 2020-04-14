package aws

import (
	"context"
	"os"

	"github.com/eberson/rootinha/helper/strs"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)

type APICredential struct {
	AccessKey string
	SecretKey string
	Region    string
}

type AWS struct {
	config aws.Config
}

func (a *AWS) build(projectName string) (string, error) {
	svc := codebuild.New(a.config)

	req := svc.StartBuildRequest(&codebuild.StartBuildInput{
		ProjectName: aws.String(projectName),
	})

	out, err := req.Send(context.TODO())

	if err != nil {
		return strs.Empty(), err
	}

	return *out.Build.Id, nil
}

func newAWS(secretKey, accessKey, region string) (*AWS, error) {
	_ = os.Setenv(external.AWSAccessKeyEnvVar, accessKey)
	_ = os.Setenv(external.AWSSecreteKeyEnvVar, secretKey)

	cfg, err := external.LoadDefaultAWSConfig()

	if err != nil {
		return nil, err
	}

	cfg.Region = region

	return &AWS{
		config: cfg,
	}, nil
}
