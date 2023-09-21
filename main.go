package main

func main() {
	model := Model{
		Name:         "Feed-Forward Neural Network",
		LearningRate: 0.0001,
		Optimizer:    Adam,
		Loss:         BinaryCrossEntropy,
	}

	model.SetLayer(Convolution{})
	model.SetLayer(Flatten{})
	model.SetLayer(Convolution{})
	model.SetLayer(Dense{})
	model.SetLayer(Dense{})

	loss, accuracy := Train(&model)
	Summary(loss, accuracy)
}
