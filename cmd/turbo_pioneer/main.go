package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/BoleBorson/TurboPioneer/models"
)

func main() {
	jsonFile, err := os.Open("/home/cole/code-projects/TurboPioneer/data/data.json")
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

	for _, value := range data.Recipes {
		for _, ivalue := range value.Ingredients {
			itemName := &ivalue.Item
			ivalue.ItemDefinition = data.Items[*itemName]
		}
	}

	fmt.Println(data.Items["Desc_GoldIngot_C"])

	for key, recipe := range data.Recipes {
		fmt.Printf("Recipe Key: %s\n", key)
		fmt.Printf("Name: %s\n", recipe.Name)
		fmt.Printf("Slug: %s\n", recipe.Slug)
		fmt.Printf("Ingredients:\n")
		for _, ing := range recipe.Ingredients {
			fmt.Printf("  Item: %s, Amount: %.2f\n", ing.Item, ing.Amount)
			item := ing.ItemDefinition // Replace this with actual retrieval if needed
			fmt.Printf("  Item:\n")
			fmt.Printf("    Slug: %s\n", item.Slug)
			fmt.Printf("    ClassName: %s\n", item.ClassName)
			fmt.Printf("    Name: %s\n", item.Name)
			fmt.Printf("    SinkPoints: %d\n", item.SinkPoints)
			fmt.Printf("    Description: %s\n", item.Description)
			fmt.Printf("    StackSize: %d\n", item.StackSize)
			fmt.Printf("    EnergyValue: %f\n", item.EnergyValue)
			fmt.Printf("    RadiativeDecay: %f\n", item.RadiativeDecay)
			fmt.Printf("    Liquid: %t\n", item.Liquid)
			fmt.Printf("    FluidColor: %v\n", item.FluidColor) // Adjust if Color type has a specific format

		}
		fmt.Printf("Products:\n")
		for _, prod := range recipe.Products {
			fmt.Printf("  Item: %s, Amount: %.2f\n", prod.Item, prod.Amount)
			item := prod.ItemDefinition // Replace this with actual retrieval if needed
			fmt.Printf("  Item:\n")
			fmt.Printf("    Slug: %s\n", item.Slug)
			fmt.Printf("    ClassName: %s\n", item.ClassName)
			fmt.Printf("    Name: %s\n", item.Name)
			fmt.Printf("    SinkPoints: %d\n", item.SinkPoints)
			fmt.Printf("    Description: %s\n", item.Description)
			fmt.Printf("    StackSize: %d\n", item.StackSize)
			fmt.Printf("    EnergyValue: %f\n", item.EnergyValue)
			fmt.Printf("    RadiativeDecay: %f\n", item.RadiativeDecay)
			fmt.Printf("    Liquid: %t\n", item.Liquid)
			fmt.Printf("    FluidColor: %v\n", item.FluidColor) // Adjust if Color type has a specific format
		}
		return
	}

}
