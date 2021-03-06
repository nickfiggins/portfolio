// Code generated by "enumer -type=Template -json -transform=kebab"; DO NOT EDIT.

//
package enum

import (
	"encoding/json"
	"fmt"
)

const _TemplateName = "tutoring-form-confirmgeneral-form-confirmfreelance-form-confirm"

var _TemplateIndex = [...]uint8{0, 21, 41, 63}

func (i Template) String() string {
	if i < 0 || i >= Template(len(_TemplateIndex)-1) {
		return fmt.Sprintf("Template(%d)", i)
	}
	return _TemplateName[_TemplateIndex[i]:_TemplateIndex[i+1]]
}

var _TemplateValues = []Template{0, 1, 2}

var _TemplateNameToValueMap = map[string]Template{
	_TemplateName[0:21]:  0,
	_TemplateName[21:41]: 1,
	_TemplateName[41:63]: 2,
}

// TemplateString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TemplateString(s string) (Template, error) {
	if val, ok := _TemplateNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Template values", s)
}

// TemplateValues returns all values of the enum
func TemplateValues() []Template {
	return _TemplateValues
}

// IsATemplate returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Template) IsATemplate() bool {
	for _, v := range _TemplateValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Template
func (i Template) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Template
func (i *Template) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Template should be a string, got %s", data)
	}

	var err error
	*i, err = TemplateString(s)
	return err
}
