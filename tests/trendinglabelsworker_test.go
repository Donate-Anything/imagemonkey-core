package tests

import (
	"testing"
	//"../src/datastructures"
	"os/exec"
	"time"
)

func runTrendingLabelsWorker(t *testing.T) {
	// Start a process
	cmd := exec.Command("go", "run", "trendinglabelsworker.go", "api_secrets.go", "shared_secrets.go", "-singleshot=true", "-treshold=5", "-use_github=false")
	cmd.Dir = "../src"
	err := cmd.Start()
	ok(t, err)

	// Wait for the process to finish or kill it after a timeout:
	done := make(chan error, 1)
	go func() {
	    done <- cmd.Wait()
	}()
	select {
	case <-time.After(5 * time.Second):
	    err := cmd.Process.Kill()
	    ok(t, err) //failed to kill process
	    t.Errorf("process killed as timeout reached")
	case err := <-done:
	    ok(t, err)
	}
}

func runMakeTrendingLabelsProductiveScript(t *testing.T, trendingLabel string, renameTo string) {
	// Start a process
	cmd := exec.Command("go", "run", "make_labels_productive.go", "api_secrets.go", "shared_secrets.go", 
						"-dryrun=false", "-trendinglabel", trendingLabel, "-renameto", renameTo, "-autoclose=false")
	cmd.Dir = "../src"
	err := cmd.Start()
	ok(t, err)

	// Wait for the process to finish or kill it after a timeout:
	done := make(chan error, 1)
	go func() {
	    done <- cmd.Wait()
	}()
	select {
	case <-time.After(5 * time.Second):
	    err := cmd.Process.Kill()
	    ok(t, err) //failed to kill process
	    t.Errorf("process killed as timeout reached")
	case err := <-done:
	    ok(t, err)
	}
}

func TestBasicTrendingLabelsWorkerFunctionalityLabelsAlreadyExists(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	testSignUp(t, "testuser", "testpassword", "testuser@imagemonkey.io")
	token := testLogin(t, "testuser", "testpassword", 200)

	testMultipleDonate(t, "apple")

	imageIds, err := db.GetAllImageIds()
	ok(t, err)

	for _, imageId := range imageIds {
		testSuggestLabelForImage(t, imageId, "red apple", token)
	}
	runTrendingLabelsWorker(t)

	numWithLabelsBefore, err := db.GetNumberOfImagesWithLabel("apple")
	ok(t, err)
	equals(t, int(numWithLabelsBefore), int(13))

	numBefore, err := db.GetNumberOfImagesWithLabelSuggestions("red apple")
	ok(t, err)
	equals(t, int(numBefore), int(13))
	runMakeTrendingLabelsProductiveScript(t, "red apple", "apple")
	numAfter, err := db.GetNumberOfImagesWithLabelSuggestions("red apple")
	ok(t, err)
	equals(t, int(numAfter), int(0))

	numWithLabelsAfter, err := db.GetNumberOfImagesWithLabel("apple")
	ok(t, err)
	equals(t, int(numWithLabelsAfter), int(13))
}


func TestBasicTrendingLabelsWorkerFunctionality(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	testSignUp(t, "testuser", "testpassword", "testuser@imagemonkey.io")
	token := testLogin(t, "testuser", "testpassword", 200)

	testMultipleDonate(t, "floor")

	imageIds, err := db.GetAllImageIds()
	ok(t, err)

	for _, imageId := range imageIds {
		testSuggestLabelForImage(t, imageId, "red apple", token)
	}
	runTrendingLabelsWorker(t)

	numWithLabelsBefore, err := db.GetNumberOfImagesWithLabel("apple")
	ok(t, err)
	equals(t, int(numWithLabelsBefore), int(0))

	numBefore, err := db.GetNumberOfImagesWithLabelSuggestions("red apple")
	ok(t, err)
	equals(t, int(numBefore), int(13))

	productiveLabelIdsBefore, err := db.GetProductiveLabelIdsForTrendingLabels()
	ok(t, err)
	equals(t, int(0), int(len(productiveLabelIdsBefore)))

	numOfTrendingLabelSuggestionsBefore, err := db.GetNumberOfTrendingLabelSuggestions()
	ok(t, err)
	equals(t, int(numOfTrendingLabelSuggestionsBefore), int(1))


	runMakeTrendingLabelsProductiveScript(t, "red apple", "apple")
	numAfter, err := db.GetNumberOfImagesWithLabelSuggestions("red apple")
	ok(t, err)
	equals(t, int(numAfter), int(0))

	numWithLabelsAfter, err := db.GetNumberOfImagesWithLabel("apple")
	ok(t, err)
	equals(t, int(numWithLabelsAfter), int(13))

	productiveLabelIdsAfter, err := db.GetProductiveLabelIdsForTrendingLabels()
	ok(t, err)
	equals(t, int(1), int(len(productiveLabelIdsAfter)))

	expectedLabelId, err := db.GetLabelIdFromName("apple")
	ok(t, err)
	equals(t, productiveLabelIdsAfter[0], expectedLabelId)
}

