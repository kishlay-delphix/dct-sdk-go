/*
Delphix DCT API

Testing ExecutionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package delphix_dct_api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/delphix/dct-sdk-go"
)

func Test_delphix_dct_api_ExecutionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExecutionsApiService CancelExecution", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var executionId string

		httpRes, err := apiClient.ExecutionsApi.CancelExecution(context.Background(), executionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExecutionsApiService GetExecutionById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var executionId string

		resp, httpRes, err := apiClient.ExecutionsApi.GetExecutionById(context.Background(), executionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExecutionsApiService GetExecutionEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var executionId string

		resp, httpRes, err := apiClient.ExecutionsApi.GetExecutionEvents(context.Background(), executionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExecutionsApiService GetExecutions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExecutionsApi.GetExecutions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExecutionsApiService SearchExecutionEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var executionId string

		resp, httpRes, err := apiClient.ExecutionsApi.SearchExecutionEvents(context.Background(), executionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExecutionsApiService SearchExecutions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExecutionsApi.SearchExecutions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
