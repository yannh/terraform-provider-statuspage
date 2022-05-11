package statuspage

type Incident struct {
	Name           string      `json:"name"`
	Status         string      `json:"status"`
	ImpactOverride string      `json:"impact_override"`
	ScheduledFor   string      `json:"scheduled_for"`
	ScheduledUntil string      `json:"scheduled_until"`
	Body           string      `json:"body"`
	ComponentIDs   []string    `json:"component_ids"`
	Components     []Component `json:"components"`
}

type IncidentFull struct {
	Incident
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CreateIncident(client *Client, pageID string, incident *Incident) (*IncidentFull, error) {
	var i IncidentFull
	err := createResource(
		client,
		pageID,
		"incident",
		struct {
			Incident *Incident `json:"incident"`
		}{incident},
		&i,
	)

	return &i, err
}

func GetIncident(client *Client, pageID, incidentID string) (*IncidentFull, error) {
	var i IncidentFull
	err := readResource(client, pageID, incidentID, "incident", &i)

	return &i, err
}

func UpdateIncident(client *Client, pageID, incidentID string, incident *Incident) (*IncidentFull, error) {
	var i IncidentFull

	err := updateResource(
		client,
		pageID,
		"incident",
		incidentID,
		struct {
			Incident *Incident `json:"incident"`
		}{incident},
		&i,
	)

	return &i, err
}

func DeleteIncident(client *Client, pageID, incidentID string) (err error) {
	return deleteResource(client, pageID, "incident", incidentID)
}
