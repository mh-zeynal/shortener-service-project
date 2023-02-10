package handlers

import (
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
	if db.IsOriginalUrlAvailable(body.IntendedUrl, "mh") {
		url, _ := db.GetUrlByOriginal(body.IntendedUrl, "mh")
		return c.String(http.StatusOK, "this url already exists:\n"+
			utils.GetVariable("DOMAIN_NAME")+url.Short_url)
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
	return c.String(http.StatusOK, "shortened URL: localhost:9090/"+short)
}

func HandleRedirects(c echo.Context) error {
	shortUrl := c.Param("s")
	if !db.IsShortUrlAvailable(shortUrl, "mh") {
		return c.String(http.StatusNotFound, "this link is unavailable")
	}
	temp, _ := db.GetUrlByShortForm(shortUrl)
	if dateExtractor.IsLinkExpired(temp.Created_at) {
		return c.String(http.StatusForbidden, "this link is expired")
	}
	return c.Redirect(http.StatusSeeOther, temp.Original_url)
}

func GetUserUrls(c echo.Context) error {
	token, err := c.Cookie("token")
	if err != nil {
		return c.String(http.StatusUnauthorized, "user not logged in")
	}
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("mh_secret"), nil
	})
	if err != nil {
		return c.String(http.StatusUnauthorized, "user not logged in")
	}
	rows, err := db.GetUrlsByUsername(claims["usr"].(string))
	return c.JSON(http.StatusOK, rows)
}
