package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Bi)

	go func() {
		defer close(out)
		for _, stage := range stages {
			if done != nil && <-done == nil {
				return
			}
			in = stage(in)
		}
		for v := range in {
			out <- v
		}
	}()

	return out
}
