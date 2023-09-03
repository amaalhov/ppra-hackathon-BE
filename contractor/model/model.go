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
	ContactType                  string    `json:"contact_type,omitempty"`
	FirstName                    string    `json:"first_name,omitempty"`
	MiddleName                   string    `json:"middle_name,omitempty"`
	Fullname                     string    `json:"full_name,omitempty"`
	LastName                     string    `json:"last_name,omitempty"`
	GovernmentIdentification     string    `json:"government_identification,omitempty"`
	GovernmentIdentificationType string    `json:"gid_type,omitempty"`
	DateOfBirth                  time.Time `json:"date_of_birth,omitempty"`
	AppointmentDate              time.Time `json:"appointment_date,omitempty"`
	Email                        string    `json:"email,omitempty"`
	Nationality                  string    `json:"nationality,omitempty"`
	CellNumber                   string    `json:"cell_number,omitempty"`
	Telephone                    string    `json:"telephone,omitempty"`
	BusinessPhoneNumber          string    `json:"business_phone_number,omitempty"`
	Country                      string    `json:"country,omitempty"`
	DistrictName                 string    `json:"disctrict_name,omitempty"`
	Town                         string    `json:"town,omitempty"`
	Street                       string    `json:"street,omitempty"`
	BoxAddress                   string    `json:"box_address,omitempty"`
	PlotNumber                   string    `json:"plot_number,omitempty"`
	PhysicalAddress              string    `json:"physical_address,omitempty"`
}

type ContractorVehicleReq struct {
	ValueOfAssets    string            `json:"value_of_assets,omitempty"`
	ValueOfEquipment string            `json:"value_of_equipment,omitempty"`
	PaidUpCapital    string            `json:"paid_up_capital,omitempty"`
	Vehicles         []Vehicle         `json:"vehicles,omitempty"`
	Plants           []Plant           `json:"plants,omitempty"`
	Properties       []Property        `json:"property,omitempty"`
	OfficeEquipments []OfficeEquipment `json:"office_equipments,omitempty"`
}

type Vehicle struct {
	RegisteredOwner    string    `json:"registered_owner,omitempty"`
	Ownership          string    `json:"ownership,omitempty"`
	RegistrationNumber string    `json:"registration_number,omitempty"`
	DateOfRegistration time.Time `json:"date_of_registration,omitempty"`
	VehicleModel       string    `json:"vehicle_model,omitempty"`
}

type Plant struct {
	RegisteredOwner    string    `json:"registered_owner,omitempty"`
	Ownership          string    `json:"ownership,omitempty"`
	RegistrationNumber string    `json:"registration_number,omitempty"`
	Description        string    `json:"description,omitempty"`
	DateOfPurchase     time.Time `json:"date_of_registration,omitempty"`
}

type Property struct {
	Ownership     string  `json:"ownership,omitempty"`
	PresentValue  float64 `json:"present_value,omitempty"`
	AttachmentUrl string  `json:"attachment_url,omitempty"`
	Locality      string  `json:"locality,omitempty"`
}

type OfficeEquipment struct {
	OfficeEquipment string  `json:"office_equipment,omitempty"`
	PresentValue    float64 `json:"present_value,omitempty"`
	AttachmentUrl   string  `json:"attachment_url,omitempty"`
}

type ContractorEmployeeReq struct {
	TotalNumberOfBotswanaCitizens    int64      `json:"botswana_employees,omitempty"`
	TotalNumberOfNonBotswanaCitizens int64      `json:"non_botswana_employees,omitempty"`
	TotalEmployees                   int64      `json:"total_employees,omitempty"`
	Employees                        []Employee `json:"employees,omitempty"`
}

type Employee struct {
	Firstname   string    `json:"first_name,omitempty"`
	Middlename  string    `json:"middle_name,omitempty"`
	Lastname    string    `json:"last_name,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
}

type ContractorSecretaryReq struct {
	Fullname        string    `json:"fullname,omitempty"`
	Nationality     string    `json:"nationality,omitempty"`
	BoxAddress      string    `json:"box_address,omitempty"`
	PhysicalAddress string    `json:"physical_address,omitempty"`
	AppointmentDate time.Time `json:"appointment_date,omitempty"`
}

type ContractorShareHolderReq struct {
	Fullname        string    `json:"fullname,omitempty"`
	Nationality     string    `json:"nationality,omitempty"`
	BoxAddress      string    `json:"box_address,omitempty"`
	PhysicalAddress string    `json:"physical_address,omitempty"`
	AppointmentDate time.Time `json:"appointment_date,omitempty"`
}

type ContractorDirectorReq struct {
	Fullname        string    `json:"fullname,omitempty"`
	Nationality     string    `json:"nationality,omitempty"`
	BoxAddress      string    `json:"box_address,omitempty"`
	PhysicalAddress string    `json:"physical_address,omitempty"`
	AppointmentDate time.Time `json:"appointment_date,omitempty"`
}

type ContractorBankDetailsReq struct {
	BankName      string `json:"bank_name,omitempty"`
	Branch        string `json:"branch,omitempty"`
	BranchAddress string `json:"branch_address,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	AccountType   string `json:"account_type,omitempty"`
}

type ContractorAffiliateReq struct {
	Fullname      string `json:"full_name,omitempty"`
	Address       string `json:"address,omitempty"`
	AttachmentUlr string `json:"attachment_url,omitempty"`
}

type ContractorProjectReq struct {
	ProjectName          string `json:"project_name,omitempty"`
	Description          string `json:"description,omitempty"`
	ClientName           string `json:"client_name,omitempty"`
	ClientRepresentative string `json:"client_representative,omitempty"`
	ClientContactNumber  string `json:"client_contact_number,omitempty"`
}
