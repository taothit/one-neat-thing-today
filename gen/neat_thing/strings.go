package neatthing

import (
	"fmt"
	"strings"
)

func (nt *NeatThing) String() string {
	if nt == nil {
		return "<nil>"
	}

	var builder = strings.Builder{}
	builder.WriteString("NeatThing{")
	if nt.Name != nil {
		builder.WriteString(fmt.Sprintf("name=(%s), ", *nt.Name))
	}

	if nt.Date != nil {
		builder.WriteString(fmt.Sprintf("date=(%s), ", *nt.Date))
	}

	if nt.Definition != nil {
		builder.WriteString(fmt.Sprintf("definition=(%s), ", *nt.Definition))
	}

	if nt.Link != nil {
		builder.WriteString(fmt.Sprintf("link=(%s), ", *nt.Link))
	}

	if nt.Bibliography != nil && len(nt.Bibliography) > 0 {
		builder.WriteString("bibliography={")
		for _, entry := range nt.Bibliography {
			builder.WriteString(fmt.Sprintf("entry=%s, ", entry))
		}
		builder.WriteRune('}')
	} else {
		builder.WriteString("bibliography=<nil>")
	}

	builder.WriteRune('}')
	return builder.String()
}
