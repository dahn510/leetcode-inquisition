package main

import (
	"fmt"
	"log"
	"math/rand"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/charmbracelet/glamour"
	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

func main() {
	// get non-premium problems
	var problem leetcodeapi.Problem
	for {
		problems, err := leetcodeapi.GetAllProblems(rand.Intn(1000), 1)
		if err != nil {
			log.Fatal(err)
		}
		problem = problems.Problems[0]
		if !problem.PaidOnly {
			break
		}
	}

	link := fmt.Sprintf("https://leetcode.com/problems/%s", problem.TitleSlug)
	fmt.Printf("Title: %s\nDifficulty: %s\n%s\n\n", problem.Title, problem.Difficulty, link)
	
	contents, err := leetcodeapi.GetProblemContentByTitleSlug(problem.TitleSlug)
	if err != nil {
		log.Fatal(err)
	}

	// convert problem in html to markdown
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(contents.Content)
	if err != nil {
		log.Fatal(err)
	}

	out, err := glamour.Render(markdown, "dark")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(out)
	return
}
