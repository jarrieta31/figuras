package figuras

type Cuadrado struct {
	Lado float64
}

func (cua *Cuadrado) perimetro() float64 {
	return cua.Lado * 4
}

func (cua *Cuadrado) area() float64 {
	return cua.Lado * cua.Lado
}
