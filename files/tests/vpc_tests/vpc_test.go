package aws

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testConf() map[string]string {
	return map[string]string{
		"availability_zones.#": "2",
		"availability_zones.0": "us-east-1a",
		"availability_zones.1": "us-east-1b",
		"availability_zones.3": "us-east-1c",
		"availability_zones.4": "us-east-1d",
		"availability_zones.5": "us-east-1e",
		"availability_zones.6": "us-east-1f",
	}
}

func TestTerraformAwsNetworkExample(t *testing.T) {
	t.Parallel()

	// Pick a random AWS region to test in. This helps ensure your code works in all regions.
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)
	awsZones, _ := aws.GetAvailabilityZonesE(t, awsRegion)
	awsZone := make([]string, 1)
	awsZone[0] = awsZones[0]

	// Give the VPC and the subnets correct CIDRs
	base_cidr_block := "10.1.0.0/16"
	base_cidr_block_public := "10.1.1.0/20"
	base_cidr_block_private := "10.1.100.0/20"

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../../aws_vpc",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"base_cidr_block":         base_cidr_block,
			"base_cidr_block_public":  base_cidr_block_public,
			"base_cidr_block_private": base_cidr_block_private,
			"aws_region":              awsRegion,
			"availability_zones":      awsZone,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	// publicSubnetId := terraform.Output(t, terraformOptions, "public_subnet_id")
	privateSubnetId := terraform.Output(t, terraformOptions, "private_subnet_id")
	vpcId := terraform.Output(t, terraformOptions, "main_vpc_id")

	subnets := aws.GetSubnetsForVpc(t, vpcId, awsRegion)

	require.Equal(t, 2, len(subnets))
	// Verify if the network that is supposed to be public is really public
	// Will not pass public test until IGW is configured
	// assert.False(t, aws.IsPublicSubnet(t, publicSubnetId, awsRegion))
	// Verify if the network that is supposed to be private is really private
	assert.False(t, aws.IsPublicSubnet(t, privateSubnetId, awsRegion))
}
