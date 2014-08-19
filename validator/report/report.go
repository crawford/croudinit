package report

type Report struct {
	errors []struct {
		line    int
		message string
	}
}

func (r *Report) Error(line int, message string) {
	r.errors = append(r.errors, struct {
		line    int
		message string
	}{line, message})
}

func NewReport() *Report {
	return &Report{
		errors: make([]struct {
			line    int
			message string
		}, 0),
	}
}
