package model

import "time"

type ContractorDetailsReq struct {
	NameOfCompany        string    `json:"name_of_company,omitempty"`
	BusinessType         string    `json:"business_type,omitempty"`
	IsLocalContractor    bool      `json:"is_local_contractor,omitempty"`
	CipaUin              string    `json:"cipa_uin,omitempty"`
	NationalIdNumber     string    `json:"omang,omitempty"`
	IsRegisteredWithCIPA bool      `json:"is_registered_with_cipa,omitempty"`
	RegistrationDate     time.Time `json:"registration_date,omitempty"`
	OwnershipType        string    `json:"ownership_type,omitempty"`
}

type ContractorAddressReq struct {
	Country      string `json:"country,omitempty"`
	DistrictName string `json:"district_name,omitempty"`
	Town         string `json:"town,omitempty"`
	PlotNumber   string `json:"plot_number,omitempty"`
	Street       string `json:"street,omitempty"`
}

type ContractorContactReq struct {
	FirstName                    string    `json:"first_name,omitempty"`
	MiddleName                   string    `json:"middle_name,omitempty"`
	Fullname                     string    `json:"full_name,omitempty"`
	LastName                     string    `json:"last_name,omitempty"`
	GovernmentIdentification     string    `json:"government_identification,omitempty"`
	GovernmentIdentificationType string    `json:"gid_type,omitempty"`
	DateOfBirth                  time.Time `json:"date_of_birth,omitempty"`
	Email                        string    `json:"email,omitempty"`
	Nationality                  string    `json:"nationality,omitempty"`
	CellNumber                   string    `json:"cell_number,omitempty"`
	Telephone                    string    `json:"telephone,omitempty"`
	BusinessPhoneNumber          string    `json:"business_phone_number,omitempty"`
}

type PostalAddressReq struct {
	Country        string `json:"country,omitempty"`
	DistrictName   string `json:"disctrict_name,omitempty"`
	Town           string `json:"town,omitempty"`
	BoxAddress     string `json:"box_address,omitempty"`
	AddressNumber  string `json:"address_number,omitempty"`
	LineAddressOne string `json:"line_address_one,omitempty"`
	LineAddressTwo string `json:"line_address_two,omitempty"`
}

type ContractorBankDetailsReq struct {
	BankName      string `json:"bank_name,omitempty"`
	Branch        string `json:"branch,omitempty"`
	BranchAddress string `json:"branch_address,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	AccountType   string `json:"account_type,omitempty"`
}

type ContractorDirectorReq struct {
	FirstName   string    `json:"first_name,omitempty"`
	MiddleName  string    `json:"middle_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
	Gender      string    `json:"gender,omitempty"`
}

type ContractorShareholderReq struct {
	FirstName   string    `json:"first_name,omitempty"`
	MiddleName  string    `json:"middle_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
	Gender      string    `json:"gender,omitempty"`
}

type ContractorAffiliate struct {
	Fullname      string `json:"full_name,omitempty"`
	Address       string `json:"address,omitempty"`
	AttachmentUlr string `json:"attachment_url,omitempty"`
}

type ContractorAffiliateReq struct {
	ContractorAffiliates []ContractorAffiliate `json:"contractor_affiliates"`
}
