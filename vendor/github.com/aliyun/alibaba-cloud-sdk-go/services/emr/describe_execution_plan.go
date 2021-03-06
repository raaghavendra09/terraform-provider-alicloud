package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeExecutionPlan invokes the emr.DescribeExecutionPlan API synchronously
// api document: https://help.aliyun.com/api/emr/describeexecutionplan.html
func (client *Client) DescribeExecutionPlan(request *DescribeExecutionPlanRequest) (response *DescribeExecutionPlanResponse, err error) {
	response = CreateDescribeExecutionPlanResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeExecutionPlanWithChan invokes the emr.DescribeExecutionPlan API asynchronously
// api document: https://help.aliyun.com/api/emr/describeexecutionplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeExecutionPlanWithChan(request *DescribeExecutionPlanRequest) (<-chan *DescribeExecutionPlanResponse, <-chan error) {
	responseChan := make(chan *DescribeExecutionPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeExecutionPlan(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeExecutionPlanWithCallback invokes the emr.DescribeExecutionPlan API asynchronously
// api document: https://help.aliyun.com/api/emr/describeexecutionplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeExecutionPlanWithCallback(request *DescribeExecutionPlanRequest, callback func(response *DescribeExecutionPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeExecutionPlanResponse
		var err error
		defer close(result)
		response, err = client.DescribeExecutionPlan(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeExecutionPlanRequest is the request struct for api DescribeExecutionPlan
type DescribeExecutionPlanRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Id              string           `position:"Query" name:"Id"`
}

// DescribeExecutionPlanResponse is the response struct for api DescribeExecutionPlan
type DescribeExecutionPlanResponse struct {
	*responses.BaseResponse
	RequestId             string      `json:"RequestId" xml:"RequestId"`
	Id                    string      `json:"Id" xml:"Id"`
	Name                  string      `json:"Name" xml:"Name"`
	Status                string      `json:"Status" xml:"Status"`
	Strategy              string      `json:"Strategy" xml:"Strategy"`
	TimeInterval          int         `json:"TimeInterval" xml:"TimeInterval"`
	StartTime             int64       `json:"StartTime" xml:"StartTime"`
	TimeUnit              string      `json:"TimeUnit" xml:"TimeUnit"`
	DayOfWeek             string      `json:"DayOfWeek" xml:"DayOfWeek"`
	DayOfMonth            string      `json:"DayOfMonth" xml:"DayOfMonth"`
	CreateClusterOnDemand bool        `json:"CreateClusterOnDemand" xml:"CreateClusterOnDemand"`
	ClusterId             string      `json:"ClusterId" xml:"ClusterId"`
	ClusterName           string      `json:"ClusterName" xml:"ClusterName"`
	WorkflowApp           string      `json:"WorkflowApp" xml:"WorkflowApp"`
	ExecutionPlanVersion  int64       `json:"ExecutionPlanVersion" xml:"ExecutionPlanVersion"`
	ClusterInfo           ClusterInfo `json:"ClusterInfo" xml:"ClusterInfo"`
	JobInfoList           JobInfoList `json:"JobInfoList" xml:"JobInfoList"`
}

// CreateDescribeExecutionPlanRequest creates a request to invoke DescribeExecutionPlan API
func CreateDescribeExecutionPlanRequest() (request *DescribeExecutionPlanRequest) {
	request = &DescribeExecutionPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeExecutionPlan", "emr", "openAPI")
	return
}

// CreateDescribeExecutionPlanResponse creates a response to parse from DescribeExecutionPlan response
func CreateDescribeExecutionPlanResponse() (response *DescribeExecutionPlanResponse) {
	response = &DescribeExecutionPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
