package xpath

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		xpath string
	}{
		{"($prod_info/@S.ADAPTATIONID='ukpsl' and $prod_info/@S.VIEW='GLPCAFULL' and not(ancestor::tm or ancestor::entry)) or ($prod_info/@S.HPROD='811' and $prod_info/@S.VIEW='GNBFULL' and not(ancestor::tm or ancestor::entry))"},
	}
	parser := Parser{}
	for _, test := range tests {
		t.Run(test.xpath, func(t *testing.T) {
			parser.Parse(test.xpath)
		})
	}
}
