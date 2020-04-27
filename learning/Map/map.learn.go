package main

import "fmt"

func main() {
	element := make(map[string]string)
	element["A"] = "Andy"

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	name, ok := elements["Al"]
	name3, ok3 := element["Al"]

	fmt.Println(element["A"])
	fmt.Println(elements)
	fmt.Println(elements["Li"])
	fmt.Println(name, ok)
	if name2, ok2 := elements["Ne"]; ok2 {
		fmt.Println(name2, ok2)
	}
	fmt.Println(name3, ok3)

}
