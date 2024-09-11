package scripts

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/BoleBorson/TurboPioneer/graph"
	"github.com/BoleBorson/TurboPioneer/models"
)

func BuildGraph() {
	jsonFile, err := os.Open("/home/cole/code-projects/TurboPioneer/data/data1.0.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	// we initialize our Users array
	var data models.Data

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	removals := make([]string, 369)
	for i := range data.Recipes {
		value := data.Recipes[i]
		for y := range value.Ingredients {
			found := data.Items[value.Ingredients[y].Item]
			value.Ingredients[y].ItemDefinition = found
		}
		for z := range value.Products {
			found, ok := data.Items[value.Products[z].Item]
			if !ok {
				removals = append(removals, i)
			}
			value.Products[z].ItemDefinition = found
		}
	}
	// remove buildings as we don't care about them for our purposes
	for item := range removals {
		delete(data.Recipes, removals[item])
	}

	recipeGraph := make(map[string]graph.Node, len(data.Recipes))

	for i := range data.Recipes {
		inputs := make([]graph.InputEdge, 0, len(data.Recipes[i].Ingredients))
		outputs := make([]graph.OutputEdge, 0, len(data.Recipes[i].Products))

		value := data.Recipes[i]
		for y := range value.Ingredients {
			ing := value.Ingredients[y]
			inputs = append(inputs, graph.InputEdge{
				Component: ing,
			})
		}
		for z := range value.Products {
			product := value.Products[z]
			outputs = append(outputs, graph.OutputEdge{
				Component: product,
			})
		}
		recipeGraph[data.Recipes[i].Name] = graph.Node{
			Recipe:  data.Recipes[i],
			Inputs:  inputs,
			Outputs: outputs,
		}
	}

	TraverseIngredients("AI Limiter", recipeGraph)

}

func TraverseIngredients(recipeName string, recipeGraph map[string]graph.Node) {
	recipe, ok := recipeGraph[recipeName]
	if !ok {
		return
	}
	fmt.Println(recipeName + " | Ingredients:")
	for ing := range recipe.Inputs {
		ingName := recipe.Inputs[ing].Component.ItemDefinition.Name
		fmt.Println(ingName)
		TraverseIngredients(ingName, recipeGraph)
	}
}
