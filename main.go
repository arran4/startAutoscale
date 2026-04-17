package main

import (
	"context"
	"flag"
	"log"
	"os/user"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

var (
	group      = flag.String("group", "", "The group you wish to set to 1")
	region     = flag.String("region", "ap-southeast-2", "Region")
	awsConfig  = flag.String("awscfg", filepath.Join(homeDir(), ".aws/credentials"), "Config file to load credentials from")
	awsProfile = flag.String("awsprofile", "default", "Profile to use")
)

func homeDir() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}
	return u.HomeDir
}

func main() {
	flag.Parse()

	if *group == "" {
		log.Fatalf("Please specify the name of the autoscale group using the -group flag")
	}

	ctx := context.Background()

	// Load the Shared AWS Configuration (~/.aws/config)
	var loadOptions []func(*config.LoadOptions) error

	if *region != "" {
		loadOptions = append(loadOptions, config.WithRegion(*region))
	}

	if *awsConfig != "" {
		loadOptions = append(loadOptions, config.WithSharedCredentialsFiles([]string{*awsConfig}))
	}

	if *awsProfile != "" {
		loadOptions = append(loadOptions, config.WithSharedConfigProfile(*awsProfile))
	}

	cfg, err := config.LoadDefaultConfig(ctx, loadOptions...)
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	svc := autoscaling.NewFromConfig(cfg)

	params := &autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String(*group),
		DesiredCapacity:      aws.Int32(1),
		MaxSize:              aws.Int32(1),
		MinSize:              aws.Int32(1),
	}

	log.Printf("Updating Auto Scaling Group '%s' to capacity 1...", *group)
	_, err = svc.UpdateAutoScalingGroup(ctx, params)

	if err != nil {
		log.Fatalf("Failed to update Auto Scaling Group: %v", err)
	}

	log.Printf("Successfully updated Auto Scaling Group '%s'.", *group)
}
