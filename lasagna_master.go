package lasagna

import "strings"

// TODO: define the 'PreparationTime()' function
// Estimate the preparation time
func PreparationTime(layers []string, avgTime int) int {
	if avgTime <= 0 {
		return len(layers) * 2
	}
	return len(layers) * avgTime
}

// TODO: define the 'Quantities()' function
// Compute the amounts of noodles and sauce needed
func Quantities(layers []string) (qtyNoodles int, qtySauce float64) {
	for _, v := range layers {
		if strings.ToLower(v) == "noodles" {
			qtyNoodles += 50
		} else if strings.ToLower(v) == "sauce" {
			qtySauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
// Add the secret ingredient
func AddSecretIngredient(friendList, myList []string) []string {
	return append(myList, friendList[len(friendList)-1])
}

// TODO: define the 'ScaleRecipe()' function
// Scale the recipe
func ScaleRecipe(quantities []float64, portions int) (qtyPortions []float64) {
	for _, v := range quantities {
		qtyPortions = append(qtyPortions, v*float64(portions)/2)
	}
	return
}
