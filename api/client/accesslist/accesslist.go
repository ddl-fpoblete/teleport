// Copyright 2023 Gravitational, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package accesslist

import (
	"context"

	"github.com/gravitational/trace"

	accesslistv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/accesslist/v1"
	"github.com/gravitational/teleport/api/types/accesslist"
	conv "github.com/gravitational/teleport/api/types/accesslist/convert/v1"
)

// Client is an access list client that conforms to the following lib/services interfaces:
// * services.AccessLists
type Client struct {
	grpcClient accesslistv1.AccessListServiceClient
}

// NewClient creates a new Access List client.
func NewClient(grpcClient accesslistv1.AccessListServiceClient) *Client {
	return &Client{
		grpcClient: grpcClient,
	}
}

// GetAccessLists returns a list of all access lists.
func (c *Client) GetAccessLists(ctx context.Context) ([]*accesslist.AccessList, error) {
	resp, err := c.grpcClient.GetAccessLists(ctx, &accesslistv1.GetAccessListsRequest{})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	accessLists := make([]*accesslist.AccessList, len(resp.AccessLists))
	for i, accessList := range resp.AccessLists {
		var err error
		accessLists[i], err = conv.FromProto(accessList)
		if err != nil {
			return nil, trace.Wrap(err)
		}
	}

	return accessLists, nil
}

// ListAccessLists returns a paginated list of access lists.
func (c *Client) ListAccessLists(ctx context.Context, pageSize int, nextToken string) ([]*accesslist.AccessList, string, error) {
	resp, err := c.grpcClient.ListAccessLists(ctx, &accesslistv1.ListAccessListsRequest{
		PageSize:  int32(pageSize),
		NextToken: nextToken,
	})
	if err != nil {
		return nil, "", trace.Wrap(err)
	}

	accessLists := make([]*accesslist.AccessList, len(resp.AccessLists))
	for i, accessList := range resp.AccessLists {
		var err error
		accessLists[i], err = conv.FromProto(accessList)
		if err != nil {
			return nil, "", trace.Wrap(err)
		}
	}

	return accessLists, resp.GetNextToken(), nil
}

// GetAccessList returns the specified access list resource.
func (c *Client) GetAccessList(ctx context.Context, name string) (*accesslist.AccessList, error) {
	resp, err := c.grpcClient.GetAccessList(ctx, &accesslistv1.GetAccessListRequest{
		Name: name,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	accessList, err := conv.FromProto(resp)
	return accessList, trace.Wrap(err)
}

// UpsertAccessList creates or updates an access list resource.
func (c *Client) UpsertAccessList(ctx context.Context, accessList *accesslist.AccessList) (*accesslist.AccessList, error) {
	resp, err := c.grpcClient.UpsertAccessList(ctx, &accesslistv1.UpsertAccessListRequest{
		AccessList: conv.ToProto(accessList),
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}
	responseAccessList, err := conv.FromProto(resp)
	return responseAccessList, trace.Wrap(err)
}

// DeleteAccessList removes the specified access list resource.
func (c *Client) DeleteAccessList(ctx context.Context, name string) error {
	_, err := c.grpcClient.DeleteAccessList(ctx, &accesslistv1.DeleteAccessListRequest{
		Name: name,
	})
	return trace.Wrap(err)
}

// DeleteAllAccessLists removes all access lists.
func (c *Client) DeleteAllAccessLists(ctx context.Context) error {
	_, err := c.grpcClient.DeleteAllAccessLists(ctx, &accesslistv1.DeleteAllAccessListsRequest{})
	return trace.Wrap(err)
}
