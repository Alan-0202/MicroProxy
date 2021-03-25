package retriever

type Retriever struct {
	Contexts string
}


func (r *Retriever) Get(url string) string{
	return r.Contexts
}


