package graph

import (
	"github.com/BoleBorson/TurboPioneer/models"
)

type InputEdge struct {
	Component models.Component
}

type OutputEdge struct {
	Component models.Component
}

type Node struct {
	Recipe  models.Recipe
	Inputs  []InputEdge
	Outputs []OutputEdge
}
