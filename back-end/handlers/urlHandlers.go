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
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		fmt.Println("no json passed to handler")
	}
	var msg *string
	if db.IsOriginalUrlAvailable(body.IntendedUrl, "mh") {
		url, _ := db.GetUrlByOriginal(body.IntendedUrl, "mh")
		msg, _ = utils.ConvertResponseMessageToJson(constants.DUPLICATE_URL,
			utils.GetVariable("DOMAIN_NAME")+url.Short_url)
		return c.String(http.StatusOK, *msg)
	}
	short := utils.GenerateShortUrl()
	claims, error := utils.ExtractTokenClaimsFromCookie(*cookie)
	user := &model.User{}
	if error == nil {
		user, err = db.GetUserByUsername(claims["usr"].(string))
	}
	db.InsertNewUrl(model.URL{Short_url: short,
		Original_url: body.IntendedUrl, Name: body.Name,
		User_id: user.User_Id})
	msg, _ = utils.ConvertResponseMessageToJson(constants.URL_SHORTENED,
		utils.GetVariable("DOMAIN_NAME")+short)
	return c.String(http.StatusOK, *msg)
}

func HandleRedirects(c echo.Context) error {
	shortUrl := c.Param("s")
	var msg *string
	if !db.IsShortUrlAvailable(shortUrl, "mh") {
		msg, _ = utils.ConvertResponseMessageToJson(constants.UNAVAILABLE_LINK, "")
		return c.String(http.StatusNotFound, *msg)
	}
	temp, _ := db.GetUrlByShortForm(shortUrl)
	if dateExtractor.IsLinkExpired(temp.Created_at) {
		msg, _ = utils.ConvertResponseMessageToJson(constants.EXPIRED_LINK, "")
		return c.String(http.StatusForbidden, *msg)
	}
	return c.Redirect(http.StatusSeeOther, temp.Original_url)
}

func GetUserUrls(c echo.Context) error {
	token, err := c.Cookie("token")
	var msg *string
	if err != nil {
		msg, _ = utils.ConvertResponseMessageToJson(constants.UNAUTHORIZED_USER, "")
		return c.String(http.StatusUnauthorized, *msg)
	}
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("mh_secret"), nil
	})
	if err != nil {
		msg, _ = utils.ConvertResponseMessageToJson(constants.UNAUTHORIZED_USER, "")
		return c.String(http.StatusUnauthorized, *msg)
	}
	rows, err := db.GetUrlsByUsername(claims["usr"].(string))
	return c.JSON(http.StatusOK, rows)
}
