package resolvers

import (
	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
)

func campaignsApplyURL(n graphqlbackend.Namespace, c graphqlbackend.CampaignSpecResolver) string {
	return n.URL() + "/campaigns/apply/" + string(c.ID())
}

func campaignURL(n graphqlbackend.Namespace, c graphqlbackend.CampaignResolver) string {
	return n.URL() + "/campaigns/" + string(c.Name())
}
