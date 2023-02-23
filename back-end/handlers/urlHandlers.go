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
	var msg *string
	ok, url := db.GetUrlByOriginalIfAvailable(body.Url, claims["usr"].(string))
	if ok && (*url != model.URL{}) {
		msg, _ = utils.ConvertResponseMessageToJson(constants.DUPLICATE_URL,
			utils.GetVariable("DOMAIN_NAME")+url.Short_url)
		return c.JSON(http.StatusOK, *msg)
	}
	short := utils.GenerateShortUrl()
	newUrl := model.URL{Short_url: short,
		Original_url: body.Url, Title: body.Title}
	user, _ := db.GetUserByUsername(claims["usr"].(string))
	newUrl.User_id = user.User_Id
	db.InsertNewUrl(newUrl)
	msg, _ = utils.ConvertResponseMessageToJson(constants.URL_SHORTENED,
		utils.GetVariable("DOMAIN_NAME")+"api/"+short)
	return c.String(http.StatusOK, *msg)
}

func HandleRedirects(c echo.Context) error {
	shortUrl := c.Param("s")
	cookie, _ := c.Cookie("token")
	claims, _ := utils.ExtractTokenClaimsFromCookie(*cookie)
	var msg *string
	ok, url := db.GetUrlByShortIfAvailable(shortUrl, claims["usr"].(string))
	if !ok {
		msg, _ = utils.ConvertResponseMessageToJson(constants.UNAVAILABLE_LINK, "")
		return c.String(http.StatusNotFound, *msg)
	}
	if dateExtractor.IsLinkExpired(url.Created_at) {
		msg, _ = utils.ConvertResponseMessageToJson(constants.EXPIRED_LINK, "")
		return c.String(http.StatusForbidden, *msg)
	}
	return c.Redirect(http.StatusSeeOther, url.Original_url)
}

func GetUserUrls(c echo.Context) error {
	token, _ := c.Cookie("token")
	var msg *string
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("mh_secret"), nil
	})
	if err != nil {
		msg, _ = utils.ConvertResponseMessageToJson(constants.UNAUTHORIZED_USER, "")
		return c.String(http.StatusUnauthorized, *msg)
	}
	rows, err := db.GetUrlsByUsername(claims["usr"].(string))
	return c.JSON(http.StatusOK, rows)
}