func TestBasicTrendingLabelsWorkerFunctionality2(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	testSignUp(t, "testuser", "testpassword", "testuser@imagemonkey.io")
	token := testLogin(t, "testuser", "testpassword", 200)

	testDonate(t, "./images/apples/apple1.jpeg", "apple", false, "", "")
	testDonate(t, "./images/apples/apple2.jpeg", "apple", false, "", "")

	imageIds, err := db.GetAllImageIds()
	ok(t, err)

	testSuggestLabelForImage(t, imageIds[0], "wooden floor", token)
	testSuggestLabelForImage(t, imageIds[1], "dirty floor", token)

	runTrendingLabelsWorker(t)

	numWithLabelsBefore, err := db.GetNumberOfImagesWithLabel("apple")
	ok(t, err)
	equals(t, int(numWithLabelsBefore), int(2))

	numBefore, err := db.GetNumberOfImagesWithLabelSuggestions("wooden floor")
	ok(t, err)
	equals(t, int(numBefore), int(1))

	numBefore2, err := db.GetNumberOfImagesWithLabel("floor")
	ok(t, err)
	equals(t, int(numBefore2), int(0))


	runMakeTrendingLabelsProductiveScript(t, "wooden floor", "floor")
	numAfter, err := db.GetNumberOfImagesWithLabelSuggestions("red apple")
	ok(t, err)
	equals(t, int(numAfter), int(0))

	numAfter2, err := db.GetNumberOfImagesWithLabelSuggestions("dirty floor")
	ok(t, err)
	equals(t, int(numAfter2), int(1))

	numAfter3, err := db.GetNumberOfImagesWithLabelSuggestions("wooden floor")
	ok(t, err)
	equals(t, int(numAfter3), int(0))

	numWithLabelsAfter, err := db.GetNumberOfImagesWithLabel("apple")
	ok(t, err)
	equals(t, int(numWithLabelsAfter), int(2))

	numWithLabelsAfter2, err := db.GetNumberOfImagesWithLabel("floor")
	ok(t, err)
	equals(t, int(numWithLabelsAfter2), int(1))
}

