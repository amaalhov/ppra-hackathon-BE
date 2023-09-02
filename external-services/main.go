package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

type GeneralDetails struct {
	UIN                       string `json:"uid"`
	CompanyStatus             string `json:"company_status"`
	ForeignCompany            string `json:"foreign_company"`
	Exempt                    string `json:"exempt"`
	IncorporationDate         string `json:"incorporation_date"`
	ReRegistrationDate        string `json:"re_registration_date"`
	HasOwnConstitution        bool   `json:"has_own_constitution"`
	ConstitutionDocUrl        string `json:"constitution_url"`
	TranslationOfConstitution string `json:"translation_of_constitution"`
	AnnualReturnFilingMonth   string `json:"annual_return_filing_month"`
	AnnualReturnLastFiledOn   string `json:"annual_return_last_filed_on"`
}

type Addresses struct {
	RegisteredOfficeAddress string `json:"registered_office_address,omitempty"`
	PostalAddress           string `json:"postal_address,omitempty"`
	PrincipalPlaceOfBusines string `json:"principal_place_of_business,omitempty"`
	ResidentialAddress      string `json:"residential_address,omitempty,omitempty"`
}

type KeyPersonnel struct {
	Fullname           string `json:"fullname"`
	Nationality        string `json:"nationality"`
	ResidentialAddress string `json:"residential_address"`
	PostalAddress      string `json:"postal_address"`
	AppointmentDate    string `json:"appointment_date"`
	IsActive           bool   `json:"is_active"`
	CeasedDate         string `json:"ceased_date,omitempty"`
}

type Auditor struct {
	Name               string `json:"auditor_name"`
	Nationality        string `json:"nationality"`
	ResidentialAddress string `json:"residential_address,omitempty"`
	PostalAddress      string `json:"postal_address,omitempty"`
	AppointmentDate    string `json:"appointment_date"`
}

type ShareHolder struct {
	Name                                 string    `json:"shareholder_name"`
	Addresses                            Addresses `json:"addresses"`
	RegistrationNumber                   string    `json:"registration_number,omitempty"`
	CountryOfRegistration                string    `json:"country_of_registration"`
	CertificateOfIncorporationDocUrl     string    `json:"certificate_of_incorporation_doc_url,omitempty"`
	TranslationOfEvidenceOfIncorporation string    `json:"translation_of_evidence_of_incorporation,omitempty"`
	IsNomineeShareholder                 bool      `json:"is_nominee_shareholder"`
	AppointmentDate                      string    `json:"appointment_date"`
}

type Share struct {
	NumberOfShares  int64  `json:"number_of_shares"`
	ShareholderName string `json:"shareholder_name"`
}

type ShareAllocations struct {
	TotalNumberOfShares int64   `json:"total_number_of_shares"`
	Shares              []Share `json:"shares"`
}

type Company struct {
	Name             string           `json:"company_name"`
	GeneralDetails   GeneralDetails   `json:"general_details"`
	Addresses        Addresses        `json:"addresses"`
	Directors        []KeyPersonnel   `json:"directors"`
	Secretaries      []KeyPersonnel   `json:"secretaries"`
	Auditors         []Auditor        `json:"auditors"`
	ShareHolders     []ShareHolder    `json:"shareholders"`
	ShareAllocations ShareAllocations `json:"share_allocations"`
}

