type TemperatureKey struct {
	Index int
	Value int
}

func NewTemperatureKey(index int, value int) TemperatureKey {
	return TemperatureKey{Index: index, Value: value}
}

func dailyTemperatures(temperatures []int) []int {
	// Define the min stack
	minTemperatureStack := make([]TemperatureKey, 0, len(temperatures))
	// Define the answer
	ans := make([]int, len(temperatures))

	for i, temperature := range temperatures {
		currentTemp := NewTemperatureKey(i, temperature)
		// Case when the temperature stack has just been initialised
		if len(minTemperatureStack) == 0 {
			minTemperatureStack = append(minTemperatureStack, currentTemp)
			continue
		}
		// Case where it is a hotter day, the days that came before it are cooler.
		// Open a window where we drain the stack.
		for len(minTemperatureStack) > 0 {
			// Edge case when as we're making our day down the stack, the days get warmer
			// Where in the stack of 50 40 30, temperature is 32
			// Peek the top of the stack
			peek := minTemperatureStack[len(minTemperatureStack)-1]
			if peek.Value >= currentTemp.Value {
				break 
			}
			ans[peek.Index] = currentTemp.Index - peek.Index
			minTemperatureStack = minTemperatureStack[:len(minTemperatureStack)-1]
		}

		// Add the cooler temperature day
		minTemperatureStack = append(minTemperatureStack, currentTemp)
	}

	return ans
}
