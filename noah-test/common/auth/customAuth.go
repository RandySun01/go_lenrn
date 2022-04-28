package auth

/*
Author: RandySun
Date: 2022/3/7 5:19 下午
*/

//// CustomAuthMd5Token 生成认证token
//func CustomAuthMd5Token() string {
//
//	timeStamp := utils.GetTimeStamp()
//	tmp := fmt.Sprintf("%s|%d", common.CUSTOMTOKENSECRETKEY, timeStamp)
//	md5Str := GetMd5Token(tmp)
//
//	//拼接认证token
//	authToken := fmt.Sprintf("%s|%d", md5Str, timeStamp)
//	zap.L().Info("CustomAuthMd5Token auth token",
//		zap.String("tmp", tmp),
//		zap.String("authToken", authToken),
//	)
//	fmt.Println(tmp, timeStamp, authToken)
//
//	return authToken
//}
//
//// CustomServiceParseToken 验证agent
//func CustomServiceParseToken(clientTokenTime string) (bool, error) {
//	// 拆分token和时间戳
//	tokenTime := strings.Split(clientTokenTime, "|")
//
//	clientToken := tokenTime[0]
//	clientTime, err := strconv.ParseInt(tokenTime[1], 10, 64)
//	if err != nil {
//		zap.L().Error("CustomServiceParseToken parse time stamp failed", zap.Error(err))
//		return false, err
//	}
//
//	fmt.Println(clientToken, clientTime)
//
//	zap.L().Info(
//		"CustomServiceParseToken token and stamp",
//		zap.String("clientToken", clientToken),
//		zap.Int64("clientTime", clientTime),
//	)
//
//	// 计算服务端和agent时间差
//	serviceNowTime := time.Now()
//	clientTimeObj := time.Unix(clientTime, 0)
//
//	if serviceNowTime.Sub(clientTimeObj).Minutes() > common.TOKENEXPIRETIME {
//		zap.L().Info("CustomServiceParseToken token expire")
//		return false, err
//	}
//
//	// 校验服务端和客户端token是否一致
//	fmt.Println(tokenTime[1])
//	tmp := fmt.Sprintf("%s|%d", common.CUSTOMTOKENSECRETKEY, clientTime)
//	serviceToken := GetMd5Token(tmp)
//	if serviceToken == clientToken {
//		return true, nil
//	}
//	// token不一致
//	return false, nil
//
//}
//
//// GetMd5Token 对字符串md5
//func GetMd5Token(tmp string) string {
//	h := md5.New()
//	h.Write([]byte(tmp))
//	md5Str := hex.EncodeToString(h.Sum(nil))
//	zap.L().Info(
//		"GetMd5Token md5  token",
//		zap.String("md5Str", md5Str),
//		zap.String("tmp", tmp),
//	)
//	return md5Str
//}
