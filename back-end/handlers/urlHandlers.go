package handlers

import (
	"back-end/constants"
	"back-end/dateExtractor"
	"back-end/db"
	"back-end/model"
	"back-end/utils"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
	"net/http"
)

func HandleShortener(c echo.Context) error {
	body := model.Input{}
	cookie, _ := c.Cookie("token")
	claims, _ := utils.ExtractTokenClaimsFromCookie(*cookie)
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		fmt.Println("no json passed to handler")
	}
	var msg model.ResponseMessage
	ok, url := db.GetUrlByOriginalIfAvailable(body.Url, claims["usr"].(string))
	if ok && (*url != model.URL{}) {
		msg = utils.GenerateResponseMessage(constants.DUPLICATE_URL,
			utils.GetVariable("DOMAIN_NAME")+url.Short_url)
		return c.JSON(http.StatusOK, msg)
	}
	short := utils.GenerateShortUrl()
	newUrl := model.URL{Short_url: short,
		Original_url: body.Url, Title: body.Title}
	user, _ := db.GetUserByUsername(claims["usr"].(string))
	newUrl.User_id = user.User_Id
	db.InsertNewUrl(newUrl)
	msg = utils.GenerateResponseMessage(constants.URL_SHORTENED,
		utils.GetVariable("DOMAIN_NAME")+"api/"+short)
	return c.JSON(http.StatusOK, msg)
}

func HandleRedirects(c echo.Context) error {
	shortUrl := c.Param("s")
	cookie, _ := c.Cookie("token")
	claims, _ := utils.ExtractTokenClaimsFromCookie(*cookie)
	var msg model.ResponseMessage
	ok, url := db.GetUrlByShortIfAvailable(shortUrl, claims["usr"].(string))
	if !ok {
		msg = utils.GenerateResponseMessage(constants.UNAVAILABLE_LINK, "")
		return c.JSON(http.StatusNotFound, msg)
	}
	if dateExtractor.IsLinkExpired(url.Created_at) {
		msg = utils.GenerateResponseMessage(constants.EXPIRED_LINK, "")
		return c.JSON(http.StatusForbidden, msg)
	}
	return c.Redirect(http.StatusSeeOther, url.Original_url)
}

func GetUserUrls(c echo.Context) error {
	token, _ := c.Cookie("token")
	var msg model.ResponseMessage
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("mh_secret"), nil
	})
	if err != nil {
		msg = utils.GenerateResponseMessage(constants.UNAUTHORIZED_USER, "")
		return c.JSON(http.StatusUnauthorized, msg)
	}
	rows, err := db.GetUrlsByUsername(claims["usr"].(string))
	return c.JSON(http.StatusOK, rows)
}
