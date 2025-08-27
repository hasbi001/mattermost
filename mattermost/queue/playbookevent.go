jobQueue := make(chan PlaybookEvent, 100)

func EventHandler(event PlaybookEvent) {
    jobQueue <- event
}

func Worker() {
    for {
        batch := collectBatch(jobQueue, 10, time.Minute)
        processBatch(batch)
    }
}
