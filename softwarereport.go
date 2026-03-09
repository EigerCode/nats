package nats

import "time"

// SoftwareReport is sent by the OpenUEM agent after reading local
// Munki (macOS) or CIMIAN (Windows) report files.
type SoftwareReport struct {
	AgentID   string                  `json:"agent_id"`
	Platform  string                  `json:"platform"` // "darwin" or "windows"
	Timestamp time.Time               `json:"timestamp"`
	Installed []SoftwareReportItem    `json:"installed,omitempty"`
	Results   []SoftwareInstallResult `json:"results,omitempty"`
	Pending   []SoftwareReportItem    `json:"pending,omitempty"`
	Errors    []string                `json:"errors,omitempty"`
	Warnings  []string                `json:"warnings,omitempty"`
}

// SoftwareReportItem represents a single managed software item.
type SoftwareReportItem struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	InstalledAt string `json:"installed_at,omitempty"`
}

// SoftwareInstallResult represents the result of a single install/update/uninstall action.
type SoftwareInstallResult struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Action  string `json:"action"` // "install", "update", "uninstall"
	Status  string `json:"status"` // "success", "failed"
	Time    string `json:"time"`
	Error   string `json:"error,omitempty"`
}
