package hubspot

const (
	// CRMAssociationContactToCompany ...
	CRMAssociationContactToCompany = 1
	// CRMAssociationCompanyToContact ...
	CRMAssociationCompanyToContact = 2
	// CRMAssociationDealToContact ...
	CRMAssociationDealToContact = 3
	// CRMAssociationContactToDeal ...
	CRMAssociationContactToDeal = 4
	// CRMAssociationDealToCompany ...
	CRMAssociationDealToCompany = 5
	// CRMAssociationCompanyToDeal ...
	CRMAssociationCompanyToDeal = 6
	// CRMAssociationCompanyToEngagement ...
	CRMAssociationCompanyToEngagement = 7
	// CRMAssociationEngagementToCompany ...
	CRMAssociationEngagementToCompany = 8
	// CRMAssociationContactToEngagement ...
	CRMAssociationContactToEngagement = 9
	// CRMAssociationEngagementToContact ...
	CRMAssociationEngagementToContact = 10
	// CRMAssociationDealToEngagement ...
	CRMAssociationDealToEngagement = 11
	// CRMAssociationEngagementToDeal ...
	CRMAssociationEngagementToDeal = 12
	// CRMAssociationParentCompanyToChildCompany ...
	CRMAssociationParentCompanyToChildCompany = 13
	// CRMAssociationChildCompanyToParentCompany ...
	CRMAssociationChildCompanyToParentCompany = 14
	// CRMAssociationContactToTicket ...
	CRMAssociationContactToTicket = 15
	// CRMAssociationTicketToContact ...
	CRMAssociationTicketToContact = 16
	// CRMAssociationTicketToEngagement ...
	CRMAssociationTicketToEngagement = 17
	// CRMAssociationEngagementToTicket ...
	CRMAssociationEngagementToTicket = 18
	// CRMAssociationDealToLineItem ...
	CRMAssociationDealToLineItem = 19
	// CRMAssociationLineItemToDeal ...
	CRMAssociationLineItemToDeal = 20
	// CRMAssociationCompanyToTicket ...
	CRMAssociationCompanyToTicket = 25
	// CRMAssociationTicketToCompany ...
	CRMAssociationTicketToCompany = 26
	// CRMAssociationDealToTicket ...
	CRMAssociationDealToTicket = 27
	// CRMAssociationTicketToDeal ...
	CRMAssociationTicketToDeal = 28
	// CRMAssociationAdvisorToCompany ...
	CRMAssociationAdvisorToCompany = 33
	// CRMAssociationCompanyToAdvisor ...
	CRMAssociationCompanyToAdvisor = 34
	// CRMAssociationBoardMemberToCompany ...
	CRMAssociationBoardMemberToCompany = 35
	// CRMAssociationCompanyToBoardMember ...
	CRMAssociationCompanyToBoardMember = 36
	// CRMAssociationContractorToCompany ...
	CRMAssociationContractorToCompany = 37
	// CRMAssociationCompanyToContractor ...
	CRMAssociationCompanyToContractor = 38
	// CRMAssociationManagerToCompany ...
	CRMAssociationManagerToCompany = 39
	// CRMAssociationCompanyToManager ...
	CRMAssociationCompanyToManager = 40
	// CRMAssociationBusinessOwnerToCompany ...
	CRMAssociationBusinessOwnerToCompany = 41
	// CRMAssociationCompanyToBusinessOwner ...
	CRMAssociationCompanyToBusinessOwner = 42
	// CRMAssociationPartnerToCompany ...
	CRMAssociationPartnerToCompany = 43
	// CRMAssociationCompanyToPartner ...
	CRMAssociationCompanyToPartner = 44
	// CRMAssociationResellerToCompany ...
	CRMAssociationResellerToCompany = 45
	// CRMAssociationCompanyToReseller ...
	CRMAssociationCompanyToReseller = 46
)

// CRMAssociations API client
type CRMAssociations struct {
	Client
}

// CRMAssociations constructor (from Client)
func (c Client) CRMAssociations() CRMAssociations {
	return CRMAssociations{
		Client: c,
	}
}

// CRMAssociationsRequest object
type CRMAssociationsRequest struct {
	FromObjectID int    `json:"fromObjectId"`
	ToObjectID   int    `json:"toObjectId"`
	Category     string `json:"category"`
	DefinitionID int    `json:"definitionId"`
}

// Create new CRM Association
func (ca CRMAssociations) Create(data CRMAssociationsRequest) error {
	// Validate Category
	if data.Category == "" {
		data.Category = "HUBSPOT_DEFINED"
	}

	return ca.Client.Request("PUT", "/crm-associations/v1/associations", data, nil)
}

// Delete CRM Association
func (ca CRMAssociations) Delete(data CRMAssociationsRequest) error {
	// Validate Category
	if data.Category == "" {
		data.Category = "HUBSPOT_DEFINED"
	}

	return ca.Client.Request("PUT", "/crm-associations/v1/associations/delete", data, nil)
}