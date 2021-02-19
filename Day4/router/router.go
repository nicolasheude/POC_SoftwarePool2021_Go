package router

import (
	"SofwareGoDay4/User"
	"SofwareGoDay4/softcrypto"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func ifUserExists(email string) int {
	lenDb := len(User.UserDB) - 1
	if lenDb < 0 {
		return 0
	}
	for i := 0; i <= lenDb; i++ {
		if User.UserDB[i].Email == email {
			return 84
		}
	}
	return 0
}

func signupSession(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if ifUserExists(email) == 84 {
			c.String(http.StatusBadRequest, "The user already exists, please try again with another email address.")
			return
		}
		User.UserDB = append(User.UserDB, User.User{Email: email, Password: password})
		c.String(http.StatusCreated, "User successfully created !")
	}
	return
}

func ifIdCorrect(email, password string) int {
	lenDb := len(User.UserDB) - 1
	if lenDb < 0 {
		return 84
	}
	for i := 0; i <= lenDb; i++ {
		if User.UserDB[i].Email == email && User.UserDB[i].Password == password {
			return 0
		}
	}
	return 84
}

func singinSession(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if ifIdCorrect(email, password) == 84 {
			c.String(http.StatusBadRequest, "No user corresponds to the transmitted identifier, please try again.")
			return
		}
		emailCry := softcrypto.Encrypt(email, hex.EncodeToString([]byte(os.Getenv("KEY"))))
		c.SetCookie("email", emailCry, 3600, "/singin-session", "0.0.0.0", true, true)
		c.String(http.StatusOK, "User successfully logged in !")
	}
	return
}

func singinSessionWhKey(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if ifIdCorrect(email, password) == 84 {
			c.String(http.StatusBadRequest, "No user corresponds to the transmitted identifier, please try again.")
			return
		}
		emailCry := softcrypto.Encrypt(email, hex.EncodeToString([]byte(os.Getenv("KEY"))))
		c.SetCookie("email", emailCry, 3600, "/singin-session", "0.0.0.0", true, true)
		c.String(http.StatusOK, "User successfully logged in !\n KEY=%s\n", emailCry)
	}
	return
}

func lookPasswordCosEmail(email string) (string, int) {
	lenDb := len(User.UserDB) - 1
	if lenDb < 0 {
		return  "", 84
	}
	for i := 0; i <= lenDb; i++ {
		if User.UserDB[i].Email == email {
			return User.UserDB[i].Password, 0
		}
	}
	return "", 84
}

func meSession(c *gin.Context) {
	keyCookie := c.GetHeader("email")
	if keyCookie == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	emailUnCry := softcrypto.Decrypt(keyCookie, hex.EncodeToString([]byte(os.Getenv("KEY"))))
	pass, err := lookPasswordCosEmail(emailUnCry)
	fmt.Println(emailUnCry, pass)
	if err == 84 {
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {
		c.String(http.StatusOK, "Email : %s\nPassword : %s\n", emailUnCry, pass)
	}
	fmt.Println(emailUnCry)
	return
}

func ifUserExistsJwt(email string) error {
	lenDb := len(User.UserDB) - 1
	if lenDb < 0 {
		return errors.New("No database")
	}
	for i := 0; i <= lenDb; i++ {
		if User.UserDB[i].Email == email {
			return nil
		}
	}
	return errors.New("No user find")
}

func signupJwt(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if err := ifUserExistsJwt(email); err != nil {
			c.String(http.StatusBadRequest, "The user already exists, please try again with another email address. ERROR : %w", err)
			return
		}
		User.UsersJWT = append(User.UsersJWT, User.UserJWT{Email: email, Password: password})
		c.String(http.StatusCreated, "User successfully created !")
	}
	return
}

func tokenParser(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return User.APISECRET, nil
}

func tokenBd(token *jwt.Token) bool {
	temp, err := tokenParser(token)
	return false
}

func singinJwt(c *gin.Context) {
	return
}

func meJwt(c *gin.Context) {
	return
}

func ApplyRoutes(r *gin.Engine) {
	r.POST("/signup-session", signupSession)
	r.POST("/singin-session", singinSession)
	r.POST("/singin-session-wh-key", singinSessionWhKey)
	r.GET("/me-session", meSession)
	r.GET("/signup-jwt", signupJwt)
	r.GET("/singin-jwt", singinJwt)
	r.GET("/me-jwt", meJwt)
}
