package main

import (
	"errors"
	"path/filepath"
)

func getOverallCoveragePercent(profiles []*Profile, ignorePatterns ...string) (float64, error) {
	var (
		totalStatements   int
		coveredStatements int
	)

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

	return coveragePercentage, nil
}

// ignoreFile checks if a file name matches any of the user-specified ignore patterns
func ignoreFile(fileName string, ignorePatterns []string) bool {
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
