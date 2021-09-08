package strike

func (t *Transaction_structure) LocationInput() *Transaction_structure {
	q := t.Question
	ms := t.Answer1.MultipleSelect

	t = &Transaction_structure{
		Question: q,
		Answer1: Answer_structure{
			MultipleSelect:  ms,
			ResponseType:    "Location-Input",
		},
	}

	Update_Question_Array(t)

	return t
}