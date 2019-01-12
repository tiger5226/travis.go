package travis

import "fmt"

// Repository is information about the repository.
type Repository struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	OwnerName string `json:"owner_name"`
	URL       string `json:"url"`
}

func (r Repository) String() string {
	return fmt.Sprintf("Repository: ID %d, Name %s, OwnerName %s, URL %s", r.ID, r.Name, r.OwnerName, r.URL)
}
