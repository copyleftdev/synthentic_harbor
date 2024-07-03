package models

import "time"

type Inquiry struct {
	ID         string    `json:"id"`
	Status     string    `json:"status"`
	CreationTS time.Time `json:"creation_ts"`
	Summary    string    `json:"summary"`
	Results    Results   `json:"results"`
	Request    Request   `json:"request"`
}

type Results struct {
	CallDetails      []CallDetail                  `json:"call_details"`
	PlanInformation  PlanInformation               `json:"plan_information"`
	Maximums         Maximums                      `json:"maximums"`
	TreatmentHistory map[string][]TreatmentHistory `json:"treatment_history"`
	ProcedureClasses map[string]ProcedureClass     `json:"procedure_classes"`
	ProcedureCodes   [][]ProcedureCode             `json:"procedure_codes"`
	ExtraInfo        ExtraInfo                     `json:"extra_info"`
}

type CallDetail struct {
	CallEndTime        time.Time `json:"call_end_time"`
	RepresentativeName string    `json:"representative_name"`
	ReferenceNumber    string    `json:"reference_number"`
	CallRecordingURL   string    `json:"call_recording_url"`
}

type PlanInformation struct {
	IsActive            bool   `json:"is_active"`
	PlanType            string `json:"plan_type"`
	EffectiveDate       string `json:"effective_date"`
	TerminationDate     string `json:"termination_date"`
	IsCalendarYearPlan  bool   `json:"is_calendar_year_plan"`
	IsProviderInNetwork bool   `json:"is_provider_in_network"`
	IsPrimaryInsurance  bool   `json:"is_primary_insurance"`
}

type Maximums struct {
	IndividualDeductible      string `json:"individual_deductible"`
	IndividualDeductibleUsed  string `json:"individual_deductible_used"`
	FamilyDeductible          string `json:"family_deductible"`
	FamilyDeductibleUsed      string `json:"family_deductible_used"`
	CoverageLimit             string `json:"coverage_limit"`
	CoverageLimitRemaining    string `json:"coverage_limit_remaining"`
	OrthodonticsCoverageLimit string `json:"orthodontics_coverage_limit"`
}

type TreatmentHistory struct {
	ProcedureCode   string   `json:"procedure_code"`
	ToothNumbers    []int    `json:"tooth_numbers"`
	Surfaces        []string `json:"surfaces"`
	QuadrantNumbers []int    `json:"quadrant_numbers"`
}

type ProcedureClass struct {
	IsCovered          bool `json:"is_covered"`
	CoveragePercentage int  `json:"coverage_percentage,omitempty"`
}

type ProcedureCode struct {
	ProcedureCode        string `json:"procedure_code"`
	ProcedureName        string `json:"procedure_name"`
	IsCovered            bool   `json:"is_covered"`
	IsPriorAuthRequired  bool   `json:"is_prior_auth_required"`
	CoveragePercentage   int    `json:"coverage_percentage,omitempty"`
	FrequencyLimitations string `json:"frequency_limitations,omitempty"`
	MoreInfo             string `json:"more_info,omitempty"`
}

type ExtraInfo struct {
	WaitingPeriods                string `json:"waiting_periods"`
	DeductibleAppliesToPreventive bool   `json:"deductible_applies_to_preventive"`
	PreventativeAppliesToMaximum  bool   `json:"preventative_applies_to_maximum"`
	MissingToothClause            string `json:"missing_tooth_clause"`
	HasPendingClaims              bool   `json:"has_pending_claims"`
}

type Request struct {
	Type                  string   `json:"type"`
	DesiredCompletionDate string   `json:"desired_completion_date"`
	PatientName           string   `json:"patient_name"`
	DOB                   string   `json:"dob"`
	MemberID              string   `json:"member_id"`
	GroupID               string   `json:"group_id"`
	InsuranceInNetwork    bool     `json:"insurance_in_network"`
	NPI                   string   `json:"npi"`
	TaxID                 string   `json:"tax_id"`
	ExternalID            string   `json:"external_id"`
	DiagnosisCodes        []string `json:"diagnosis_codes"`
	ClaimsDateOfService   string   `json:"claims_date_of_service"`
	ClaimNumber           string   `json:"claim_number"`
	Insurance             string   `json:"insurance"`
	BenefitsQuery         []string `json:"benefits_query"`
	BenefitsCodes         []string `json:"benefits_codes"`
	IsSpecialist          bool     `json:"is_specialist"`
}
