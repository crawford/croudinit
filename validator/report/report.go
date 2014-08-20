package report

type entry struct {
	line    int
	message string
}

type Report struct {
	errors   []entry
	warnings []entry
}

func (r *Report) Error(line int, message string) {
	r.errors = append(r.errors, entry{line, message})
}

func (r *Report) Warning(line int, message string) {
	r.warnings = append(r.warnings, entry{line, message})
}

func NewReport() *Report {
	return &Report{
		errors: make([]entry, 0),
	}
}
