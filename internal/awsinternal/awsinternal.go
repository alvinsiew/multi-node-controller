package awsinternal

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// FilterInstances to get instance ip addresses base on filtering pass in
func FilterInstances(filterName string) ([]string, []string) {
	var servers []string
	var names []string
	awsRegion := "ap-southeast-1"

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create an EC2 service client.
	svc := ec2.New(sess, &aws.Config{Region: aws.String(awsRegion)})

	// Make the API request to EC2 filtering for the addresses in the
	// account's VPC.
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(filterName),
				},
			},
		},
	}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		
	}

	// fmt.Println(result.Reservations)
	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			if *i.State.Name == "running" {
				for _, t := range i.Tags {
					if *t.Key == "Name" {
						names = append(names, (*t.Value))
					}
				}
				for _, r := range i.NetworkInterfaces {
					servers = append(servers, (*r.PrivateIpAddress))
				}
			}
		}
	}
	return servers, names
}
