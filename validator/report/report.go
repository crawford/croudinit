package report

type Entry struct {
	Line    int
	Message string
}

type Report struct {
	errors   []Entry
	warnings []Entry
}

func (r *Report) Error(line int, message string) {
	r.errors = append(r.errors, Entry{line, message})
}

func (r *Report) Warning(line int, message string) {
	r.warnings = append(r.warnings, Entry{line, message})
}

func (r *Report) Errors() []Entry {
	return r.errors
}

func (r *Report) Warnings() []Entry {
	return r.warnings
}
