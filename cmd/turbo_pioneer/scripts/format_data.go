package scripts

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/BoleBorson/TurboPioneer/models"
)

func FormatData() {
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

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	// Open or create a file
	file, err := os.Create("recipes.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after writing

	// Write JSON data to file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data successfully written to recipes.json")

	// for key, recipe := range data.Recipes {
	// 	fmt.Printf("Recipe Key: %s\n", key)
	// 	fmt.Printf("Name: %s\n", recipe.Name)
	// 	fmt.Printf("Slug: %s\n", recipe.Slug)
	// 	fmt.Printf("Ingredients:\n")
	// 	for _, ing := range recipe.Ingredients {
	// 		fmt.Printf("  Item: %s, Amount: %.2f\n", ing.Item, ing.Amount)
	// 		item := ing.ItemDefinition // Replace this with actual retrieval if needed
	// 		fmt.Printf("  Item:\n")
	// 		fmt.Printf("    Slug: %s\n", item.Slug)
	// 		fmt.Printf("    ClassName: %s\n", item.ClassName)
	// 		fmt.Printf("    Name: %s\n", item.Name)
	// 		fmt.Printf("    SinkPoints: %d\n", item.SinkPoints)
	// 		fmt.Printf("    Description: %s\n", item.Description)
	// 		fmt.Printf("    StackSize: %d\n", item.StackSize)
	// 		fmt.Printf("    EnergyValue: %f\n", item.EnergyValue)
	// 		fmt.Printf("    RadiativeDecay: %f\n", item.RadiativeDecay)
	// 		fmt.Printf("    Liquid: %t\n", item.Liquid)
	// 		fmt.Printf("    FluidColor: %v\n", item.FluidColor) // Adjust if Color type has a specific format

	// 	}
	// 	fmt.Printf("Products:\n")
	// 	for _, prod := range recipe.Products {
	// 		fmt.Printf("  Item: %s, Amount: %.2f\n", prod.Item, prod.Amount)
	// 		item := prod.ItemDefinition // Replace this with actual retrieval if needed
	// 		fmt.Printf("  Item:\n")
	// 		fmt.Printf("    Slug: %s\n", item.Slug)
	// 		fmt.Printf("    ClassName: %s\n", item.ClassName)
	// 		fmt.Printf("    Name: %s\n", item.Name)
	// 		fmt.Printf("    SinkPoints: %d\n", item.SinkPoints)
	// 		fmt.Printf("    Description: %s\n", item.Description)
	// 		fmt.Printf("    StackSize: %d\n", item.StackSize)
	// 		fmt.Printf("    EnergyValue: %f\n", item.EnergyValue)
	// 		fmt.Printf("    RadiativeDecay: %f\n", item.RadiativeDecay)
	// 		fmt.Printf("    Liquid: %t\n", item.Liquid)
	// 		fmt.Printf("    FluidColor: %v\n", item.FluidColor) // Adjust if Color type has a specific format
	// 	}
	// 	return
	// }

}
