package testutils

import "os"

func SimulateStdinInput(input string) (func(), error) {
	oldStdin := os.Stdin
	r, w, err := os.Pipe()

	if err != nil {
		return nil, err
	}

	_, err = w.WriteString(input)
	if err != nil {
		return nil, err
	}

	w.Close()

	os.Stdin = r

	return func() {
		os.Stdin = oldStdin
	}, nil
}
