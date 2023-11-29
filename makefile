test-manipulate:
	go test -run TestManipulate golang/manipulate_test.go golang/manipulate.go -v

test-odd:
	go test -run TestFindOddNumber golang/odd_number_test.go golang/odd_number.go -v

test-smily:
	go test -run TestCountSmilyFace golang/smily_test.go golang/smily.go -v