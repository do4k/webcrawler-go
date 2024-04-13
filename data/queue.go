package data

type Queue []string

func (q *Queue) Enqueue(str string) {
	*q = append(*q, str)
}

func (q *Queue) Dequeue() (string, bool) {
	if len(*q) == 0 {
		return "", false
	}

	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

func (q *Queue) Contains(str string) bool {
	for _, element := range *q {
		if element == str {
			return true
		}
	}

	return false
}

func (q *Queue) Count() int {
	return len(*q)
}
