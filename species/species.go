package species

func AnimalSpecies(input int) string {
	species := [4]string{
		"Cow",
		"Pig",
		"Horse",
		"Chicken",
		}
	return species[input - 1]
}