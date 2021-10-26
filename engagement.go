package dodjo

import (
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Engagement struct {
	*Client
	Id                         int       `json:"id,omitempty"`
	Tags                       []string  `json:"tags,omitempty"`
	Name                       string    `json:"name"`
	Description                string    `json:"description,omitempty"`
	Version                    string    `json:"version,omitempty"`
	FirstContacted             string    `json:"first_contacted,omitempty"`
	TargetStart                string    `json:"target_start"`
	TargetEnd                  string    `json:"target_end"`
	Reason                     string    `json:"reason,omitempty"`
	Updated                    time.Time `json:"updated,omitempty"`
	Created                    time.Time `json:"created"`
	Active                     bool      `json:"active,omitempty"`
	Tracker                    string    `json:"tracker,omitempty"`
	TestStrategy               string    `json:"test_strategy,omitempty"`
	ThreatModel                bool      `json:"threat_model,omitempty"`
	ApiTest                    bool      `json:"api_test,omitempty"`
	PenTest                    bool      `json:"pen_test,omitempty"`
	CheckList                  bool      `json:"check_list,omitempty"`
	Status                     string    `json:"status,omitempty"`
	Progress                   string    `json:"progress,omitempty"`
	TmodelPath                 string    `json:"tmodel_path,omitempty"`
	DoneTesting                bool      `json:"done_testing,omitempty"`
	EngagementType             string    `json:"engagement_type,omitempty"`
	BuildId                    string    `json:"build_id,omitempty"`
	CommitHash                 string    `json:"commit_hash,omitempty"`
	BranchTag                  string    `json:"branch_tag,omitempty"`
	SourceCodeManagementUri    string    `json:"source_code_management_uri,omitempty"`
	DeduplicationOnEngagement  bool      `json:"deduplication_on_engagement"`
	Lead                       int       `json:"lead,omitempty"`
	Requester                  string    `json:"requester,omitempty"`
	Preset                     string    `json:"preset,omitempty"`
	ReportType                 string    `json:"report_type,omitempty"`
	Product                    int       `json:"product"`
	BuildServer                string    `json:"build_server,omitempty"`
	SourceCodeManagementServer string    `json:"source_code_management_server,omitempty"`
	OrchestrationEngine        string    `json:"orchestration_engine,omitempty"`
	Notes                      []string  `json:"notes,omitempty"`
	Files                      []string  `json:"files,omitempty"`
	RiskAcceptance             []string  `json:"risk_acceptance,omitempty"`
}

func (engagement *Engagement) ImportSecretDetectionReport(filename string) error {
	mFormDataVars := make(map[string]string, 0)
	mFormDataVars["scan_type"] = "GitLab Secret Detection Report"
	mFormDataVars["engagement"] = strconv.FormatInt(int64(engagement.Id), 10)
	mFormDataVars["file"] = filename
	pFileImportInfo := &FileImport{}
	err := engagement.makeRequest("POST", "/import-scan/", []byte{}, true, mFormDataVars, pFileImportInfo)
	if err != nil {
		logrus.Panic(err)
	}
	if pFileImportInfo.Engagement == engagement.Id {
		return nil
	} else {
		return errors.New("cannot import file")
	}
}
