// Copyright 2020 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestActionsService_CreateOrganizationRunnerGroup(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/actions/runner-groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
			  "id": 2,
			  "name": "octo-runner-group",
			  "visibility": "selected",
			  "default": false,
			  "selected_repositories_url": "https://api.github.com/orgs/octo-org/actions/runner-groups/2/repositories",
			  "runners_url": "https://api.github.com/orgs/octo-org/actions/runner_groups/2/runners",
			  "inherited": false
			}`,
		)
	})

	req := &CreateRunnerGroupRequest{
		Name:                  "octo-runner-group",
		Visibility:            "selected",
		SelectedRepositoryIDs: []int64{2},
		Runners:               []int64{3},
	}
	runnerGroup, _, err := client.Actions.CreateOrganizationRunnerGroup(context.Background(), "o", req)
	if err != nil {
		t.Errorf("Actions.CreateRegistrationToken returned error: %v", err)
	}

	id := int64(2)
	name := "octo-runner-group"
	visibility := "selected"
	selectedRepositoriesURL := "https://api.github.com/orgs/octo-org/actions/runner-groups/2/repositories"
	runnersURL := "https://api.github.com/orgs/octo-org/actions/runner_groups/2/runners"
	isDefault := false
	inherited := false
	want := &RunnerGroup{
		ID: &id,
		Name: &name,
		Visibility: &visibility,
		Default: &isDefault,
		SelectedRepositoriesURL: &selectedRepositoriesURL,
		RunnersURL: &runnersURL,
		Inherited: &inherited,
	}
	if !reflect.DeepEqual(runnerGroup, want) {
		t.Errorf("Actions.CreateOrganizationRunnerGroup returned %+v, want %+v", runnerGroup, want)
	}
}

func TestActionsService_DeleteOrganizationRunnerGroup(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/actions/runner-groups/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Actions.DeleteOrganizationRunnerGroup(context.Background(), "o", 2)
	if err != nil {
		t.Errorf("Actions.DeleteOrganizationRunnerGroup returned error: %v", err)
	}
}
