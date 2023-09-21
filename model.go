package main

type Model struct {
	Name         string
	Layers       []Layer
	Optimizer    string
	Loss         string
	LearningRate float64
}

func (m *Model) SetLayer(layer Layer) {
	m.Layers = append(m.Layers, layer)
}

func (m *Model) SetOptimizer(opt string) {
	m.Optimizer = opt
}

func (m *Model) SetLoss(loss string) {
	m.Loss = loss
}
