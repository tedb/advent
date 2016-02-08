package advent

import (
	"bufio"
	"fmt"
	"strings"
)

// Advent19DistinctMoleculesFile reads data from a file in puzzle format
func Advent19DistinctMoleculesFile(arg string) (count int) {
	parts := strings.SplitN(arg, "\n\n", 2)
	swaps, mol := parts[0], parts[1]
	return Advent19DistinctMolecules(mol, swaps)
}

// Advent19DistinctMolecules takes a "molecule" like "HOH"
// and tries all the substring swaps listed in swaps,
// then returns the count of distinct outcomes.
// Basic algorithm: for each swap ("A => B"), recursively
// find the index of the first occurrence of A, compose a new string with B
// in its place, then call the function with
func Advent19DistinctMolecules(mol, swaps string) (count int) {
	var swapped []string

	scanner := bufio.NewScanner(strings.NewReader(swaps))
	for scanner.Scan() {
		var src, dest string
		fmt.Sscanf(scanner.Text(), "%s => %s", &src, &dest)
		swapped = append(swapped, MoleculeSwaps(mol, src, dest, 0)...)
	}

	err := scanner.Err()
	checkErr(err)

	return len(UniqueStrings(swapped))
}

// MoleculeSwaps takes one molecule and a src and dest, and returns
// all the versions of the molecule with the src swapped exactly once in each position.
// Runs recursively to process the string one replacement per function call
func MoleculeSwaps(mol, src, dest string, offset int) (swapped []string) {
	fmt.Println(mol, "offset", offset)
	if offset == len(mol) {
		return
	}
	parts := strings.SplitN(mol[offset:], src, 2)
	// return if we didn't find another match for src
	if len(parts) <= 1 {
		return
	}
	newString := mol[:offset] + parts[0] + dest + parts[1]
	//nextOffset := offset + len(parts[0]) + len(dest)
	//println("new string:", newString, "/ next offset:", nextOffset)
	fmt.Printf("%s(%s => %s) = %s\n", mol, src, dest, newString)
	swapped = append(swapped, newString)
	swapped = append(swapped, MoleculeSwaps(mol, src, dest, offset+1)...)
	return
}
