package drds

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

// DescribeDrdsRegions invokes the drds.DescribeDrdsRegions API synchronously
// api document: https://help.aliyun.com/api/drds/describedrdsregions.html
func (client *Client) DescribeDrdsRegions(request *DescribeDrdsRegionsRequest) (response *DescribeDrdsRegionsResponse, err error) {
	response = CreateDescribeDrdsRegionsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsRegionsWithChan invokes the drds.DescribeDrdsRegions API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsregions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsRegionsWithChan(request *DescribeDrdsRegionsRequest) (<-chan *DescribeDrdsRegionsResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsRegionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsRegions(request)
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

// DescribeDrdsRegionsWithCallback invokes the drds.DescribeDrdsRegions API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsregions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsRegionsWithCallback(request *DescribeDrdsRegionsRequest, callback func(response *DescribeDrdsRegionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsRegionsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsRegions(request)
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

// DescribeDrdsRegionsRequest is the request struct for api DescribeDrdsRegions
type DescribeDrdsRegionsRequest struct {
	*requests.RpcRequest
}

// DescribeDrdsRegionsResponse is the response struct for api DescribeDrdsRegions
type DescribeDrdsRegionsResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Regions   Regions `json:"Regions" xml:"Regions"`
}

// CreateDescribeDrdsRegionsRequest creates a request to invoke DescribeDrdsRegions API
func CreateDescribeDrdsRegionsRequest() (request *DescribeDrdsRegionsRequest) {
	request = &DescribeDrdsRegionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsRegions", "drds", "openAPI")
	return
}

// CreateDescribeDrdsRegionsResponse creates a response to parse from DescribeDrdsRegions response
func CreateDescribeDrdsRegionsResponse() (response *DescribeDrdsRegionsResponse) {
	response = &DescribeDrdsRegionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}