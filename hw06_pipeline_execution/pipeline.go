package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		stg := stage(in)
		in = checkDone(stg, done)
	}

	return in
}

func checkDone(stg In, done In) Out {
	out := make(Bi)

	go func() {
		defer func() {
			close(out)
		}()

		for {
			select {
			case v, ok := <-stg:
				if !ok {
					return
				}
				out <- v
			case <-done:
				return
			}
		}
	}()
	return out
}
