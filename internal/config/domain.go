package config

import (
	"fmt"

	"github.com/google/uuid"
)

func roundRobin() func() string {

	domains := []string{
		"https://webhook.site/f66bb23d-6a44-4e3b-834b-25272e13bc2d",
		"https://webhook.site/4cefcea4-d8ba-489e-a884-9ff7f9838c28",
		"https://webhook.site/8d746520-3635-47e9-bff6-2ebaa4fd4c28",
		"https://webhook.site/43028638-6e86-43b7-a7e1-218655ad334b",
		"https://webhook.site/46ed42ca-d817-4877-8d67-218bb11df92a",
	}

	for i := 0; i < 1000; i++ {
		domains = append(domains, fmt.Sprintf(`http://%s.acefone.himanshu-gunwant.com:3000`, uuid.New()))
	}

	index := 0 // Initial index
	n := len(domains)

	return func() string {
		if n == 0 {
			return "" // Return empty string if the array is empty
		}
		// Select the current domain and update the index
		selectedDomain := domains[index]
		index = (index + 1) % n // Update index in a circular manner
		return selectedDomain
	}
}
