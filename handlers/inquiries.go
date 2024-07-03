package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"synthtic-harbor/models"
)

var insuranceProviders = []string{"Aetna", "Cigna", "Humana", "MetLife", "United Healthcare", "Delta Dental", "Blue Cross Blue Shield"}

func GetDentalInquiry(c *gin.Context) {
	id := c.Param("id")
	inquiry := generateFakeInquiry(id)
	c.JSON(http.StatusOK, inquiry)
}

func GetDentalInquiries(c *gin.Context) {
	limit := 10 // Default limit
	inquiries := []models.Inquiry{}
	for i := 0; i < limit; i++ {
		id := uuid.New().String()
		inquiry := generateFakeInquiry(id)
		inquiries = append(inquiries, inquiry)
	}
	c.JSON(http.StatusOK, inquiries)
}

func generateFakeInquiry(id string) models.Inquiry {
	return models.Inquiry{
		ID:         id,
		Status:     randomStatus(),
		CreationTS: time.Now(), // Use current time for simplicity
		Summary:    faker.Sentence(),
		Results: models.Results{
			CallDetails: []models.CallDetail{
				{
					CallEndTime:        time.Now(),
					RepresentativeName: faker.Name(),
					ReferenceNumber:    faker.UUIDDigit(),
					CallRecordingURL:   faker.URL(),
				},
			},
			PlanInformation: models.PlanInformation{
				IsActive:            randomBool(),
				PlanType:            randomPlanType(),
				EffectiveDate:       faker.Date(),
				TerminationDate:     faker.Date(),
				IsCalendarYearPlan:  randomBool(),
				IsProviderInNetwork: randomBool(),
				IsPrimaryInsurance:  randomBool(),
			},
			Maximums: models.Maximums{
				IndividualDeductible:      strconv.Itoa(rand.Intn(999)),
				IndividualDeductibleUsed:  strconv.Itoa(rand.Intn(999)),
				FamilyDeductible:          strconv.Itoa(rand.Intn(999)),
				FamilyDeductibleUsed:      strconv.Itoa(rand.Intn(999)),
				CoverageLimit:             strconv.Itoa(rand.Intn(9999)),
				CoverageLimitRemaining:    strconv.Itoa(rand.Intn(9999)),
				OrthodonticsCoverageLimit: strconv.Itoa(rand.Intn(9999)),
			},
			TreatmentHistory: map[string][]models.TreatmentHistory{
				faker.Date(): {
					{
						ProcedureCode:   "D1110",
						ToothNumbers:    []int{rand.Intn(32) + 1},
						Surfaces:        []string{"occlusal"},
						QuadrantNumbers: []int{rand.Intn(4) + 1},
					},
					{
						ProcedureCode:   "D1120",
						ToothNumbers:    []int{rand.Intn(32) + 1},
						Surfaces:        []string{"distal"},
						QuadrantNumbers: []int{rand.Intn(4) + 1},
					},
				},
			},
			ProcedureClasses: map[string]models.ProcedureClass{
				"DIAGNOSTIC": {
					IsCovered:          randomBool(),
					CoveragePercentage: rand.Intn(100),
				},
				"IMPLANTS": {
					IsCovered: randomBool(),
				},
			},
			ProcedureCodes: [][]models.ProcedureCode{
				{
					{
						ProcedureCode:        "D1110",
						ProcedureName:        "Prophylaxis",
						IsCovered:            randomBool(),
						IsPriorAuthRequired:  randomBool(),
						CoveragePercentage:   rand.Intn(100),
						FrequencyLimitations: "1 every 6 floating months",
						MoreInfo:             faker.Sentence(),
					},
				},
			},
			ExtraInfo: models.ExtraInfo{
				WaitingPeriods:                faker.Sentence(),
				DeductibleAppliesToPreventive: randomBool(),
				PreventativeAppliesToMaximum:  randomBool(),
				MissingToothClause:            faker.Sentence(),
				HasPendingClaims:              randomBool(),
			},
		},
		Request: models.Request{
			Type:                  "BENEFITS",
			DesiredCompletionDate: faker.Date(),
			PatientName:           faker.Name(),
			DOB:                   faker.Date(),
			MemberID:              faker.UUIDDigit(),
			GroupID:               faker.UUIDDigit(),
			InsuranceInNetwork:    randomBool(),
			NPI:                   faker.UUIDDigit(),
			TaxID:                 faker.UUIDDigit(),
			ExternalID:            faker.UUIDDigit(),
			DiagnosisCodes:        []string{"F41.1", "F42.23"},
			ClaimsDateOfService:   faker.Date(),
			ClaimNumber:           faker.UUIDDigit(),
			Insurance:             randomInsuranceProvider(),
			BenefitsQuery:         []string{"CODE_LOOKUP_BENEFITS"},
			BenefitsCodes:         []string{"D1110", "D1120", "1130"},
			IsSpecialist:          randomBool(),
		},
	}
}

func randomStatus() string {
	statuses := []string{"SCHEDULED", "IN_PROGRESS", "SUCCESS", "UNSUCCESSFUL", "UNKNOWN"}
	return statuses[rand.Intn(len(statuses))]
}

func randomPlanType() string {
	planTypes := []string{"PPO", "HMO", "EPO", "POS", "Indemnity"}
	return planTypes[rand.Intn(len(planTypes))]
}

func randomInsuranceProvider() string {
	return insuranceProviders[rand.Intn(len(insuranceProviders))]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}
