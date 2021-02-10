package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// FoodDisplayTable ...
type FoodDisplayTable struct {
    XMLName xml.Name `xml:"Food_Display_Table" json:"-"`
	FoodDisplayRow []FoodDisplayRow `xml:"Food_Display_Row" json:"data"`
}

// FoodDisplayRow ...
type FoodDisplayRow struct {
	XMLName xml.Name `xml:"Food_Display_Row" json:"-"`
	FoodCode int `xml:"Food_Code" json:"food_code"`
	DisplayName string `xml:"Display_Name" json:"display_name"`
	PortionDefault float32 `xml:"Portion_Default" json:"portion_default"`
	PortionAmount float32 `xml:"Portion_Amount" json:"portion_amount"`
	PortionDisplayName string `xml:"Portion_Display_Name" json:"portion_display_name"`
	Factor float32 `json:"factor"`
	Increment float32 `json:"increment"`
	Multiplier float32 `json:"multiplier"`
	Grains float32 `json:"grains"`
	WholeGrains float32 `xml:"Whole_Grains" json:"whole_grains"`
	Vegetables float32 `json:"vegetables"`
	OrangeVegetables float32 `xml:"Orange_Vegetables" json:"orange_vegetables"`
	DarkgreenVegetables  float32 `xml:"Drkgreen_Vegetables" json:"darkgreen_vegetables"`
	Starchyvegetables float32 `xml:"Starchy_vegetables" json:"starchy_vegetables"`
	OtherVegetables float32 `xml:"Other_Vegetables" json:"other_vegetables"`
	Fruits float32 `json:"fruits"`
	Milk float32 `json:"milk"`
	Meats float32 `json:"meats"`
	Soy float32 `json:"soy"`
	DrybeansPeas float32 `xml:"Drybeans_Peas" json:"drybeans_peas"`
	Oils float32 `json:"oils"`
	SolidFats float32 `xml:"Solid_Fats" json:"solid_fats"`
	AddedSugars float32 `xml:"Added_Sugars" json:"added_sugars"`
	Alcohol float32 `json:"alcohol"`
	Calories float32 `json:"calories"`
	SaturatedFats float32 `xml:"Saturated_Fats" json:"saturated_fats"`
}

func main()  {
	startTimer := time.Now()
	xmlFile, err := os.Open("DataSet/myfoodapediadata/Food_Display_Table.xml")
	if err != nil {
        panic(err.Error())
    }
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
        panic(err.Error())
    }

	var foodDisplayTable FoodDisplayTable
	xml.Unmarshal(byteValue, &foodDisplayTable)

	jsonBytes, err := json.MarshalIndent(foodDisplayTable.FoodDisplayRow, "", "  ")
	if err != nil {
		panic(err.Error())
	}

	createdFile, err := os.Create("some.json")
	if err != nil {
		panic(err.Error())
	}

	_, err = createdFile.Write(jsonBytes)
	if err != nil {
		panic(err.Error())
	}

	err = createdFile.Close()
	if err != nil {
		panic(err.Error())
	}

	timeElapsed := time.Since(startTimer)
	fmt.Printf("converted in %s ✨\n", timeElapsed)
}