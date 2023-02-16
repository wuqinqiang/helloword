package base

type Subject struct {
	title string
	text  string
}

func New(title string, text string) Subject {
	return Subject{
		title: title,
		text:  text,
	}
}

func (subject *Subject) Text() string {
	return subject.text
}

func (subject *Subject) Title() string {
	return subject.title
}
