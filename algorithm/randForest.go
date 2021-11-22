package algorithm

import (
	"apigo/algorithm/RF"
	"fmt"
	"io/ioutil"
	"strings"

	//"strconv"
	"os"
	//"encoding/csv"
	"time"
	//"math"
)

var forest *RF.Forest

func RunRandomForestTraining() {

	start := time.Now()
	//aqui carga el dataset
	f, err := os.Open("./algorithm/dataPredecirConcurrencia.csv")
	if err != nil {
		panic("Error leyendo el archivo")
	}
	defer f.Close()

	content, _ := ioutil.ReadAll(f)
	s_content := string(content)
	lines := strings.Split(s_content, "\n")

	inputs := make([][]interface{}, 0)
	targets := make([]string, 0)
	//no tomar el nombre de la columna
	for _, line := range lines[1:] {

		line = strings.TrimRight(line, "\r\n")

		if len(line) == 0 {
			continue
		}
		tup := strings.Split(line, ",")
		pattern := tup[:len(tup)-1]
		target := tup[len(tup)-1]
		X := make([]interface{}, 0)
		for _, x := range pattern {
			X = append(X, x)
		}
		inputs = append(inputs, X)

		targets = append(targets, target)
	}
	//separacion de la data de entrenamiento y de prueba
	train_inputs := make([][]interface{}, 0)

	train_targets := make([]string, 0)

	test_inputs := make([][]interface{}, 0)
	test_targets := make([]string, 0)

	for i, x := range inputs {
		if i%2 == 1 {
			test_inputs = append(test_inputs, x)
		} else {
			train_inputs = append(train_inputs, x)
		}
	}

	for i, y := range targets {
		if i%2 == 1 {
			test_targets = append(test_targets, y)
		} else {
			train_targets = append(train_targets, y)
		}
	}
	//fmt.Println(train_inputs)

	//forest := RF.BuildForest(inputs, targets, 10, 500, len(train_inputs[0])) //100 trees
	forest = RF.BuildForest(train_inputs, train_targets, 5, 50, len(train_inputs[0])) //100 trees

	err_count := 0.0
	correct := 0
	for i := 0; i < len(test_inputs); i++ {
		output := forest.Predicate(test_inputs[i])
		expect := test_targets[i]
		//fmt.Println(output,expect)
		if output != expect {
			err_count += 1
		}
		if output == expect {
			correct += 1
		}
	}
	//para la prediccion
	var parapredecir = make([][]interface{}, 0)
	//parapredecir=make([0,1,2,3,4,5]interface{},1)
	//valores con los que se va a predecir
	var prueba = string("14.0,3.0,4,3,5")
	var tupla = strings.Split(prueba, ",")
	E := make([]interface{}, 0)
	patron := tupla[:len(tupla)]
	for _, e := range patron {
		E = append(E, e)
	}
	parapredecir = append(parapredecir, E)
	algo := make([][]interface{}, 0)
	for _, x := range parapredecir {
		algo = append(algo, x)

	}
	fmt.Println(algo[0])
	//aa = make([][] interface{},1)
	//ee = append([14,17.0,1.0,2,8,4], 1)
	//ee := forest.Predicate(parapredecir)
	prediccion := forest.Predicate(algo[0])
	fmt.Println("Prediccion de ", prueba, ": ", prediccion)
	fmt.Println("Precision: ", correct)
	fmt.Println("Cantidad de Predecidos: ", len(test_inputs))
	fmt.Println("Porcentaje de predecidos correctamente: ", (float64(correct)/float64(len(test_inputs)))*100, "%")
	fmt.Println("success rate:", 1.0-err_count/float64(len(test_inputs)))
	//aaaa := forest.Predicate(test_inputs[i])

	fmt.Println(time.Since(start))

}

func RandomForestPredict(dato1, dato2, dato3, dato4, dato5 string) string {
	//para la prediccion
	var parapredecir = make([][]interface{}, 0)

	//valores con los que se va a predecir
	tupla := []string{dato1, dato2, dato3, dato4, dato5}
	E := make([]interface{}, 0)
	patron := tupla[:len(tupla)]
	for _, e := range patron {
		E = append(E, e)
	}
	parapredecir = append(parapredecir, E)
	algo := make([][]interface{}, 0)
	for _, x := range parapredecir {
		algo = append(algo, x)

	}
	fmt.Println(algo[0])
	//aa = make([][] interface{},1)
	//ee = append([14,17.0,1.0,2,8,4], 1)
	//ee := forest.Predicate(parapredecir)
	prediccion := forest.Predicate(algo[0])
	fmt.Println("Prediccion de ", tupla, ": ", prediccion)
	return prediccion
}
