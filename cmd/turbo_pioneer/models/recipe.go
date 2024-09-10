package models

type Color struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
	A int `json:"a"`
}

// Define a struct for the main object
type Item struct {
	Slug           string  `json:"slug"`
	ClassName      string  `json:"className"`
	Name           string  `json:"name"`
	SinkPoints     int     `json:"sinkPoints"`
	Description    string  `json:"description"`
	StackSize      int     `json:"stackSize"`
	EnergyValue    float64 `json:"energyValue"`
	RadiativeDecay float64 `json:"radioactiveDecay"`
	Liquid         bool    `json:"liquid"`
	FluidColor     Color   `json:"fluidColor"`
}

type Component struct {
	Item           string  `json:"item"`
	Amount         float64 `json:"amount"`
	ItemDefinition Item    `json:"itemDefinition"`
}

type Recipe struct {
	Slug                 string      `json:"slug"`
	Name                 string      `json:"name"`
	ClassName            string      `json:"className"`
	Alternate            bool        `json:"alternate"`
	Time                 int64       `json:"time"`
	ManualTimeMultiplier float64     `json:"manualTimeMultiplier"`
	Ingredients          []Component `json:"ingredients"`
	ForBuilding          bool        `json:"forBuilding"`
	InMachine            bool        `json:"inMachine"`
	InWorkshop           bool        `json:"inWorkshop"`
	Products             []Component `json:"products"`
	ProducedIn           []string    `json:"producedIn"`
	IsVariablePower      bool        `json:"isVariablePower"`
}

type Data struct {
	Recipes map[string]Recipe `json:"recipes"`
	Items   map[string]Item   `json:"items"`
}
