package main

import (
	"testing"
)

func TestAdv(t *testing.T) {
    // Test case 1
    registers := map[string]int{"A": 16, "B": 0, "C": 0}
    operand := 2
    expected := 4

    adv(operand, registers)

    if registers["A"] != expected {
        t.Errorf("adv(%d, registers) = %d; want %d", operand, registers["A"], expected)
    }

    // Test case 2
    registers = map[string]int{"A": 32, "B": 0, "C": 0}
    operand = 3
    expected = 4

    adv(operand, registers)

    if registers["A"] != expected {
        t.Errorf("adv(%d, registers) = %d; want %d", operand, registers["A"], expected)
    }

    // Test case 3
    registers = map[string]int{"A": 64, "B": 0, "C": 0}
    operand = 1
    expected = 32

    adv(operand, registers)

    if registers["A"] != expected {
        t.Errorf("adv(%d, registers) = %d; want %d", operand, registers["A"], expected)
    }
}

func TestBxl(t *testing.T) {
	// Test case 1
	registers := map[string]int{"A": 0, "B": 5, "C": 0}
	operand := 3
	expected := 6 // 5 XOR 3 = 6

	bxl(operand, registers)

	if registers["B"] != expected {
		t.Errorf("bxl(%d, registers) = %d; want %d", operand, registers["B"], expected)
	}

	// Test case 2
	registers = map[string]int{"A": 0, "B": 10, "C": 0}
	operand = 7
	expected = 13 // 10 XOR 7 = 13

	bxl(operand, registers)

	if registers["B"] != expected {
		t.Errorf("bxl(%d, registers) = %d; want %d", operand, registers["B"], expected)
	}

	// Test case 3
	registers = map[string]int{"A": 0, "B": 15, "C": 0}
	operand = 15
	expected = 0 // 15 XOR 15 = 0

	bxl(operand, registers)

	if registers["B"] != expected {
		t.Errorf("bxl(%d, registers) = %d; want %d", operand, registers["B"], expected)
	}
}

func TestBst(t *testing.T){
	// Test case 1: Combo operand 6
    registers := map[string]int{"A": 0, "B": 0, "C": 9}
    operand := 6
    expected := 1 // 9 % 8 = 1

    bst(operand, registers)

    if registers["B"] != expected {
        t.Errorf("bst(%d, registers) = %d; want %d", operand, registers["B"], expected)
    }

    // Test case 2: Combo operand 15 (literal value 15)
	registers = map[string]int{"A": 0, "B": 0, "C": 0}
    operand = 0
    expected = 0 // 0 % 8 = 0

    bst(operand, registers)

    if registers["B"] != expected {
        t.Errorf("bst(%d, registers) = %d; want %d", operand, registers["B"], expected)
    }

    // // Test case 3: Combo operand 8 (literal value 8)
    registers = map[string]int{"A": 0, "B": 0, "C": 0}
    operand = 3
    expected = 3 // 3 % 8 = 3

    bst(operand, registers)

    if registers["B"] != expected {
        t.Errorf("bst(%d, registers) = %d; want %d", operand, registers["B"], expected)
    }

    // // Test case 4: Combo operand 4 (value of register A)
    registers = map[string]int{"A": 10, "B": 0, "C": 0}
    operand = 4
    expected = 2 // 10 % 8 = 2

    bst(operand, registers)

    if registers["B"] != expected {
        t.Errorf("bst(%d, registers) = %d; want %d", operand, registers["B"], expected)
    }

    // Test case 5: Combo operand 5 (value of register B)
    registers = map[string]int{"A": 0, "B": 15, "C": 0}
    operand = 5
    expected = 7 // 15 % 8 = 7

    bst(operand, registers)

    if registers["B"] != expected {
        t.Errorf("bst(%d, registers) = %d; want %d", operand, registers["B"], expected)
    }

	// Test case 5: Combo operand 6 (value of register C)
    registers = map[string]int{"A": 0, "B": 0, "C": 8}
    operand = 6
    expected = 0 // 8 % 8 = 0

    bst(operand, registers)

    if registers["B"] != expected {
        t.Errorf("bst(%d, registers) = %d; want %d", operand, registers["B"], expected)
    }

    // Test case 6: Combo operand 7 (reserved, should not appear in valid programs)
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("bst(7, registers) did not panic")
        }
    }()
    registers = map[string]int{"A": 0, "B": 0, "C": 0}
    operand = 7

    bst(operand, registers)
}


