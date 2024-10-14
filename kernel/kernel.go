package kernel

import "fmt"

type Kernel struct {
	Matrix [][]float64
	Size   int
}

func NewKernel(size int, matrix [][]float64) *Kernel {
	k := &Kernel{
		Size:   size,
		Matrix: make([][]float64, size),
	}
	for i := range k.Matrix {
		k.Matrix[i] = make([]float64, size)
	}
	if matrix != nil {
		err := k.SetAll(matrix)
		if err != nil {
			panic(err)
		}
	}
	return k
}

func (k *Kernel) Get(x, y int) float64 {
	return k.Matrix[x][y]
}

func (k *Kernel) Set(x, y int, value float64) error {
	if x < 0 || y < 0 || x >= k.Size || y >= k.Size {
		return fmt.Errorf("index out of bounds")
	}
	k.Matrix[x][y] = value
	return nil
}

func (k *Kernel) SetAll(values [][]float64) error {
	if len(values) != k.Size {
		return fmt.Errorf("incorrect number of rows")
	}
	for i := range values {
		if len(values[i]) != k.Size {
			return fmt.Errorf("incorrect number of columns")
		}
	}
	k.Matrix = values
	return nil
}
