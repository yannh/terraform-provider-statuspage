package statuspage

type Component struct {
	Name               *string `json:"name"`
	Description        *string `json:"description,omitempty"`
	GroupID            *string `json:"group_id,omitempty"`
	Showcase           *bool   `json:"showcase,omitempty"`
	Status             *string `json:"status,omitempty"`
	OnlyShowIfDegraded *bool   `json:"only_show_if_degraded,omitempty"`
}

type ComponentFull struct {
	Component
	ID              *string `json:"id"`
	PageID          *string `json:"page_id"`
	Position        *int32  `json:"position"`
	CreatedAt       *string `json:"created_at"`
	UpdatedAt       *string `json:"updated_at"`
	AutomationEmail *string `json:"automation_email"`
}

func CreateComponent(client *Client, pageID string, component *Component) (*ComponentFull, error) {
	var c ComponentFull
	err := createResource(
		client,
		pageID,
		"component",
		struct {
			Component *Component `json:"component"`
		}{component},
		&c,
	)

	return &c, err
}

func GetComponent(client *Client, pageID string, componentID string) (*ComponentFull, error) {
	var c ComponentFull
	err := readResource(client, pageID, componentID, "component", &c)

	return &c, err
}

func UpdateComponent(client *Client, pageID, componentID string, component *Component) (*ComponentFull, error) {
	var c ComponentFull

	err := updateResource(
		client,
		pageID,
		"component",
		componentID,
		struct {
			Component *Component `json:"component"`
		}{component},
		&c,
	)

	return &c, err
}

func DeleteComponent(client *Client, pageID, componentID string) (err error) {
	return deleteResource(client, pageID, "component", componentID)
}
