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
	problems, err := leetcodeapi.GetAllProblems(rand.Intn(1000), 1)
	if err != nil {
		log.Fatal(err)
	}

	problem := problems.Problems[0]

	link := fmt.Sprintf("https://leetcode.com/problems/%s", problem.TitleSlug)
	fmt.Printf("Title: %s Difficulty: %s\n%s\n\n", problem.TitleSlug, problem.Difficulty, link)
	
	contents, err := leetcodeapi.GetProblemContentByTitleSlug(problems.Problems[0].TitleSlug)
	if err != nil {
		log.Fatal(err)
	}

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