func main() {
	router := bunrouter.New()

	router.POST("/api/external_services/cipa", func(w http.ResponseWriter, req bunrouter.Request) error {
		// req embeds *http.Request and has all the same fields and methods
		type reqData struct {
			Fullname string `json:"representative_name"`
		}

		var r reqData

		json.NewDecoder(req.Body).Decode(&r)

		uin := req.URL.Query().Get("cipa_uin")
		if uin == "" {
			w.WriteHeader(http.StatusBadRequest)
			return bunrouter.JSON(w, bunrouter.H{
				"message": "cipa identifier required",
				"status":  false,
			})
		}

		log.Println(req.Method, req.Route(), req.Param("cipa_uin"))

		company := Company{
			Name: "Sefalana Holding Company Limited",
			GeneralDetails: GeneralDetails{
				UIN:                       "BW00001731678",
				CompanyStatus:             "registered",
				ForeignCompany:            "not specified",
				Exempt:                    "not specified",
				IncorporationDate:         "17 December 1986",
				ReRegistrationDate:        "19 May 2020",
				HasOwnConstitution:        true,
				ConstitutionDocUrl:        "https://www.cipa.co.bw/ng-cipa-companies/document/XP-zGFLGghDXol337Ytop00p3GxP__G5-IFh57jU9BZv6Mhd0F-94nHtFlHgjstlQcaOudj2QkyO7eQjc-3ZJI6qQPArtDkr0injfLl_Ay_7bE-SS?nodeId=a97c7575f1198f7e",
				TranslationOfConstitution: "",
				AnnualReturnFilingMonth:   "November",
				AnnualReturnLastFiledOn:   "14 November 2022",
			},
			Addresses: Addresses{
				RegisteredOfficeAddress: "Plot 10038,Cnr Nelson Mandela Drive & Kubu Rd,Broadhurst Industrial,Gaborone,Botswana",
				PostalAddress:           "Private Bag 0080,Gaborone,Botswana",
				PrincipalPlaceOfBusines: "Plot 10038,Cnr Nelson Mandela Drive & Kubu Rd,Broadhurst Industrial,Gaborone,Botswana",
			},
			Directors: []KeyPersonnel{
				{
					Fullname:           "Chandrakant Chauhan",
					Nationality:        "Botswana",
					ResidentialAddress: "Plot 53698,Phakalane Golf Estate,Gaborone,Botswana",
					PostalAddress:      "Private Bag 0080,Gaborone,Botswana",
					AppointmentDate:    "30 August 2002",
					IsActive:           true,
				},
				{
					Fullname:           "Mohamed Sajid Osman",
					Nationality:        "Botswana",
					ResidentialAddress: "Plot 21147,Village,Gaborone,Botswana",
					PostalAddress:      "Private Bag 0800,Gaborone,Botswana",
					AppointmentDate:    "23 January 2014",
					IsActive:           true,
				},
				{
					Fullname:           "Keneilwe Patricia Mere",
					Nationality:        "Botswana",
					ResidentialAddress: "Plot 42898,Phakalane,Gaborone,Botswana",
					PostalAddress:      "P O Box 46271,Gaborone,Botswana",
					AppointmentDate:    "25 January 2017",
					IsActive:           true,
				},
				{
					Fullname:           "Moatlhodi Sebabole",
					Nationality:        "Botswana",
					ResidentialAddress: "Plot 71215, Phakalane, Gaborone, Botswana",
					PostalAddress:      "P O Box 50331, Gaborone, Botswana",
					IsActive:           false,
					AppointmentDate:    "25 January 2008",
					CeasedDate:         "30 October 2020",
				},
			},
			Secretaries: []KeyPersonnel{
				{
					Fullname:           "Joanne Sylvia Friginal Robinson",
					Nationality:        "United Kingdom",
					ResidentialAddress: "Plot 7861, Maru-A-Pula, Gaborone, Botswana",
					PostalAddress:      "Private Bag 0080, Gaborone, Botswana",
					IsActive:           true,
					AppointmentDate:    "06 April 2022",
				},
				{
					Fullname:           "Gofaone M Senna",
					Nationality:        "Botswana",
					ResidentialAddress: "Plot 961/2, Block 1, Mmopane, Botswana",
					PostalAddress:      "Private Bag 0800, Gaborone, Botswana",
					IsActive:           false,
					AppointmentDate:    "01 February 2020",
					CeasedDate:         "31 January 2022",
				},
			},
			Auditors: []Auditor{
				{
					Name:               "Magritha Juanita Wotherspoon (Deloitte)",
					Nationality:        "South Africa",
					ResidentialAddress: "Plot 301,Extension 5,Gaborone,Botswana",
					AppointmentDate:    "23 October 2018",
				},
			},
			ShareHolders: []ShareHolder{
				{
					Name: "Fnb Botswana Nominees (Pty) Ltd Re:Ag Bpopf Equity",
					Addresses: Addresses{
						RegisteredOfficeAddress: "Plot 54362,Central Business District,Gaborone,Botswana",
						PostalAddress:           "P O Box 1552,Gaborone,Botswana",
					},
					CountryOfRegistration:            "Botswana",
					CertificateOfIncorporationDocUrl: "https://www.cipa.co.bw/ng-cipa-companies/document/XP-zGFLGghDXol337Ytop00p3GxP__G5-IFh57jU9BZv6Mhd0F-94nHtFlHgjstlQcaOudj2QkyO7eQjc-3ZJI6qQPArtDkr0injfLl_Ay_7bE-SS?nodeId=741cc6279c90ec1b",
					IsNomineeShareholder:             true,
					AppointmentDate:                  "15 August 2014",
				},
				{
					Name: "Motor Vehicle Accident Fund",
					Addresses: Addresses{
						RegisteredOfficeAddress: "Plot 50367,Fairgrounds Office Park,Gabane,Botswana",
						PostalAddress:           "Private Bag 00438,Gaborone,Botswana",
					},
					CountryOfRegistration:            "Botswana",
					CertificateOfIncorporationDocUrl: "https://www.cipa.co.bw/ng-cipa-companies/document/XP-zGFLGghDXol337Ytop00p3GxP__G5-IFh57jU9BZv6Mhd0F-94nHtFlHgjstlQcaOudj2QkyO7eQjc-3ZJI6qQPArtDkr0injfLl_Ay_7bE-SS?nodeId=b5a1c0ead2f17580",
					IsNomineeShareholder:             true,
					AppointmentDate:                  "14 August 2014",
				},
				{
					Name: "Stanbic Nominees Botswana Re Bifm Plef",
					Addresses: Addresses{
						RegisteredOfficeAddress: "Plot 50672, Fairgrounds Office Park, Gaborone, Botswana",
						PostalAddress:           "Private Bag 0800, Gaborone, Botswana",
					},
					CountryOfRegistration:            "Botswana",
					CertificateOfIncorporationDocUrl: "https://www.cipa.co.bw/ng-cipa-companies/document/XP-zGFLGghDXol337Ytop00p3GxP__G5-IFh57jU9BZv6Mhd0F-94nHtFlHgjstlQcaOudj2QkyO7eQjc-3ZJI6qQPArtDkr0injfLl_Ay_7bE-SS?nodeId=1680bcd0caf9cdac",
					IsNomineeShareholder:             true,
					AppointmentDate:                  "14 August 2014",
				},
				{
					Name: "Mootiemang Reginald Motswaiso",
					Addresses: Addresses{
						ResidentialAddress: "House 21399, Phakalane, Gaborone, Botswana",
						PostalAddress:      "Private Bag 0800, Gaborone, Botswana",
					},
					CountryOfRegistration: "Botswana",
					IsNomineeShareholder:  false,
					AppointmentDate:       "14 August 2014",
				},
			},
			ShareAllocations: ShareAllocations{
				TotalNumberOfShares: 177127430,
				Shares: []Share{
					{
						NumberOfShares:  31046939,
						ShareholderName: "Fnb Botswana Nominees (Pty) Ltd Re:Ag Bpopf Equity",
					},
					{
						NumberOfShares:  27339765,
						ShareholderName: "Fnb Botswana Nominees (Pty) Ltd Re:Aa Bpopf Equity",
					},
					{
						NumberOfShares:  25083138,
						ShareholderName: "Motor Vehicle Accident Fund",
					},
					{
						NumberOfShares:  22894296,
						ShareholderName: "Fnb Botswana Nominees (Pty) Ltd Re:Bifm Bpopf-Equity",
					},
					{
						NumberOfShares:  14637532,
						ShareholderName: "Fnb Botswana Nominees (Pty) Ltd Re:Iam Bpopf Equity",
					},
					{
						NumberOfShares:  13373679,
						ShareholderName: "Fnbbn (Pty) Ltd Re: Ag Bpopf Equity Portfolio B",
					},
					{
						NumberOfShares:  8088948,
						ShareholderName: "Stanbic Nominees Botswana Re Allan Gray Debswana Pf",
					},
					{
						NumberOfShares:  7230321,
						ShareholderName: "Stanbic Nominees Botswana Re Bifm Plef",
					},
					{
						NumberOfShares:  6996588,
						ShareholderName: "Stanbic Nominees Botswana Re Bifm Mlf",
					},
					{
						NumberOfShares:  6069644,
						ShareholderName: "Sbb Nominees (Pty) Ltd Re: A/C C00021",
					},
					{
						NumberOfShares:  13860448,
						ShareholderName: "Chandrakant Chauhan",
					},
					{
						NumberOfShares:  247620,
						ShareholderName: "Mootiemang Reginald Motswaiso",
					},
					{
						NumberOfShares:  223708,
						ShareholderName: "Mohamed Sajid Osman",
					},
					{
						NumberOfShares:  34804,
						ShareholderName: "Keith Robert Jefferis",
					},
				},
			},
		}

		return bunrouter.JSON(w, company)
	})

	router.GET("/api/external_services/NIR", func(w http.ResponseWriter, req bunrouter.Request) error {
		id := req.URL.Query().Get("national_id_number")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return bunrouter.JSON(w, bunrouter.H{
				"message": "national id number is required",
				"status":  false,
			})
		}

		if id == "123456789" {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			return bunrouter.JSON(w, bunrouter.H{
				"id_number":             "123456789",
				"surname":               "Doe",
				"forenames":             "Jane",
				"date_of_birth":         "10/01/2000",
				"place_of_birth":        "Gaborone",
				"signature":             "",
				"nationality":           "Motswana",
				"sex":                   "Female",
				"color_of_eyes":         "Blue",
				"date_of_expiry":        "14/12/2030",
				"place_of_application":  "Borakanelo",
				"signature_of_register": "",
			})
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "national id not found on national registry",
			"status":  false,
		})
	})

	log.Println("external services listening on http://localhost:3001")
	log.Println(http.ListenAndServe(":3001", router))
}