func TestBasicTrendingLabelsWorkerFunctionalityRecurringLabelSuggestion(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	testSignUp(t, "testuser", "testpassword", "testuser@imagemonkey.io")
	token := testLogin(t, "testuser", "testpassword", 200)

	testMultipleDonate(t, "floor")

	imageIds, err := db.GetAllImageIds()
	ok(t, err)

	for i, imageId := range imageIds {
		if i == 0 {
			continue
		}
		testSuggestLabelForImage(t, imageId, "red apple", token)
	}
	runTrendingLabelsWorker(t)

	numWithLabelsBefore, err := db.GetNumberOfImagesWithLabel("apple")
	ok(t, err)
	equals(t, int(numWithLabelsBefore), int(0))

	numBefore, err := db.GetNumberOfImagesWithLabelSuggestions("red apple")
	ok(t, err)
	equals(t, int(numBefore), int(12))

	productiveLabelIdsBefore, err := db.GetProductiveLabelIdsForTrendingLabels()
	ok(t, err)
	equals(t, int(0), int(len(productiveLabelIdsBefore)))

	numOfTrendingLabelSuggestionsBefore, err := db.GetNumberOfTrendingLabelSuggestions()
	ok(t, err)
	equals(t, int(numOfTrendingLabelSuggestionsBefore), int(1))

	runMakeTrendingLabelsProductiveScript(t, "red apple", "apple")
	numAfter, err := db.GetNumberOfImagesWithLabelSuggestions("red apple")
	ok(t, err)
	equals(t, int(numAfter), int(0))

	numWithLabelsAfter, err := db.GetNumberOfImagesWithLabel("apple")
	ok(t, err)
	equals(t, int(numWithLabelsAfter), int(12))

	productiveLabelIdsAfter, err := db.GetProductiveLabelIdsForTrendingLabels()
	ok(t, err)
	equals(t, int(1), int(len(productiveLabelIdsAfter)))

	expectedLabelId, err := db.GetLabelIdFromName("apple")
	ok(t, err)
	equals(t, productiveLabelIdsAfter[0], expectedLabelId)


	testSuggestLabelForImage(t, imageIds[0], "red apple", token)
	recurringLabelSuggestionNumBefore, err := db.GetNumberOfImagesWithLabelSuggestions("red apple")
	ok(t, err)
	equals(t, int(recurringLabelSuggestionNumBefore), int(1))

	runTrendingLabelsWorker(t)

	recurringLabelSuggestionNumAfter, err := db.GetNumberOfImagesWithLabelSuggestions("red apple")
	ok(t, err)
	equals(t, int(recurringLabelSuggestionNumAfter), int(0))

	numLabelsAfterRecurringLabelSuggestion, err := db.GetNumberOfImagesWithLabel("apple")
	ok(t, err)
	equals(t, int(numLabelsAfterRecurringLabelSuggestion), int(13))
}


func TestBasicTrendingLabelsWorkerFunctionalityNumOfSent(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	testSignUp(t, "testuser", "testpassword", "testuser@imagemonkey.io")
	token := testLogin(t, "testuser", "testpassword", 200)

	testMultipleDonate(t, "floor")

	imageIds, err := db.GetAllImageIds()
	ok(t, err)

	for i, imageId := range imageIds {
		if i == 0 {
			continue
		}
		testSuggestLabelForImage(t, imageId, "red apple", token)
	}
	runTrendingLabelsWorker(t)

	numWithLabelsBefore, err := db.GetNumberOfImagesWithLabel("apple")
	ok(t, err)
	equals(t, int(numWithLabelsBefore), int(0))

	numBefore, err := db.GetNumberOfImagesWithLabelSuggestions("red apple")
	ok(t, err)
	equals(t, int(numBefore), int(12))

	productiveLabelIdsBefore, err := db.GetProductiveLabelIdsForTrendingLabels()
	ok(t, err)
	equals(t, int(0), int(len(productiveLabelIdsBefore)))

	numOfTrendingLabelSuggestionsBefore, err := db.GetNumberOfTrendingLabelSuggestions()
	ok(t, err)
	equals(t, int(numOfTrendingLabelSuggestionsBefore), int(1))

	numOfTrendingLabelSentBefore, err := db.GetNumOfSentOfTrendingLabel("red apple")
	ok(t, err)
	equals(t, int(numOfTrendingLabelSentBefore), int(0))

	runTrendingLabelsWorker(t)

	numOfTrendingLabelSentAfter, err := db.GetNumOfSentOfTrendingLabel("red apple")
	ok(t, err)
	equals(t, int(numOfTrendingLabelSentAfter), int(12))


	runMakeTrendingLabelsProductiveScript(t, "red apple", "apple")
	numAfter, err := db.GetNumberOfImagesWithLabelSuggestions("red apple")
	ok(t, err)
	equals(t, int(numAfter), int(0))

	numWithLabelsAfter, err := db.GetNumberOfImagesWithLabel("apple")
	ok(t, err)
	equals(t, int(numWithLabelsAfter), int(12))

	productiveLabelIdsAfter, err := db.GetProductiveLabelIdsForTrendingLabels()
	ok(t, err)
	equals(t, int(1), int(len(productiveLabelIdsAfter)))

	expectedLabelId, err := db.GetLabelIdFromName("apple")
	ok(t, err)
	equals(t, productiveLabelIdsAfter[0], expectedLabelId)

	runTrendingLabelsWorker(t)

	numOfTrendingLabelSentAfterRunAgain, err := db.GetNumOfSentOfTrendingLabel("red apple")
	ok(t, err)
	equals(t, int(numOfTrendingLabelSentAfterRunAgain), int(12))
}