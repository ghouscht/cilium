// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for ModifyVpcEndpoint.
type ModifyVpcEndpointInput struct {
	_ struct{} `type:"structure"`

	// (Gateway endpoint) One or more route tables IDs to associate with the endpoint.
	AddRouteTableIds []string `locationName:"AddRouteTableId" locationNameList:"item" type:"list"`

	// (Interface endpoint) One or more security group IDs to associate with the
	// network interface.
	AddSecurityGroupIds []string `locationName:"AddSecurityGroupId" locationNameList:"item" type:"list"`

	// (Interface endpoint) One or more subnet IDs in which to serve the endpoint.
	AddSubnetIds []string `locationName:"AddSubnetId" locationNameList:"item" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// A policy to attach to the endpoint that controls access to the service. The
	// policy must be in valid JSON format.
	PolicyDocument *string `type:"string"`

	// (Interface endpoint) Indicates whether a private hosted zone is associated
	// with the VPC.
	PrivateDnsEnabled *bool `type:"boolean"`

	// (Gateway endpoint) One or more route table IDs to disassociate from the endpoint.
	RemoveRouteTableIds []string `locationName:"RemoveRouteTableId" locationNameList:"item" type:"list"`

	// (Interface endpoint) One or more security group IDs to disassociate from
	// the network interface.
	RemoveSecurityGroupIds []string `locationName:"RemoveSecurityGroupId" locationNameList:"item" type:"list"`

	// (Interface endpoint) One or more subnets IDs in which to remove the endpoint.
	RemoveSubnetIds []string `locationName:"RemoveSubnetId" locationNameList:"item" type:"list"`

	// (Gateway endpoint) Specify true to reset the policy document to the default
	// policy. The default policy allows full access to the service.
	ResetPolicy *bool `type:"boolean"`

	// The ID of the endpoint.
	//
	// VpcEndpointId is a required field
	VpcEndpointId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyVpcEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpcEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVpcEndpointInput"}

	if s.VpcEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcEndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyVpcEndpointOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s ModifyVpcEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyVpcEndpoint = "ModifyVpcEndpoint"

// ModifyVpcEndpointRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies attributes of a specified VPC endpoint. The attributes that you
// can modify depend on the type of VPC endpoint (interface or gateway). For
// more information, see VPC Endpoints (https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using ModifyVpcEndpointRequest.
//    req := client.ModifyVpcEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVpcEndpoint
func (c *Client) ModifyVpcEndpointRequest(input *ModifyVpcEndpointInput) ModifyVpcEndpointRequest {
	op := &aws.Operation{
		Name:       opModifyVpcEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpcEndpointInput{}
	}

	req := c.newRequest(op, input, &ModifyVpcEndpointOutput{})

	return ModifyVpcEndpointRequest{Request: req, Input: input, Copy: c.ModifyVpcEndpointRequest}
}

// ModifyVpcEndpointRequest is the request type for the
// ModifyVpcEndpoint API operation.
type ModifyVpcEndpointRequest struct {
	*aws.Request
	Input *ModifyVpcEndpointInput
	Copy  func(*ModifyVpcEndpointInput) ModifyVpcEndpointRequest
}

// Send marshals and sends the ModifyVpcEndpoint API request.
func (r ModifyVpcEndpointRequest) Send(ctx context.Context) (*ModifyVpcEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyVpcEndpointResponse{
		ModifyVpcEndpointOutput: r.Request.Data.(*ModifyVpcEndpointOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyVpcEndpointResponse is the response type for the
// ModifyVpcEndpoint API operation.
type ModifyVpcEndpointResponse struct {
	*ModifyVpcEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyVpcEndpoint request.
func (r *ModifyVpcEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}