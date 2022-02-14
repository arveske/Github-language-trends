package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	TotalCount int64            `json:"total_count"`
	Title      bool             `json:"incomplete_results"`
	Items      []RepositoryInfo `json:"items"`
}

type RepositoryInfo struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	HtmlUrl  string `json:"html_url"`
	Language string `json:"language"`
}

type LanguageInfo struct {
	Count   int64    `json:"repository_count"`
	RepList []string `json:"repository_list"`
}

func main() {
	router := gin.Default()
	router.GET("/languages", getLanguages)
	router.GET("/languages/:name", getLanguageByName)

	router.Run("localhost:8080")
}

func getLanguages(c *gin.Context) {
	apiResponse, err := getRepositoryList()
	if err == nil {
		languageList := getLanguageList(apiResponse)
		c.IndentedJSON(http.StatusOK, languageList)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}
}

func getLanguageByName(c *gin.Context) {
	name := c.Param("name")
	apiResponse, err := getRepositoryList()
	if err == nil {
		languageList := getLanguageList(apiResponse)
		if _, ok := languageList[name]; ok {
			responseMap := make(map[string]LanguageInfo)
			responseMap[name] = languageList[name]
			c.IndentedJSON(http.StatusOK, responseMap)
		} else {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "language not found"})
		}
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}
}

func getDate30DaysAgo() string {
	return time.Now().AddDate(0, 0, -30).Format("2006-01-02")
}

func getRepositoryList() ([]RepositoryInfo, error) {
	url := fmt.Sprintf("https://api.github.com/search/repositories?q=created:>%v&sort=stars&order=desc&per_page=100",
		getDate30DaysAgo())

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseParsed Response
	errParsing := json.Unmarshal(body, &responseParsed)
	if errParsing != nil {
		return nil, errParsing
	}

	return responseParsed.Items, nil
}

func getLanguageList(repositories []RepositoryInfo) map[string]LanguageInfo {
	languageList := make(map[string]LanguageInfo)
	for _, resp := range repositories {
		if resp.Language != "" {
			if _, ok := languageList[resp.Language]; ok {
				languageData := languageList[resp.Language]
				languageData.Count += 1
				languageData.RepList = append(languageData.RepList, resp.HtmlUrl)
				languageList[resp.Language] = languageData
			} else {
				languageList[resp.Language] = LanguageInfo{1, []string{resp.HtmlUrl}}
			}
		}
	}
	return languageList
}
