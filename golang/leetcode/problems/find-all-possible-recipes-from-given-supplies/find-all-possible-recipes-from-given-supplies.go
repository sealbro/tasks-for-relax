package find_all_possible_recipes_from_given_supplies

import "slices"

// https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	mapSupplies := make(map[string]struct{}, len(supplies))
	for _, supplie := range supplies {
		mapSupplies[supplie] = struct{}{}
	}

	mapRecepts := make(map[string][]string, len(recipes))
	for _, recipe := range recipes {
		mapRecepts[recipe] = nil
	}

	for i, ingredient := range ingredients {
		recipe := recipes[i]
		for _, ing := range ingredient {
			if _, ok := mapSupplies[ing]; !ok {
				if _, ok := mapRecepts[ing]; ok {
					mapRecepts[recipe] = append(mapRecepts[recipe], ing)
				} else {
					delete(mapRecepts, recipe)
				}
			}
		}
	}

	var results []string
	for k, v := range mapRecepts {
		if len(v) == 0 {
			results = append(results, k)
		} else {
			allIng := true
			for _, subRec := range v {
				if vv, ok := mapRecepts[subRec]; ok {
					// recursive
					for _, sr := range vv {
						if vvv, ok := mapRecepts[sr]; ok {
							if slices.Contains(vvv, subRec) {
								allIng = false
								break
							}
						}
					}
				} else {
					allIng = false
				}
				if !allIng {
					break
				}
			}
			if allIng {
				results = append(results, k)
			}
		}
	}

	return results
}
