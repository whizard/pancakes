# Revenge of the Pancakes<!-- omit in toc -->
<p align="center">pfmackay@gmail.com</p>

---

## Abstract <!-- omit in toc -->

This problem came from a Google Code Jam project that was made available online in 2016.

The Infinite House of Pancakes has just introduced a new kind of pancake! It has a happy face made of chocolate chips on one side (the "happy side"), and nothing on the other side (the "blank side").

You are the head waiter on duty, and the kitchen has just given you a stack of pancakes to serve to a customer. Like any good pancake server, you have X­ray pancake vision, and you can see whether each pancake in the stack has the happy side up or the blank side up. You think the customer will be happiest if every pancake is happy side up when you serve them.

---

## Table of Contents <!-- omit in toc -->

- [Background](#background)
- [Goal](#goal)
- [Components](#components)
  - [Building The CLI Application](#building-the-cli-application)
  - [Running The CLI Application](#running-the-cli-application)
  - [Testing](#testing)
    - [Test Cases File Format](#test-cases-file-format)
      - [Test Cases File Example](#test-cases-file-example)
- [References](#references)

---

## Background

The Infinite House of Pancakes has just introduced a new kind of pancake! It has a happy face made of chocolate chips on one side (the "happy side"), and nothing on the other side (the "blank side").

You are the head waiter on duty, and the kitchen has just given you a stack of pancakes to serve to a customer. Like any good pancake server, you have X­ray pancake vision, and you can see whether each pancake in the stack has the happy side up or the blank side up. You think the customer will be happiest if every pancake is happy side up when you serve them.

You know the following maneuver: carefully lift up some number of pancakes (possibly all of them) from the top of the stack, flip that entire group over, and then put the group back down on top of any pancakes that you did not lift up. When flipping a group of pancakes, you flip the entire group in one motion; you do not individually flip each pancake. Formally: if we number the pancakes 1, 2, ..., N from top to bottom, you choose the top i pancakes to flip. Then, after the flip, the stack is i, i­1, ..., 2, 1, i+1, i+2, ..., N. Pancakes 1, 2, ..., now have the opposite side up, whereas pancakes i+1, i+2, ..., N have the same side up that they had up before.

For example, let's denote the happy side as '+' and the blank side as '­-'. Suppose that the stack, starting from the top, is ­­'+­'. One valid way to execute the maneuver would be to pick up the top three, flip the entire group, and put them back down on the remaining fourth pancake (which would stay where it is and remain unchanged). The new state of the stack would then be "++"­. The other valid ways would be to pick up and flip the top one, the top two, or all four. It would not be valid to choose and flip the middle two or the bottom one, for example; you can only take some number off the top.

## Goal

You will not serve the customer until every pancake is happy side up, but you don't want the pancakes to get cold, so you have to act fast! What is the smallest number of times you will need to execute the maneuver to get all the pancakes happy side up, if you make optimal choices?

---

## Components

The solution to this challenge has the following components:

- `pkg/ihop/pancakeflips.go` This file contains the pancake flipping function. The function returns the number of flips required to have all of the pancakes in a stack with the happy face side up
- `pkg/ihop/testdata` This directory has files that contain test cases
- `pkg/ihop/pancakeflips_test.go` This file is used to test the PancakeFlips function
- `cmd/main.go` A simple CLI that takes a string, representing a stack of pancakes, to be used as an argument to the PancakeFlips function

### Building The CLI Application

To build the command line application, from within the `./cmd/happyfaces` directory, do the following:

```sh

go build -o program .

```

### Running The CLI Application

An executable file called `program` will be found in the `./cmd/happyfaces` directory. To run the CLI application, from within the `./cmd/happyfaces` directory, type the following:

```sh
./program -s <stack-string>
```

For example:

```sh
./program -s "--+-"
```

The output will be the number of flips required to have all of the pancakes with the happy face side up. For example after running `./program -s "+--+"` this will produce the following output:

```sh
Number of Pancake Flips = 2
```

---

### Testing

To run tests for the `PancakeFlips` function, from within the `./pkg/ihop` directory, type the following:

```sh
go test
```

This will run all of the test cases contained in each of files found in the `pkg/ihop/testdata` directory. To add another set of test cases, create a file containing the desired test cases. Then place the file in the `pkg/ihop/testdata` directory.

Running `go test` will display all of the output from all test cases found in each file. The following is an example output with a single file containing test cases:

```sh
Test Filename: test1.txt
Case #1: 1
Case #2: 1
Case #3: 2
Case #4: 0
Case #5: 3
-----------
```

#### Test Cases File Format

The following describes the file format for a file containing test cases:

- The first line specifies the number of test cases contained in the file.
- Each subsequent line is then a unique test case. A test case line consists of string of characters with either a '+' (representing a happy face side) or a '-' (representing a blank side).

##### Test Cases File Example

```sh
5
-
-+
+-
+++
--+-
```

---

## References

- [Google Code Jam - Revenge of the Pancakes](docs/Revenge_of_the_Pancakes_-_Weave_Take_Home_Challenge.pdf)