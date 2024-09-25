package main

import (
	"errors"
	"fmt"
	"path/filepath"
)

func getOverallCoveragePercent(profiles []*Profile, ignorePatterns ...string) (float64, error) {
	var (
		totalStatements   int
		coveredStatements int
	)

	if len(ignorePatterns) > 0 {
		fmt.Printf("ignoring files from overall coverage: %s\n", ignorePatterns)
	}
	for _, profile := range profiles {
		if ignoreFile(profile.FileName, ignorePatterns) {
			continue
		}
		for _, block := range profile.Blocks {
			totalStatements += block.NumStmt
			if block.Count > 0 { // count only if the block was hit
				coveredStatements += block.NumStmt
			}
		}
	}

	if totalStatements == 0 {
		return 0, errors.New("no statements found in coverage data")
	}

	coveragePercentage := float64(coveredStatements) / float64(totalStatements) * 100

	fmt.Printf("Total Statements: %d\n", totalStatements)
	fmt.Printf("Covered Statements: %d\n", coveredStatements)
	fmt.Printf("Coverage Percentage: %.2f%%\n", coveragePercentage)
	return coveragePercentage, nil
}

// ignoreFile checks if a file name matches any of the user-specified ignore patterns
func ignoreFile(fileName string, ignorePatterns []string) bool {
	if len(ignorePatterns) > 0 {
	}
	for _, pattern := range ignorePatterns {
		match, err := filepath.Match(pattern, filepath.Base(fileName))
		if err != nil {
			continue
		}
		if match {
			return true
		}
	}
	return false
}
