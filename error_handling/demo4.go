package error_handling

import "fmt"

type borderExeption struct {
	param   int
	message string
}

func (b *borderExeption) Error() string {
	return fmt.Sprintf("%d---%s", b.param, b.message)
}

func Demo4(param int) (string, error) {
	if param < 1 || param > 100 {
		return "", &borderExeption{param, "sınırların dışı"}
	}
	return "Doğru", nil
}
