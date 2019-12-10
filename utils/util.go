/**
 * @Author: cristi
 * @Description:
 * @File:  util.go
 * @Version: 1.0.0
 * @Date: 2019/12/10 0010 14:08
 */

package utils

import "regexp"

//mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
