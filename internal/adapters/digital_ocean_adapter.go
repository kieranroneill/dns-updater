package adapters

import (
	"context"

	"github.com/digitalocean/godo"
)

type DigitalOceanAdapter struct {
	client *godo.Client
}

/**
 * constructors
 */

func NewDigitalOceanAdapter(token string) *DigitalOceanAdapter {
	return &DigitalOceanAdapter{
		client: godo.NewFromToken(token),
	}
}

/**
 * public methods
 */

// DomainRecordByNameAndType Gets the domain record by the name and type.
//
// Uses: /v2/domains/{domain_name}/records
//
// Returns:
//   - The domain record, or nil if no record is found.
//   - An error if there was an issue sending the request.
func (do *DigitalOceanAdapter) DomainRecordByNameAndType(ctx context.Context, domain string, name string, recordType string) (*godo.DomainRecord, error) {
	records, _, err := do.client.Domains.Records(ctx, domain, &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	})
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		if record.Type == recordType && record.Name == name {
			return &record, nil
		}
	}

	return nil, nil
}

// DomainRecordByID Gets the domain record by the record ID.
//
// Uses: /v2/domains/{domain_name}/records/{domain_record_id}
//
// Returns:
//   - The domain record, or nil if no record is found.
//   - An error if there was an issue sending the request.
func (do *DigitalOceanAdapter) DomainRecordByID(ctx context.Context, domain string, id int) (*godo.DomainRecord, error) {
	record, _, err := do.client.Domains.Record(ctx, domain, id)
	if err != nil {
		return nil, err
	}

	return record, nil
}
