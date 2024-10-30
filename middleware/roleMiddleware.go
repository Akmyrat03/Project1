package middleware

// import (
// 	"net/http"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// type RoleBasedClaims struct {
// 	jwt.StandardClaims        //It IssuedAt , ExpiresAt informations hold
// 	UserId             int    `json:"user_id"`
// 	Role               string `json:"role"`
// }

// // RoleBasedAuth istenilen rolun kullanici olup olmadigni kontrol eder
// func RoleBasedAuth(requiredRole string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Adim 1: Authorization basligindan token'i al
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error": "Authorization token required",
// 			})
// 			return
// 		}

// 		// Adim 2: Token'i coz ve dogrula
// 		token, err := jwt.ParseWithClaims(tokenString, &RoleBasedClaims{}, func(token *jwt.Token) (interface{}, error) {
// 			return []byte("###%5645646566"), nil
// 		})

// 		if err != nil {
// 			// Handle parsing errors (e.g., expired or invalid token)
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
// 			c.Abort()
// 			return
// 		}

// 		// Adim 3: Token gecerli mi ve rol uyuyor mu?
// 		if claims, ok := token.Claims.(*RoleBasedClaims); ok && token.Valid {
// 			if claims.Role != requiredRole {
// 				// Eger kullanici rolu gereken role uymuyorsa, erisimi reddet
// 				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 					"error": "Erisim reddedildi",
// 				})
// 				return
// 			}
// 			// Eger rol uyuyorsa kullanici bilgilerni baglama ekle
// 			c.Set("user_id", claims.UserId)
// 			c.Set("role", claims.Role)
// 		} else {
// 			// Token gecerli degilse, yetkisiz hatasi dondur
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error": "gecersiz token",
// 			})
// 			return
// 		}

// 		// Tum kontroller gecilirse istegi devam etdir
// 		c.Next()
// 	}
// }
