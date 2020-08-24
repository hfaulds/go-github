// Copyright 2020 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
)

// RunnerGroup represents a self-hosted runner group
type RunnerGroup struct {
	ID     *int64  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	Default *bool `json:"default,omitempty"`
	RunnersURL *string `json:"runners_url,omitempty"`
	SelectedRepositoriesURL *string `json:"selected_repositories_url,omitempty"`
	Inherited *bool `json:"inherited,omitempty"`
}

type CreateRunnerGroupRequest struct {
	Name   string `json:"name,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	SelectedRepositoryIDs []int64 `json:"selected_repository_ids,omitempty"`
	Runners []int64 `json:"runners"`
}

// CreateOrganizationRunnerGroup create a self-hosted runner group for an organization.
// GitHub API docs: https://docs.github.com/en/rest/reference/actions#create-a-self-hosted-runner-group-for-an-organization
func (s *ActionsService) CreateOrganizationRunnerGroup(ctx context.Context, org string, runnerGroupReq *CreateRunnerGroupRequest) (*RunnerGroup, *Response, error) {
	u := fmt.Sprintf("orgs/%v/actions/runner-groups", org)

	req, err := s.client.NewRequest("POST", u, runnerGroupReq)
	if err != nil {
		return nil, nil, err
	}

	runnerGroup := &RunnerGroup{}
	resp, err := s.client.Do(ctx, req, &runnerGroup)
	if err != nil {
		return nil, resp, err
	}

	return runnerGroup, resp, nil
}

// DeleteOrganizationRunnerGroup deletes a self-hosted runner group from an organization.
// GitHub API docs: https://docs.github.com/en/rest/reference/actions#delete-a-self-hosted-runner-group-from-an-organization
func (s *ActionsService) DeleteOrganizationRunnerGroup(ctx context.Context, org string, runnerGroupID int64) (*Response, error) {
	u := fmt.Sprintf("orgs/%v/actions/runner-groups/%d", org, runnerGroupID)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	runnerGroup := &RunnerGroup{}
	resp, err := s.client.Do(ctx, req, &runnerGroup)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
