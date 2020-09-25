package ihop

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

const (
	TestDataDir   = "./testdata"
	StackLenLimit = 100
	TestCaseLimit = 100
)

//TestPancakeFlips Read all files in the testdata directory and for each file run all its test cases
func TestPancakeFlips(t *testing.T) {
	var testCaseFiles []string

	err := getTestFileNames(&testCaseFiles)

	if err != nil {
		t.Errorf("%v", err)
	}

	for _, name := range testCaseFiles {
		err := runTestCases(t, name)
		if err != nil {
			t.Errorf("%s: %v", name, err)
		}
	}
}

// runTestCase - given a file name open the file and run all of the test cases contained in the file
func runTestCases(t *testing.T, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	bf := bufio.NewReader(file)

	testCount := 0

	// read in first line for the number of test cases contained in the file
	line, _, err := bf.ReadLine()
	if err == io.EOF {
		return fmt.Errorf("test file is empty")
	}

	numTests, err := strconv.Atoi(string(line))
	if err != nil {
		return fmt.Errorf("test file must start with the number of test cases contained in the file")
	}

	fmt.Println("Test Filename: ", filename)

	for {
		line, _, err := bf.ReadLine()

		if err == io.EOF {
			break
		}

		if len(line) > StackLenLimit {
			t.Logf("WARNING: test file contains a stack string that is longer that the current limit (%d)", StackLenLimit)
			continue
		}

		str := string(line)

		if err = ValidateStack(str); err != nil {
			return err
		}

		if testCount > TestCaseLimit {
			return fmt.Errorf("test file contains too many test cases (%d). Limit is %d.", testCount, TestCaseLimit)
		}

		testCount++
		flips := PancakeFlips(str)

		fmt.Printf("Case #%v:%4v\n", testCount, flips)
	}

	if testCount != numTests {
		return fmt.Errorf("test file does not have the number of test cases specified in the file (%d)", numTests)
	}

	fmt.Println("-----------")

	return nil
}

// getTestFileNames Get a list of filenames for each file found in the testdata directory
func getTestFileNames(testCaseFiles *[]string) error {
	f, err := os.Open(TestDataDir)
	if err != nil {
		return errors.New("Error: cannot read from the testdata directory")
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return errors.New("Error: listing the testdata directory")
	}

	for _, file := range files {
		*testCaseFiles = append(*testCaseFiles, fmt.Sprintf("%s/%s", TestDataDir, file.Name()))
	}
	return nil
}
