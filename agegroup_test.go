package parkrunparser

import (
	"testing"
)

func TestParseAgeGroup(t *testing.T) {
	for _, test := range []struct {
		input          string
		age            string
		sex            Sex
		expected_error bool
	}{
		{"M10", "10", SEX_MALE, false},
		{"M11-14", "11-14", SEX_MALE, false},
		{"M15-17", "15-17", SEX_MALE, false},
		{"M18-19", "18-19", SEX_MALE, false},
		{"M20-24", "20-24", SEX_MALE, false},
		{"M25-29", "25-29", SEX_MALE, false},
		{"M30-34", "30-34", SEX_MALE, false},
		{"M35-39", "35-39", SEX_MALE, false},
		{"M40-44", "40-44", SEX_MALE, false},
		{"M45-49", "45-49", SEX_MALE, false},
		{"M50-54", "50-54", SEX_MALE, false},
		{"M55-59", "55-59", SEX_MALE, false},
		{"M60-64", "60-64", SEX_MALE, false},
		{"M65-69", "65-69", SEX_MALE, false},
		{"M70-74", "70-74", SEX_MALE, false},
		{"M75-79", "75-79", SEX_MALE, false},
		{"M80-84", "80-84", SEX_MALE, false},
		{"M85-89", "85-89", SEX_MALE, false},

		{"JM10", "10", SEX_MALE, false},
		{"JM11-14", "11-14", SEX_MALE, false},
		{"JM15-17", "15-17", SEX_MALE, false},
		{"SM18-19", "18-19", SEX_MALE, false},
		{"SM20-24", "20-24", SEX_MALE, false},
		{"SM25-29", "25-29", SEX_MALE, false},
		{"SM30-34", "30-34", SEX_MALE, false},
		{"VM35-39", "35-39", SEX_MALE, false},
		{"VM40-44", "40-44", SEX_MALE, false},
		{"VM45-49", "45-49", SEX_MALE, false},
		{"VM50-54", "50-54", SEX_MALE, false},
		{"VM55-59", "55-59", SEX_MALE, false},
		{"VM60-64", "60-64", SEX_MALE, false},
		{"VM65-69", "65-69", SEX_MALE, false},
		{"VM70-74", "70-74", SEX_MALE, false},
		{"VM75-79", "75-79", SEX_MALE, false},
		{"VM80-84", "80-84", SEX_MALE, false},
		{"VM85-89", "85-89", SEX_MALE, false},
		{"MWC", "WC", SEX_MALE, false},

		{"JH10", "10", SEX_MALE, false},
		{"JH11-14", "11-14", SEX_MALE, false},
		{"JH15-17", "15-17", SEX_MALE, false},
		{"SH18-19", "18-19", SEX_MALE, false},
		{"SH20-24", "20-24", SEX_MALE, false},
		{"SH25-29", "25-29", SEX_MALE, false},
		{"SH30-34", "30-34", SEX_MALE, false},
		{"VH35-39", "35-39", SEX_MALE, false},
		{"VH40-44", "40-44", SEX_MALE, false},
		{"VH45-49", "45-49", SEX_MALE, false},
		{"VH50-54", "50-54", SEX_MALE, false},
		{"VH55-59", "55-59", SEX_MALE, false},
		{"VH60-64", "60-64", SEX_MALE, false},
		{"VH65-69", "65-69", SEX_MALE, false},
		{"VH70-74", "70-74", SEX_MALE, false},
		{"VH75-79", "75-79", SEX_MALE, false},
		{"VH80-84", "80-84", SEX_MALE, false},
		{"VH85-89", "85-89", SEX_MALE, false},
		{"HWC", "WC", SEX_MALE, false},

		{"K10", "10", SEX_FEMALE, false},
		{"K11-14", "11-14", SEX_FEMALE, false},
		{"K15-17", "15-17", SEX_FEMALE, false},
		{"K18-19", "18-19", SEX_FEMALE, false},
		{"K20-24", "20-24", SEX_FEMALE, false},
		{"K25-29", "25-29", SEX_FEMALE, false},
		{"K30-34", "30-34", SEX_FEMALE, false},
		{"K35-39", "35-39", SEX_FEMALE, false},
		{"K40-44", "40-44", SEX_FEMALE, false},
		{"K45-49", "45-49", SEX_FEMALE, false},
		{"K50-54", "50-54", SEX_FEMALE, false},
		{"K55-59", "55-59", SEX_FEMALE, false},
		{"K60-64", "60-64", SEX_FEMALE, false},
		{"K65-69", "65-69", SEX_FEMALE, false},
		{"K70-74", "70-74", SEX_FEMALE, false},
		{"K75-79", "75-79", SEX_FEMALE, false},
		{"K80-84", "80-84", SEX_FEMALE, false},
		{"K85-89", "85-89", SEX_FEMALE, false},

		{"JF10", "10", SEX_FEMALE, false},
		{"JF11-14", "11-14", SEX_FEMALE, false},
		{"JF15-17", "15-17", SEX_FEMALE, false},
		{"SF18-19", "18-19", SEX_FEMALE, false},
		{"SF20-24", "20-24", SEX_FEMALE, false},
		{"SF25-29", "25-29", SEX_FEMALE, false},
		{"SF30-34", "30-34", SEX_FEMALE, false},
		{"VF35-39", "35-39", SEX_FEMALE, false},
		{"VF40-44", "40-44", SEX_FEMALE, false},
		{"VF45-49", "45-49", SEX_FEMALE, false},
		{"VF50-54", "50-54", SEX_FEMALE, false},
		{"VF55-59", "55-59", SEX_FEMALE, false},
		{"VF60-64", "60-64", SEX_FEMALE, false},
		{"VF65-69", "65-69", SEX_FEMALE, false},
		{"VF70-74", "70-74", SEX_FEMALE, false},
		{"VF75-79", "75-79", SEX_FEMALE, false},
		{"VF80-84", "80-84", SEX_FEMALE, false},
		{"VF85-89", "85-89", SEX_FEMALE, false},
		{"FWC", "WC", SEX_FEMALE, false},

		{"JN10", "10", SEX_FEMALE, false},
		{"JN11-14", "11-14", SEX_FEMALE, false},
		{"JN15-17", "15-17", SEX_FEMALE, false},
		{"SN18-19", "18-19", SEX_FEMALE, false},
		{"SN20-24", "20-24", SEX_FEMALE, false},
		{"SN25-29", "25-29", SEX_FEMALE, false},
		{"SN30-34", "30-34", SEX_FEMALE, false},
		{"VN35-39", "35-39", SEX_FEMALE, false},
		{"VN40-44", "40-44", SEX_FEMALE, false},
		{"VN45-49", "45-49", SEX_FEMALE, false},
		{"VN50-54", "50-54", SEX_FEMALE, false},
		{"VN55-59", "55-59", SEX_FEMALE, false},
		{"VN60-64", "60-64", SEX_FEMALE, false},
		{"VN65-69", "65-69", SEX_FEMALE, false},
		{"VN70-74", "70-74", SEX_FEMALE, false},
		{"VN75-79", "75-79", SEX_FEMALE, false},
		{"VN80-84", "80-84", SEX_FEMALE, false},
		{"VN85-89", "85-89", SEX_FEMALE, false},
		{"NWC", "WC", SEX_FEMALE, false},

		{"JV10", "10", SEX_FEMALE, false},
		{"JV11-14", "11-14", SEX_FEMALE, false},
		{"JV15-17", "15-17", SEX_FEMALE, false},
		{"SV18-19", "18-19", SEX_FEMALE, false},
		{"SV20-24", "20-24", SEX_FEMALE, false},
		{"SV25-29", "25-29", SEX_FEMALE, false},
		{"SV30-34", "30-34", SEX_FEMALE, false},
		{"VV35-39", "35-39", SEX_FEMALE, false},
		{"VV40-44", "40-44", SEX_FEMALE, false},
		{"VV45-49", "45-49", SEX_FEMALE, false},
		{"VV50-54", "50-54", SEX_FEMALE, false},
		{"VV55-59", "55-59", SEX_FEMALE, false},
		{"VV60-64", "60-64", SEX_FEMALE, false},
		{"VV65-69", "65-69", SEX_FEMALE, false},
		{"VV70-74", "70-74", SEX_FEMALE, false},
		{"VV75-79", "75-79", SEX_FEMALE, false},
		{"VV80-84", "80-84", SEX_FEMALE, false},
		{"VV85-89", "85-89", SEX_FEMALE, false},
		{"VWC", "WC", SEX_FEMALE, false},

		{"JW10", "10", SEX_FEMALE, false},
		{"JW11-14", "11-14", SEX_FEMALE, false},
		{"JW15-17", "15-17", SEX_FEMALE, false},
		{"SW18-19", "18-19", SEX_FEMALE, false},
		{"SW20-24", "20-24", SEX_FEMALE, false},
		{"SW25-29", "25-29", SEX_FEMALE, false},
		{"SW30-34", "30-34", SEX_FEMALE, false},
		{"VW35-39", "35-39", SEX_FEMALE, false},
		{"VW40-44", "40-44", SEX_FEMALE, false},
		{"VW45-49", "45-49", SEX_FEMALE, false},
		{"VW50-54", "50-54", SEX_FEMALE, false},
		{"VW55-59", "55-59", SEX_FEMALE, false},
		{"VW60-64", "60-64", SEX_FEMALE, false},
		{"VW65-69", "65-69", SEX_FEMALE, false},
		{"VW70-74", "70-74", SEX_FEMALE, false},
		{"VW75-79", "75-79", SEX_FEMALE, false},
		{"VW80-84", "80-84", SEX_FEMALE, false},
		{"VW85-89", "85-89", SEX_FEMALE, false},
		{"WWC", "WC", SEX_FEMALE, false},

		{"", "??", SEX_UNKNOWN, false},
		{"BAD", "", SEX_UNKNOWN, true},
	} {
		agegroup, err := ParseAgeGroup(test.input)
		if err != nil {
			if !test.expected_error {
				t.Errorf("unexpected error when parsing '%s': %v", test.input, err)
			}
		} else {
			if test.expected_error {
				t.Errorf("expected error not raised when parsing '%s'", test.input)
			}
			if agegroup.Age != test.age {
				t.Errorf("unexpected group when parsing '%s': %s; expected: %s", test.input, agegroup.Age, test.age)
			}
			if agegroup.Sex != test.sex {
				t.Errorf("unexpected sex when parsing '%s': %s; expected: %s", test.input, agegroup.Sex, test.sex)
			}
		}

	}
}
