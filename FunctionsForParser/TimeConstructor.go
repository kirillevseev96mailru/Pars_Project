package FunctionsForParser

import ("strconv")

func TimeConstructor(a ...int) (string) {
	OurString := strconv.Itoa(a[0]) + "." + strconv.Itoa(a[1]) + "." + strconv.Itoa(a[2]) + "-" + 
	strconv.Itoa(a[3]) + ":" + strconv.Itoa(a[4]) + ":" + strconv.Itoa(a[5])
	return (OurString)
}