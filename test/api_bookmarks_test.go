/*
Delphix DCT API

Testing BookmarksApiService

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

func Test_delphix_dct_api_BookmarksApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BookmarksApiService CreateBookmark", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BookmarksApi.CreateBookmark(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BookmarksApiService CreateBookmarkTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bookmarkId string

		resp, httpRes, err := apiClient.BookmarksApi.CreateBookmarkTags(context.Background(), bookmarkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BookmarksApiService DeleteBookmark", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bookmarkId string

		httpRes, err := apiClient.BookmarksApi.DeleteBookmark(context.Background(), bookmarkId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BookmarksApiService DeleteBookmarkTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bookmarkId string

		httpRes, err := apiClient.BookmarksApi.DeleteBookmarkTags(context.Background(), bookmarkId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BookmarksApiService GetBookmarkById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bookmarkId string

		resp, httpRes, err := apiClient.BookmarksApi.GetBookmarkById(context.Background(), bookmarkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BookmarksApiService GetBookmarkTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bookmarkId string

		resp, httpRes, err := apiClient.BookmarksApi.GetBookmarkTags(context.Background(), bookmarkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BookmarksApiService GetBookmarks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BookmarksApi.GetBookmarks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BookmarksApiService GetVdbGroupsByBookmark", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bookmarkId string

		resp, httpRes, err := apiClient.BookmarksApi.GetVdbGroupsByBookmark(context.Background(), bookmarkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BookmarksApiService SearchBookmarks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BookmarksApi.SearchBookmarks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BookmarksApiService UpdateBookmark", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bookmarkId string

		resp, httpRes, err := apiClient.BookmarksApi.UpdateBookmark(context.Background(), bookmarkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
