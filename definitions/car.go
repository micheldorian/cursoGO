package definitions

//Car con acceso publico
type CarType struct {
	Brand string
	Year  int
}

func UpdateCar(car *CarType, brand string, year int) {
	(*car).Brand = brand
	(*car).Year = year
}

func (car CarType) GetBrand() string {
	return car.Brand
}

func (car CarType) GetYear() int {
	return car.Year
}
