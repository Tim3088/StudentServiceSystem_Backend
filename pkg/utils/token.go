package utils

import (
	"StudentServiceSystem/internal/global"
	"StudentServiceSystem/internal/model"
	"StudentServiceSystem/internal/pkg/redis"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"time"
)

var Secret = []byte(global.Config.GetString("jwt.pass")) // 获取签发jwt密钥

const AccessTokenExpireDuration = time.Hour * 24      // 访问令牌有效期1天
const RefreshTokenExpireDuration = time.Hour * 24 * 7 // 刷新令牌有效期7天
// GenerateTokens 生成访问令牌和刷新令牌
func GenerateTokens(userID int) (string, string, error) {
	// 生成访问令牌
	accessTokenString, err := generateAccessToken(userID)
	if err != nil {
		zap.L().Error("生成访问令牌失败", zap.Error(err))
		return "", "", err
	}
	// 生成刷新令牌
	refreshClaims := model.MyClaims{
		UserID:    userID,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpireDuration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                 // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                 // 生效时间
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(Secret)
	if err != nil {
		zap.L().Error("生成刷新令牌失败", zap.Error(err))
		return "", "", err
	}
	// 将令牌存储到Redis
	ctx := context.Background()
	err = redis.Rdb.Set(ctx, "access_token:"+accessTokenString, userID, AccessTokenExpireDuration).Err()
	if err != nil {
		zap.L().Error("访问令牌存储失败", zap.Error(err))
		return "", "", err
	}
	err = redis.Rdb.Set(ctx, "refresh_token:"+refreshTokenString, userID, RefreshTokenExpireDuration).Err()
	if err != nil {
		zap.L().Error("刷新令牌存储失败", zap.Error(err))
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// ParseToken 解析令牌
func ParseToken(tokenString string) (*model.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		zap.L().Error("解析Token失败", zap.String("token", tokenString), zap.Error(err))
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		// 检查Redis中是否存在该令牌
		ctx := context.Background()
		key := "access_token:" + tokenString
		if claims.TokenType == "refresh" {
			key = "refresh_token:" + tokenString
		}
		_, err := redis.Rdb.Get(ctx, key).Result()
		if err != nil {
			zap.L().Error("无效的令牌", zap.String("token", tokenString), zap.Error(err))
			return nil, err
		}
		return claims, nil
	}
	return nil, errors.New("无效的令牌")
}

// generateAccessToken 生成访问令牌
func generateAccessToken(userID int) (string, error) {
	accessClaims := model.MyClaims{
		UserID:    userID,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpireDuration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                // 生效时间
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(Secret)
	if err != nil {
		zap.L().Error("生成访问令牌失败", zap.Error(err))
		return "", err
	}
	return accessTokenString, nil
}

// RefreshTokenHandler 处理刷新令牌请求
func RefreshTokenHandler(c *gin.Context) {
	var data struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("处理刷新令牌请求失败", zap.Error(err))
		JsonFail(c, 200501, "参数错误")
	}

	// 解析刷新令牌
	token, err := jwt.ParseWithClaims(data.RefreshToken, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil || !token.Valid {
		JsonFail(c, 200509, "无效的刷新令牌")
	}

	// 验证刷新令牌的有效性
	claims, ok := token.Claims.(*model.MyClaims)
	if !ok || claims.ExpiresAt.Time.Before(time.Now()) || claims.TokenType != "refresh" {
		JsonFail(c, 200509, "无效的刷新令牌")
	}

	// 生成新的访问令牌
	newToken, err := generateAccessToken(claims.UserID)
	if err != nil {
		zap.L().Error("生成新的刷新令牌失败", zap.Error(err))
	}

	// 将新的访问令牌存储到Redis
	err = redis.Rdb.Set(c, "access_token:"+newToken, claims.UserID, AccessTokenExpireDuration).Err()
	if err != nil {
		zap.L().Error("新的刷新令牌存储失败", zap.Error(err))
	}

	JsonSuccess(c, gin.H{
		"access_token": newToken,
	})
}
