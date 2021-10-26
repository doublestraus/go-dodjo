package dodjo

type FileImport struct {
	ScanDate         string      `json:"scan_date"`
	MinimumSeverity  string      `json:"minimum_severity"`
	Active           bool        `json:"active"`
	Verified         bool        `json:"verified"`
	ScanType         string      `json:"scan_type"`
	EndpointToAdd    interface{} `json:"endpoint_to_add,omitempty"`
	File             interface{} `json:"file,omitempty"`
	Engagement       int         `json:"engagement"`
	Lead             interface{} `json:"lead,omitempty"`
	CloseOldFindings bool        `json:"close_old_findings"`
	PushToJira       bool        `json:"push_to_jira"`
	SonarqubeConfig  interface{} `json:"sonarqube_config,omitempty"`
	Test             int         `json:"test"`
}
