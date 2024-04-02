package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vpc, err := ec2.NewVpc(ctx, "nat_host_vpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.128.0/18"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("nat_host_vpc"),
			},
		})
		if err != nil {
			return err
		}

		subnet, err := ec2.NewSubnet(ctx, "private", &ec2.SubnetArgs{
			VpcId:               vpc.ID(),
			CidrBlock:           pulumi.String("10.0.128.0/20"),
			MapPublicIpOnLaunch: pulumi.Bool(true),
			AvailabilityZone:    pulumi.String("ap-northeast-1a"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("private"),
			},
		})
		if err != nil {
			return err
		}

		igw, err := ec2.NewInternetGateway(ctx, "igw", &ec2.InternetGatewayArgs{
			VpcId: vpc.ID(),
		})
		if err != nil {
			return err
		}

		_, err = ec2.NewInternetGatewayAttachment(ctx, "igw_att", &ec2.InternetGatewayAttachmentArgs{
			VpcId:             vpc.ID(),
			InternetGatewayId: igw.ID(),
		})
		if err != nil {
			return err
		}

		eip, err := ec2.NewEip(ctx, "eip", &ec2.EipArgs{
			Domain: pulumi.String("vpc"),
		})
		if err != nil {
			return err
		}

		_, err = ec2.NewNatGateway(ctx, "nat", &ec2.NatGatewayArgs{
			SubnetId:     subnet.ID(),
			AllocationId: eip.ID(),
		})
		if err != nil {
			return err
		}

		vpc1, err := ec2.NewVpc(ctx, "nat_multiple_vpc1", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/18"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("nat_multiple_vpc1"),
			},
		})
		if err != nil {
			return err
		}

		_, err = ec2.NewSubnet(ctx, "private1", &ec2.SubnetArgs{
			VpcId:               vpc1.ID(),
			CidrBlock:           pulumi.String("10.0.0.0/20"),
			MapPublicIpOnLaunch: pulumi.Bool(false),
			AvailabilityZone:    pulumi.String("ap-northeast-1a"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("private1"),
			},
		})
		if err != nil {
			return err
		}

		vpc2, err := ec2.NewVpc(ctx, "nat_multiple_vpc2", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.64.0/18"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("nat_multiple_vpc2"),
			},
		})
		if err != nil {
			return err
		}

		_, err = ec2.NewSubnet(ctx, "private2", &ec2.SubnetArgs{
			VpcId:               vpc2.ID(),
			CidrBlock:           pulumi.String("10.0.64.0/20"),
			MapPublicIpOnLaunch: pulumi.Bool(false),
			AvailabilityZone:    pulumi.String("ap-northeast-1a"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("private2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
