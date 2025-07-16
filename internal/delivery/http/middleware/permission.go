package middleware

import (
	"net/http"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/pkg/errorsx"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(permissionName string, db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		if userID == 0 {
			httpx.ResponseError(c, "unauthorized", http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		var count int64

		err := db.
			Table("user_roles").
			Joins("JOIN role_permissions ON user_roles.role_id = role_permissions.role_id").
			Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
			Where("user_roles.user_id = ? AND permissions.name = ?", userID, permissionName).
			Count(&count).Error
		if err != nil {
			err = errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
			httpx.HandleServiceError(c, err)
			c.Abort()
			return
		}

		if count == 0 {
			httpx.ResponseError(c, "forbidden", http.StatusForbidden, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
