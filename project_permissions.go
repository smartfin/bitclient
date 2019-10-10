package bitclient

import "fmt"

type GetProjectUserPermissionRequest struct {
	PagedRequest
	Filter string
}

type GetProjectUserPermissionResponse struct {
	PagedResponse
	Values []UserPermission
}

func (bc *BitClient) GetProjectUserPermission(projectKey string, params GetProjectUserPermissionRequest) (GetProjectUserPermissionResponse, error) {
	response := GetProjectUserPermissionResponse{}

	_, err := bc.DoGet(
		fmt.Sprintf("/projects/%s/permissions/users", projectKey),
		params,
		&response,
	)

	return response, err
}

type SetUserPermissionRequest struct {
	Users      []string `json:"name" url:"name,omitempty"`
	Permission string   `json:"permission" url:"permission,omitempty"`
}

func (bc *BitClient) SetProjectUserPermission(projectKey string, params SetUserPermissionRequest) error {
	response := GetProjectUserPermissionResponse{}

	_, err := bc.DoPutUrl(
		fmt.Sprintf("/projects/%s/permissions/users", projectKey),
		params,
		&response,
	)

	return err
}

type GetProjectGroupPermissionRequest struct {
	PagedRequest
	Filter string
}

type GetProjectGroupPermissionResponse struct {
	PagedResponse
	Values []GroupPermission
}

func (bc *BitClient) GetProjectGroupPermission(projectKey string, params GetProjectGroupPermissionRequest) (GetProjectGroupPermissionResponse, error) {
	response := GetProjectGroupPermissionResponse{}

	_, err := bc.DoGet(
		fmt.Sprintf("/projects/%s/permissions/groups", projectKey),
		params,
		&response,
	)

	return response, err
}
