package Perceptron

import (
	. "gocurry"
)

/////// Evaluation functions

func ExternalPerceptronDot_primUs_dotProduct(task *Task) {
	root := task.GetControl()
	loop1 := root.GetChild(0)
	loop2 := root.GetChild(1)

	var result float64 = 0.0

	for loop2.GetNumChildren() > 1 {
		value2 := loop2.GetChild(0)
		if loop1.GetNumChildren() > 1 {
			value1 := loop1.GetChild(0)
			result += value1.GetFloat() * value2.GetFloat()
			loop1 = loop1.GetChild(1)
			if !loop1.IsHnf() {
				task.ToHnf(loop1)
			}
		} else {
			result += value2.GetFloat()
		}
		loop2 = loop2.GetChild(1)
		if !loop2.IsHnf() {
			task.ToHnf(loop2)
		}
	}

	FloatLitCreate(root, result)
}
