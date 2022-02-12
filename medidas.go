package medidas //Se crea paquete

//declara constantes
const km_milla = 1.60934
const m_pie = 0.3048
const cm_in = 2.54

func km_a_milla (a float64) float64{
	return a*km_milla
}

func milla_a_km (a float64) float64{
	return a/km_a_milla
}

func pies_a_m (a float64) float64{
	return a*m_pie
}

func m_a_pie (a float64) float64{
	return a/m_pie
}

func in_a_cm (a float64) float64{
	return a*cm_in
}

func cm_a_in (a float64) float64{
	return a/cm_in
}