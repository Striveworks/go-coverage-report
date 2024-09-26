package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReport_Markdown(t *testing.T) {
	oldCov, err := ParseCoverage("testdata/01-old-coverage.txt")
	require.NoError(t, err)

	newCov, err := ParseCoverage("testdata/01-new-coverage.txt")
	require.NoError(t, err)

	changedFiles, err := ParseChangedFiles("testdata/01-changed-files.json", "github.com/fgrosse/prioqueue")
	require.NoError(t, err)

	report := NewReport(oldCov, newCov, changedFiles)
	actual := report.Markdown()

	expected := `### Merging this branch will **decrease** overall coverage
#### Overall Project Coverage: 100.00% :arrow_right: 90.20%
| Impacted Packages | Coverage Δ | :robot: |
|-------------------|------------|---------|
| github.com/fgrosse/prioqueue | 90.20% (**-9.80%**) | :thumbsdown: |
| github.com/fgrosse/prioqueue/foo/bar | 0.00% (ø) |  |

---

<details>

<summary>Coverage by file</summary>

### Changed files (no unit tests)

| Changed File | Coverage Δ | Total | Covered | Missed | :robot: |
|--------------|------------|-------|---------|--------|---------|
| github.com/fgrosse/prioqueue/foo/bar/baz.go | 0.00% (ø) | 0 | 0 | 0 |  |
| github.com/fgrosse/prioqueue/min_heap.go | 80.77% (**-19.23%**) | 52 (+2) | 42 (-8) | 10 (+10) | :skull:  |

_Please note that the "Total", "Covered", and "Missed" counts above refer to ***code statements*** instead of lines of code. The value in brackets refers to the test coverage of that file in the old version of the code._

</details>`
	assert.Equal(t, expected, actual)
}

func TestReport_Markdown_OnlyChangedUnitTests(t *testing.T) {
	oldCov, err := ParseCoverage("testdata/02-old-coverage.txt")
	require.NoError(t, err)

	newCov, err := ParseCoverage("testdata/02-new-coverage.txt")
	require.NoError(t, err)

	changedFiles, err := ParseChangedFiles("testdata/02-changed-files.json", "github.com/fgrosse/prioqueue")
	require.NoError(t, err)

	report := NewReport(oldCov, newCov, changedFiles)
	actual := report.Markdown()

	expected := `### Merging this branch will **increase** overall coverage
#### Overall Project Coverage: 90.20% :arrow_right: 99.02%
| Impacted Packages | Coverage Δ | :robot: |
|-------------------|------------|---------|
| github.com/fgrosse/prioqueue | 99.02% (**+8.82%**) | :thumbsup: |

---

<details>

<summary>Coverage by file</summary>

### Changed unit test files

- github.com/fgrosse/prioqueue/min_heap_test.go

</details>`
	assert.Equal(t, expected, actual)
}

func TestReport_Markdown_Realistic(t *testing.T) {
	oldCov, err := ParseCoverage("testdata/03-old-coverage.txt")
	require.NoError(t, err)

	newCov, err := ParseCoverage("testdata/03-new-coverage.txt")
	require.NoError(t, err)

	changedFiles, err := ParseChangedFiles("testdata/03-changed-files.json", "chariot/go/apps/training-v2/")
	require.NoError(t, err)

	report := NewReport(oldCov, newCov, changedFiles)
	actual := report.Markdown()

	expected := `### Merging this branch will **decrease** overall coverage
#### Overall Project Coverage: 55.76% :arrow_right: 54.99%
| Impacted Packages | Coverage Δ | :robot: |
|-------------------|------------|---------|
| chariot/go/apps/training-v2/internal/api/handlers | 9.37% (**-0.09%**) | :thumbsdown: |

---

<details>

<summary>Coverage by file</summary>

### Changed files (no unit tests)

| Changed File | Coverage Δ | Total | Covered | Missed | :robot: |
|--------------|------------|-------|---------|--------|---------|
| chariot/go/apps/training-v2/internal/api/handlers/blueprints.go | 9.00% (ø) | 3779 | 340 | 3439 |  |

_Please note that the "Total", "Covered", and "Missed" counts above refer to ***code statements*** instead of lines of code. The value in brackets refers to the test coverage of that file in the old version of the code._

</details>`
	assert.Equal(t, expected, actual)
}

func TestOverallCoveragePercentage(t *testing.T) {
	pp, err := ParseProfiles("testdata/03-old-coverage.txt")
	require.NoError(t, err)
	total, _ := getOverallCoveragePercent(pp)
	require.NoError(t, err)
	assert.Equal(t, 55.75731681264687, total)
}

func TestOverallCoverageWithIgnorePercentage(t *testing.T) {
	pp, err := ParseProfiles("testdata/03-old-coverage.txt")
	require.NoError(t, err)
	total, _ := getOverallCoveragePercent(pp, "*mock.go", "testingdeps.go")
	require.NoError(t, err)
	assert.Equal(t, 56.108083560399635, total)
}