func TestJnz(t *testing.T){
	// Test case 1: Register A is 0, jnz should return false
    registers := map[string]int{"A": 0, "B": 0, "C": 0}
    operand := 5
    expectedBool := false

    _, resultBool := jnz(operand, registers)

    if resultBool != expectedBool {
        t.Errorf("jnz(%d, registers) = %v; want %v", operand, resultBool, expectedBool)
    }

	// Test case 2: Register A is not 0, jnz should return true
    registers = map[string]int{"A": 1, "B": 0, "C": 0}
    operand = 5
    expectedBool = true

    _, resultBool = jnz(operand, registers)

    if resultBool != expectedBool {
        t.Errorf("jnz(%d, registers) = %v; want %v", operand, resultBool, expectedBool)
    }
}

func TestBxc(t *testing.T){
	// Test case 1: Registers B and C have different values
    registers := map[string]int{"A": 0, "B": 5, "C": 3}
    operand := 0 // Operand is ignored
    expectedB := 5 ^ 3 // 5 XOR 3

    bxc(operand, registers)

    if registers["B"] != expectedB {
        t.Errorf("bxc(%d, registers) = %d; want %d", operand, registers["B"], expectedB)
    }

    // Test case 2: Registers B and C have the same values
    registers = map[string]int{"A": 0, "B": 7, "C": 7}
    operand = 0 // Operand is ignored
    expectedB = 7 ^ 7 // 7 XOR 7

    bxc(operand, registers)

    if registers["B"] != expectedB {
        t.Errorf("bxc(%d, registers) = %d; want %d", operand, registers["B"], expectedB)
    }

    // Test case 3: Register B is 0
    registers = map[string]int{"A": 0, "B": 0, "C": 4}
    operand = 0 // Operand is ignored
    expectedB = 0 ^ 4 // 0 XOR 4

    bxc(operand, registers)

    if registers["B"] != expectedB {
        t.Errorf("bxc(%d, registers) = %d; want %d", operand, registers["B"], expectedB)
    }

    // Test case 4: Register C is 0
    registers = map[string]int{"A": 0, "B": 6, "C": 0}
    operand = 0 // Operand is ignored
    expectedB = 6 ^ 0 // 6 XOR 0

    bxc(operand, registers)

    if registers["B"] != expectedB {
        t.Errorf("bxc(%d, registers) = %d; want %d", operand, registers["B"], expectedB)
    }
}

func TestOut(t *testing.T){
		// Test case 1: Operand is less than 8
		registers := map[string]int{"A": 0, "B": 8, "C": 0}
		operand := 5
		expectedOutput := "0"
	
		result,_ := out(operand, registers)
	
		if result != expectedOutput {
			t.Errorf("out(%d) = %s; want %s", operand, result, expectedOutput)
		}
}

func TestBDV(t *testing.T) {
    // Reset the registers before starting the test
    registers := map[string]int{"A": 10, "B": 0, "C": 0}
    operand := 2 // The divisor

    // Expected result: B = A / operand = 10 / 4 = 2.5 -> 2
    expectedB := 2

    // Call the bdv function
    _, _ = bdv(operand, registers)

    if registers["B"] != expectedB {
        t.Errorf("bdv(%d) = %d; want %d", operand, registers["B"], expectedB)
    }
}

func TestCDV(t *testing.T) {
    // Reset the registers before starting the test
    registers := map[string]int{"A": 10, "B": 3, "C": 0}
    operand := 5 // The divisor

    // Expected result: C = A / operand = 10 / 2^1
    expectedC := 1

    // Call the cdv function
    _, _ = cdv(operand, registers)

    if registers["C"] != expectedC {
        t.Errorf("cdv(%d) = %d; want %d", operand, registers["C"], expectedC)
    }
}