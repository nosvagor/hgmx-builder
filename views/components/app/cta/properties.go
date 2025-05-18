package cta

type Props struct {
	Label string
	Path  string
}

func (p *Props) Classes() string {
	return "px-2 py-1 transition-colors rounded cursor-pointer bg-primary-501 text-base-500 hover:bg-primary-401 text-sm font-medium "
}
