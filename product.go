package dodjo

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type Product struct {
	*Client
	Id            int      `json:"id,omitempty"`
	FindingsCount int      `json:"findings_count,omitempty"`
	FindingsList  []int    `json:"findings_list,omitempty"`
	Tags          []string `json:"tags"`
	ProductMeta   []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"product_meta,omitempty"`
	Name                       string    `json:"name"`
	Description                string    `json:"description"`
	Created                    time.Time `json:"created"`
	ProdNumericGrade           int       `json:"prod_numeric_grade"`
	BusinessCriticality        string    `json:"business_criticality"`
	Platform                   string    `json:"platform"`
	Lifecycle                  string    `json:"lifecycle"`
	Origin                     string    `json:"origin"`
	UserRecords                int64     `json:"user_records"`
	Revenue                    string    `json:"revenue"`
	ExternalAudience           bool      `json:"external_audience"`
	InternetAccessible         bool      `json:"internet_accessible"`
	EnableSimpleRiskAcceptance bool      `json:"enable_simple_risk_acceptance"`
	EnableFullRiskAcceptance   bool      `json:"enable_full_risk_acceptance"`
	ProductManager             int       `json:"product_manager,omitempty"`
	TechnicalContact           int       `json:"technical_contact,omitempty"`
	TeamManager                int       `json:"team_manager,omitempty"`
	ProdType                   int       `json:"prod_type"`
	Members                    []int     `json:"members,omitempty"`
	AuthorizationGroups        []int     `json:"authorization_groups,omitempty"`
	Regulations                []int     `json:"regulations,omitempty"`
}

func (product *Product) GetEngagements() []*Engagement {
	engagements := &Engagements{}
	err := product.makeRequest("GET", fmt.Sprintf("/engagements/?product=%d", product.Id), []byte{}, false, nil, engagements)
	if err != nil {
		logrus.Fatal(err)
	}
	engagementsData := make([]*Engagement, 0)
	for _, eng := range engagements.Results {
		eng.Client = product.Client
		engagementsData = append(engagementsData, &eng)
	}
	return engagementsData
}

func (product *Product) AddEngagement(name, description string, startTime, endTime *time.Time) (*Engagement, error) {
	engagement := &Engagement{
		DeduplicationOnEngagement: false,
		Name:                      name,
		Description:               description,
		Created:                   time.Now(),
		Product:                   product.Id,
		Active:                    true,
		Status:                    "In Progress",
		TargetStart:               startTime.Format("2006-01-02"),
		TargetEnd:                 endTime.Format("2006-01-02"),
	}
	body, err := json.Marshal(engagement)
	if err != nil {
		return nil, err
	}
	err = product.makeRequest("POST", "/engagements/", body, false, nil, engagement)
	if err != nil {
		return nil, err
	}
	engagement.Client = product.Client
	return engagement, nil
}
