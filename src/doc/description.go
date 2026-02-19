package doc

import (
	"fmt"

	"mirage/src/models"
)

func PrintDescription(ep *models.Endpoint) {
	desc := ""
	if ep.Description != nil {
		desc = "-> " + *ep.Description
	}
	fmt.Printf("%s '%s' %s\n", ep.Method, ep.Path, desc)
}
