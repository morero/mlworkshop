package main

import (
  "log"
  "math/rand"
  "github.com/sjwhitworth/golearn/base"
  "fmt"
  "github.com/sjwhitworth/golearn/filters"
 // "github.com/sjwhitworth/golearn/ensemble"
  //"github.com/sjwhitworth/golearn/evaluation"
  //"github.com/sjwhitworth/golearn/ensemble"
  //"github.com/sjwhitworth/golearn/evaluation"
  "github.com/sjwhitworth/golearn/ensemble"
  //"github.com/sjwhitworth/golearn/evaluation"
  "github.com/sjwhitworth/golearn/evaluation"
)
/*
import "fmt"
import (
  "math/rand"
  "github.com/sjwhitworth/golearn/base"
  "github.com/sjwhitworth/golearn/ensemble"
  "github.com/sjwhitworth/golearn/evaluation"
  "github.com/sjwhitworth/golearn/filters"
)
*/
func main() {
  //
  // Finally, Random Forests
  //

  log.Println("wewer")

  rand.Seed(44111342)

  // Load in the people dataset
  people, err := base.ParseCSVToInstances("data/people.csv", true)
  if err != nil {
    panic(err)
  }



    // Discretise the iris dataset with Chi-Merge
    filt := filters.NewChiMergeFilter(people, 0.999)
    for _, a := range base.NonClassFloatAttributes(people) {
      filt.AddAttribute(a)
    }

    filt.Train()
    peoplef := base.NewLazilyFilteredInstances(people, filt)

    // Create a 60-40 training-test split
    trainData, testData := base.InstancesTrainTestSplit(peoplef, 0.60)

    tree := ensemble.NewRandomForest(10, 3)

   err = tree.Fit(trainData)
    if err != nil {
      panic(err)
    }

    predictions, err := tree.Predict(testData)
    if err != nil {
      panic(err)
    }
    fmt.Println("RandomForest Performance")
    cf, err := evaluation.GetConfusionMatrix(testData, predictions)
    if err != nil {
      panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
    }
    fmt.Println(evaluation.GetSummary(cf))

  fmt.Println("test")
  /*
    */

}