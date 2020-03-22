package main

// ErrorService .
type ErrorService struct {
}

// ShowAndRecordNotTrueErrors .
func (eD *ErrorService) ShowAndRecordNotTrueErrors(boolFlag ...bool) {
	for _, value := range boolFlag {
		if value == false {
			// Logger.Printf()
		}
	}
}
