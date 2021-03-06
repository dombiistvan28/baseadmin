package FElement

import (
	h "baseadmin/helper"
)

const (
	RADIOGROUP_TEMPLATE = `%label%%inputs%`
)

type RadioGroup struct {
	Label  string
	Radio  []InputRadio
	Static Static
}

func (rg RadioGroup) Render(errs map[string]error) string {
	h.PrintlnIf("Rendering radiogroup", h.GetConfig().Mode.Debug)
	var replaces map[string]string = make(map[string]string)
	output := RADIOGROUP_TEMPLATE
	replaces["%label%"] = rg.Label

	replaces["%inputs%"] = ""
	var inpErrors []error
	var inputName string
	for _, v := range rg.Radio {
		inpError, contains := errs[v.Name]
		if contains && inputName != v.Name {
			inpErrors = append(inpErrors, inpError)
			inputName = v.Name
		}
		replaces["%inputs%"] += "\n" + v.Render(nil)
	}

	replaces["%inputs%"] += "\n" + rg.Static.Render(nil)

	for i, v := range replaces {
		output = h.Replace(output, []string{i}, []string{v})
	}

	return GroupRender(output, rg.HasPreOrPost(), false, inpErrors, "")
}

func (rg RadioGroup) HasPreOrPost() bool {
	return false
}
