package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = worker(in, done, stage)
	}
	return in
}

func worker(in In, done In, stage Stage) Out {
	outCh := make(Bi)
	go func() {
		defer close(outCh)
		stageOutCh := stage(in)
		for {
			select {
			case <-done:
				return
			case v, ok := <-stageOutCh:
				if ok {
					outCh <- v
				} else {
					return
				}
			}
		}
	}()
	return outCh
}
