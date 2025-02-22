package threatintel

type FeedsBundle struct {
	Version      string       `json:"version"`
	CreatedAt    int64        `json:"created_at"`
	ScannerFeeds ScannerFeeds `json:"scanner_feeds"`
	TracerFeeds  TracerFeeds  `json:"tracer_feeds"`
	Extra        []string     `json:"extra"`
}

type ScannerFeeds struct {
	VulnerabilityRules   []KhulnasoftRule `json:"vulnerability_rules"`
	SecretRules          []KhulnasoftRule `json:"secret_rules"`
	MalwareRules         []KhulnasoftRule `json:"malware_rules"`
	ComplianceRules      []KhulnasoftRule `json:"compliance_rules"`
	CloudComplianceRules []KhulnasoftRule `json:"cloud_compliance_rules"`
}

type TracerFeeds struct {
	NetworkRules      []KhulnasoftRule `json:"network_rules"`
	FilesystemRules   []KhulnasoftRule `json:"filesystem_rules"`
	ProcessRules      []KhulnasoftRule `json:"process_rules"`
	ExternalArtefacts []Artefact       `json:"external_artefacts"`
}

type Artefact struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content []byte `json:"content"`
}

type KhulnasoftRule struct {
	RuleID      string `json:"rule_id"`
	Type        string `json:"type"`
	Payload     string `json:"payload"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
}
