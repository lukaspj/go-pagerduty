package pagerduty

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
)

// StatusDashboard represents a status dashboard.
type StatusDashboard struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	UrlSlug string `json:"url_slug,omitempty"`
}

// ListStatusDashboardsResponse represents a list response of business services.
type ListStatusDashboardsResponse struct {
	StatusDashboards []*StatusDashboard `json:"status_dashboards,omitempty"`
	NextCursor       string             `json:"next_cursor,omitempty"`
	Limit            uint               `json:"limit,omitempty"`
}

// ListStatusDashboardsOptions is the data structure used when calling the ListStatusDashboards API endpoint.
type ListStatusDashboardsOptions struct {
	// NextCursor is an opaque string than will deliver the next set of results when
	// provided as the cursor parameter in a subsequent request. A null value for this
	// field indicates that there are no additional results.
	NextCursor string `json:"next_cursor,omitempty"`
}

type ImpactedServiceAdditionalFields struct {
	HighestImpactingPriority *struct {
		ID    string `json:"id,omitempty"`
		Order int    `json:"order,omitempty"`
	}
}

type ImpactedService struct {
	ID               string                          `json:"id,omitempty"`
	Name             string                          `json:"name,omitempty"`
	Type             string                          `json:"type,omitempty"`
	Status           string                          `json:"status,omitempty"`
	AdditionalFields ImpactedServiceAdditionalFields `json:"additional_fields"`
}

// ListStatusDashboardsResponse represents a list response of business services.
type ImpactedServicesListResponse struct {
	APIListObject
	Services []*ImpactedService `json:"services,omitempty"`
}

// ListStatusDashboardServiceImpactsOptions is the data structure used when calling the ListStatusDashboards API endpoint.
type ListStatusDashboardServiceImpactsOptions struct {
	AdditionalFields string `json:"additional_fields,omitempty"`
}

// ListStatusDashboards lists existing status dashboards.
func (c *Client) ListStatusDashboards(ctx context.Context, o ListStatusDashboardsOptions) ([]*StatusDashboard, error) {
	queryParms, err := query.Values(o)
	if err != nil {
		return nil, err
	}

	response, err := c.get(ctx, "/status_dashboards?"+queryParms.Encode(), nil)

	var result ListStatusDashboardsResponse
	if err := c.decodeJSON(response, &result); err != nil {
		return nil, err
	}

	return result.StatusDashboards, nil
}

// ImpactedServicesByStatusDashboardUrlSlug lists existing status dashboards.
func (c *Client) ImpactedServicesByStatusDashboardUrlSlug(ctx context.Context, url_slug string, o ListStatusDashboardServiceImpactsOptions) ([]*ImpactedService, error) {
	queryParms, err := query.Values(o)
	if err != nil {
		return nil, err
	}

	response, err := c.get(ctx, fmt.Sprintf("/status_dashboards/url_slugs/%s/service_impacts?", url_slug)+queryParms.Encode(), nil)

	var result ImpactedServicesListResponse
	if err := c.decodeJSON(response, &result); err != nil {
		return nil, err
	}

	return result.Services, nil
}
