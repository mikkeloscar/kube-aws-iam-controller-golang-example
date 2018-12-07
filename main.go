package main

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	svc := ec2.New(sess)

	for {

		log.Print("Getting instances")
		params := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				{
					Name: aws.String("instance-state-name"),
					Values: []*string{
						aws.String("running"),
					},
				},
			},
		}
		resp, err := svc.DescribeInstances(params)
		if err != nil {
			log.Printf("ERROR: %v", err)
		} else {
			for _, res := range resp.Reservations {
				for _, instance := range res.Instances {
					log.Printf("%s - %s",
						aws.StringValue(instance.InstanceId),
						aws.StringValue(instance.InstanceType),
					)
				}
			}
		}

		// sleep for 5 min.
		time.Sleep(5 * time.Minute)
	}
}
