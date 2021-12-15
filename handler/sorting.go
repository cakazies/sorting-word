package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"sorting-word/helper"
	utils "sorting-word/utils/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func SortingMap(c *gin.Context) {

	var form FormRequest
	var listResp []Item

	err := c.ShouldBind(&form)
	fmt.Println("err ", err)
	if err != nil {
		utils.Response(c, helper.StatusFailed, http.StatusBadRequest, err)
		return
	}

	if form.Paragraf == "" {
		utils.Response(c, helper.StatusFailed, http.StatusBadRequest, "please set your input paragraf")
		return
	}

	text := form.Paragraf
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	processedString := strings.ToLower(reg.ReplaceAllString(text, " "))
	listWord := strings.Split(processedString, " ")

	mapWord := make(map[string]int)

	for _, v := range listWord {

		if val, ok := mapWord[v]; !ok {
			mapWord[v] = 1
		} else {
			mapWord[v] = (val + 1)
		}

	}

	p := make(SortList, len(mapWord))

	i := 0
	for k, v := range mapWord {
		p[i] = Item{k, v}
		i++
	}

	sort.Sort(p)

	i = 0
	for _, k := range p {
		listResp = append(listResp, k)
		i++
		if i == 10 {
			break
		}
	}

	utils.Response(c, helper.StatusOk, http.StatusOK, listResp)

}
