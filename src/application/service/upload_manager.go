package service

import "runtime"

type UploadManager struct {
	Errors []error
	Paths  []string
}

func (up *UploadManager) UploadObject(string) error {
	return nil
}

func (up *UploadManager) ProcessUpload(concurrency int, doneUpload chan string) error {

	// channel with type that will have the number of CPU's
	/*
		If the computer is an octa-core, then the channel will have 8 values (size of the channel)
		which will be used to process the upload
	*/
	in := make(chan int, runtime.NumCPU())
	// channel with type string that will be used to return the result of the upload by each goroutine
	returnChannel := make(chan string)

	for process := 0; process < concurrency; process++ {
		// invoke worker as gourtine
		go up.uploadWorker(in, returnChannel)
	}

	// Go routine that will send the file paths index to the channel
	go func() {
		for p := 0; p < len(up.Paths); p++ {
			/*
				This will work like a queue, the first goroutine that is available will receive the file path index (from paths array)
				The gourtine after geting this index value will get the value from paths array based on this passed index
				It the channel is full, then the goroutine will wait until the channel is available (The size of the channel is the number of CPU's)
			*/
			in <- p
		}
		// close the channel
		close(in)
	}()

	// Go routine that will receive the result of the upload
	for r := range returnChannel {
		if r != "" {
			doneUpload <- r
			break
		}
	}

	return nil
}

// Worker responsible to upload the file
func (up *UploadManager) uploadWorker(in chan int, returnChannel chan string) {
	// get the file path index from the channel
	// it will read the channel until it is closed
	for x := range in {
		err := up.UploadObject(up.Paths[x])
		if err != nil {
			// add the error to list of errors
			up.Errors = append(up.Errors, err)
			// notify the channel that the upload failed
			returnChannel <- err.Error()
		}

		// notify the channel that the upload was successful
		returnChannel <- "completed"
	}
}
