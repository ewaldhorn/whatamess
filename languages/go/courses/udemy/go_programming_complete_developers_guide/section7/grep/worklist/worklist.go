package worklist

type Entry struct {
	Path string
}

type WorkList struct {
	jobs chan Entry
}

func (w *WorkList) Add(work Entry) {
	w.jobs <- work
}

func (w *WorkList) Next() Entry {
	j := <-w.jobs
	return j
}

func New(bufferSize int) WorkList {
	return WorkList{make(chan Entry, bufferSize)}
}

func NewJob(path string) Entry {
	return Entry{path}
}

func (w *WorkList) Finalize(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		w.Add(Entry{""})
	}
}
