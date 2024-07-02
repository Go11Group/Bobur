package handlers

import "api_get_way/genproto"

type Handler struct {
	Composition   genproto.CompositionServiceClient
	Discovery     genproto.DiscoveryServiceClient
	Collaboration genproto.CollaborationServiceClient
}
