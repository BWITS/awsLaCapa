package models

import (
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/ec2"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/awserr"
        "github.com/astaxie/beego"
)

/* At the moment, I couldn't figure out a way to use type alias to generate
 object model. Hence, use a model to describe ec2.DescribeVpcs and not being
 used anywhere else except for Swagger API documentation
 */
type VPC struct {
        CidrBlock       string `json:"CidrBlock"`
        DhcpOptionsID   string `json:"DhcpOptionsId"`
        InstanceTenancy string `json:"InstanceTenancy"`
        IsDefault       bool `json:"IsDefault"`
        State           string `json:"State"`
        Tags []struct {
            Key string `json:"Key"`
            Value string `json:"Value"`
        } `json:"Tags"`
        VpcID string `json:"VpcId"`
}

type VPCs struct {
        VPCs []VPC `json:"vpcs"`
}

func VPCList() (*ec2.DescribeVpcsOutput, error) {
        svc := ec2.New(session.New(), &aws.Config{
                Region: aws.String("ap-southeast-2"),
        })

        params := &ec2.DescribeVpcsInput{}

        resp, err := svc.DescribeVpcs(params)

        if err != nil {
          if awsErr, ok := err.(awserr.Error); ok {
           beego.Error(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
           if reqErr, ok := err.(awserr.RequestFailure); ok {
               beego.Error(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
            }
          } else {
              beego.Debug(err.Error())
          }
          return nil, err
        }
        return resp, nil
}
