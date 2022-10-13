package mypackage

type AutoGenerated struct {
	Topics struct {
		HubAdminEvents          []string `json:"hubAdminEvents"`
		HubStatus               []string `json:"hubStatus"`
		ScenarioEventsGen2      []string `json:"scenarioEvents_gen2"`
		HubEvents               []string `json:"hubEvents"`
		DeviceEvents            []string `json:"deviceEvents"`
		BankEvents              []string `json:"bankEvents"`
		NfcEvents               []string `json:"nfcEvents"`
		ScenarioTriggeredEvents []string `json:"scenarioTriggeredEvents"`
		UserEventsGen2          []string `json:"userEvents_gen2"`
		DeviceReplaceEvents     []string `json:"deviceReplaceEvents"`
		DetectionClusterEvents  []string `json:"detectionClusterEvents"`
		CameraStatusEvents      []string `json:"cameraStatusEvents"`
	} `json:"topics"`
}