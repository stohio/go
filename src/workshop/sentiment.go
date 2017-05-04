package main
import "github.com/cdipaolo/sentiment"
import "fmt"

func main(){
	model, err := sentiment.Restore()
	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}
	analysis := model.SentimentAnalysis("Your mother is an awful lady", sentiment.English) // 0
	fmt.Print(analysis);
}
