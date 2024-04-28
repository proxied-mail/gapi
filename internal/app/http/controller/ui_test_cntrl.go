package controller

import (
	"fmt"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"io/ioutil"
	"net/http"
	"strings"
)

type UiTestController struct {
	fx.In
	JbsRep repository.JobsRepository
}

func (jbCntrl UiTestController) Basic(c echo.Context) error {
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIzIiwianRpIjoiNGU4OGNkMmFhZjU1ODM5ZTA2ZWU0YzcyOGViMzY3NWNhNTQ5NTFmMTA3YTczMjBlZWJlMTY4ZTY5ZTFlZDUyNzllNGI4OTBiNGI1NmNhNzciLCJpYXQiOjE3MTQyNjUzNzYsIm5iZiI6MTcxNDI2NTM3NiwiZXhwIjoxNzQ1ODAxMzc2LCJzdWIiOiIxMzQ4MiIsInNjb3BlcyI6W119.O0DZ0-E-meypYlVfj6En-yKe5Yv57Nzgsoi1jOhgCSXStz790Zic27uED54cjuNLhYDJdeJmdBwWl19CI6c_yIlUSLEAkPH3HhKtKFTlo-s20Qmty8VNXPNN6cg09CMKr1pHSWXVndwIsEH4SYQ8OvkCmyddxlAm0FUj646rzeJy7avqXrzkWSU5G_P3dJBwIy5tCUhFMwzSLelwUFopNNV7e-_nsNT8BnL1JY524_OBDsTsNhuDAh3NBQY4HLVoPIanEiy38e3JuVRxYYoIOLlWR5Vi3vxxQQGevKUHUC6o6QjsM5bgoWhQs7Yk0sH-oqWJfGI3xXrLqHkiT9wG1gBlho-oqXiZFPbUThnU72QlwuRe3z6G_j3Bygu4qT4b_GV6Pp7emwj8iHz04RMl9uENBo1jnYf-Ski4FjUk_Q6QKynARO72OOU11rYZmhmqA57ngRQX222DmLhbA12UdWkXgxQoeGxOhoRbf6X6D5PcQHvo3lPpz89_BvW1a-0v3Qdd_BVfelb2-4eOpHRuSPjyBXR1iWmt55xcf9V6sEC6Ijhmd2E6yjUE23YBRjJQSNd5mpl-JDMVYC5sgV9HN6E-wKKf0Ylnxap8b1Rcb7BbM2jnGhKC4UqWRiX1p9lSSNKWhiQ39sWneuv9kOTOnnuPLdiqeUKZbdyJRWPTijY"
	PostData := strings.NewReader("")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://proxiedmail.com/en/board", PostData)
	req.Header.Set("Cookie", "token="+token)
	resp, err := client.Do(req)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error = %s \n", err)
	}

	// Print response
	fmt.Printf("Response = %s", string(data))

	return c.String(http.StatusOK, "OK")
}
