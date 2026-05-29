.PHONY: test fmt vet cover world

world: test fmt vet cover 

test:
	go test ./learning/1_fundamentals
	go test ./learning/2_Data_Modeling

test_verbose:
	go test -v ./learning/1_fundamentals
	go test -v ./learning/2_Data_Modeling

vet:
	go vet ./learning/1_fundamentals
	go vet ./learning/2_Data_Modeling


cover:
	go test -cover ./learning/1_fundamentals
	go test -cover ./learning/2_Data_Modeling

fmt:
	go fmt ./learning/1_fundamentals
	go fmt ./learning/2_Data_Modeling