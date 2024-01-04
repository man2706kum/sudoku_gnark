package sudoku

import (
	"fmt"

	"github.com/consensys/gnark/frontend"
)

func Print2() {
	fmt.Println("hello")
}

type Circuit struct {
	Puzzle   [9][9]frontend.Variable `gnark:",public"`
	Solution [9][9]frontend.Variable
}

func (circuit *Circuit) check_puzzle_and_solution_validity(api frontend.API) error {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//All entries in the puzzle is less than or equal to 9
			api.AssertIsLessOrEqual(circuit.Puzzle[i][j], 9)

			//All entries in the solution is less than or equal to 9
			api.AssertIsLessOrEqual(circuit.Solution[i][j], 9)

			// either puzzle[i][j] == 0 or both puzzle[i][j] and solution[i][j] agrees ie P[i][j] * (P[i][j] - S[i][j]) == 0
			api.AssertIsEqual(api.Mul(circuit.Puzzle[i][j], api.Sub(circuit.Puzzle[i][j], circuit.Solution[i][j])), 0)
		}
	}

	return nil
}

func (circuit *Circuit) check_solution(api frontend.API) error {

	//check uniqueness in rows & column
	for k := 0; k < 9; k++ {
		for i := 0; i < 9; i++ {
			for j := i + 1; j < 9; j++ {
				api.AssertIsDifferent(circuit.Solution[k][i], circuit.Solution[k][j]) //rows uniqueness
				api.AssertIsDifferent(circuit.Solution[i][k], circuit.Solution[j][k]) // column uniqueness
			}
		}
	}

	//check uniqueness in 3*3 box
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			api.AssertIsDifferent(circuit.Solution[3*i+0][3*j+0], circuit.Solution[3*i+0][3*j+1])
			api.AssertIsDifferent(circuit.Solution[3*i+0][3*j+0], circuit.Solution[3*i+0][3*j+2])

			api.AssertIsDifferent(circuit.Solution[3*i+0][3*j+0], circuit.Solution[3*i+1][3*j+0])
			api.AssertIsDifferent(circuit.Solution[3*i+0][3*j+0], circuit.Solution[3*i+1][3*j+1])
			api.AssertIsDifferent(circuit.Solution[3*i+0][3*j+0], circuit.Solution[3*i+1][3*j+2])

			api.AssertIsDifferent(circuit.Solution[3*i+0][3*j+0], circuit.Solution[3*i+2][3*j+0])
			api.AssertIsDifferent(circuit.Solution[3*i+0][3*j+0], circuit.Solution[3*i+2][3*j+1])
			api.AssertIsDifferent(circuit.Solution[3*i+0][3*j+0], circuit.Solution[3*i+2][3*j+2])
		}

	}

	return nil
}

func (circuit *Circuit) Define(api frontend.API) error {
	circuit.check_puzzle_and_solution_validity(api)
	circuit.check_solution(api)
	return nil
}
