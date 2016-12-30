package main

import (
	_ "github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
	"github.com/muesli/smolder"
)

// ProposalResource is the resource responsible for /proposals
type ProposalResource struct {
	smolder.Resource
}

var (
	_ smolder.GetIDSupported = &ProposalResource{}
	_ smolder.GetSupported   = &ProposalResource{}
	_ smolder.PostSupported  = &ProposalResource{}
	_ smolder.PutSupported   = &ProposalResource{}
)

// Register this resource with the container to setup all the routes
func (r *ProposalResource) Register(container *restful.Container, config smolder.APIConfig, context smolder.APIContextFactory) {
	r.Name = "ProposalResource"
	r.TypeName = "proposal"
	r.Endpoint = "proposals"
	r.Doc = "Manage proposals"

	r.Config = config
	r.Context = context

	r.Init(container, r)
}