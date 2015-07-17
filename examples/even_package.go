package even

// public function
func Even(i int) bool {
    return i % 2 == 0
}

// private function
func odd(int i) bool {
    return i % 2 == 1
}

// install package
// mkdir $GOPATH/src/even
// cp even.go $GOPATH/src/even
// go build
// go install
