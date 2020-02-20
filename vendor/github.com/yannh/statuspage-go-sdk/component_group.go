package statuspage

type ComponentGroup struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description,omitempty"`
	Components  []string `json:"components,omitempty"`
}

type ComponentGroupFull struct {
	ComponentGroup
	ID        *string `json:"id"`
	PageID    *string `json:"page_id"`
	Position  *int32  `json:"position"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func CreateComponentGroup(client *Client, pageID string, componentGroup *ComponentGroup) (*ComponentGroupFull, error) {
	var cg ComponentGroupFull
	err := createResource(
		client,
		pageID,
		"component-group",
		struct {
			ComponentGroup *ComponentGroup `json:"component_group"`
		}{componentGroup},
		&cg,
	)

	return &cg, err
}

func GetComponentGroup(client *Client, pageID, componentGroupID string) (*ComponentGroupFull, error) {
	var cg ComponentGroupFull
	err := readResource(
		client,
		pageID,
		componentGroupID,
		"component-group",
		&cg,
	)

	return &cg, err
}

func UpdateComponentGroup(client *Client, pageID, componentGroupID string, componentGroup *ComponentGroup) (*ComponentGroupFull, error) {
	var cg ComponentGroupFull

	err := updateResource(
		client,
		pageID,
		"component-group",
		componentGroupID,
		struct {
			ComponentGroup *ComponentGroup `json:"component_group"`
		}{componentGroup},
		&cg,
	)

	return &cg, err
}

func DeleteComponentGroup(client *Client, pageID, componentGroupID string) (err error) {
	return deleteResource(client, pageID, "component-group", componentGroupID)
}
