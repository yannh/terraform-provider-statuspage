package statuspage

type IncidentTemplate struct {
	Name         *string  `json:"name"`
	GroupID      *string  `json:"group_id"`
	UpdateStatus *string  `json:"update_status"`
	Title        *string  `json:"suffix"`
	Body         *string  `json:"y_axis_min"`
	ComponentIDs []string `json:"component_ids"`
	ShouldTweet  *bool    `json:"should_tweet"`
}

type IncidentTemplateFull struct {
	IncidentTemplate
	ID        *string `json:"id"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func CreateIncidentTemplate(client *Client, pageID string, incidentTemplate *IncidentTemplate) (*IncidentTemplateFull, error) {
	var i IncidentTemplateFull
	err := createResource(
		client,
		pageID,
		"incident_template",
		struct {
			IncidentTemplate *IncidentTemplate `json:"incident_template"`
		}{incidentTemplate},
		&i,
	)

	return &i, err
}

func GetIncidentTemplate(client *Client, pageID, incidentTemplateID string) (*IncidentTemplateFull, error) {
	var i IncidentTemplateFull
	err := readResource(client, pageID, incidentTemplateID, "incident_template", &i)

	return &i, err
}

func UpdateIncidentTemplate(client *Client, pageID, incidentTemplateID string, incidentTemplate *IncidentTemplate) (*IncidentTemplateFull, error) {
	var i IncidentTemplateFull

	err := updateResource(
		client,
		pageID,
		"incident_template",
		incidentTemplateID,
		struct {
			IncidentTemplate *IncidentTemplate `json:"incident_template"`
		}{incidentTemplate},
		&i,
	)

	return &i, err
}

func DeleteIncidentTemplate(client *Client, pageID, incidentTemplateID string) (err error) {
	return deleteResource(client, pageID, "incident_template", incidentTemplateID)
}
