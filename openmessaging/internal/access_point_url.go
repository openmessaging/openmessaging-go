package internal

import (
	"fmt"
	. "git.ucloudadmin.com/paas_vppd/prj-wintersoldier/openmessaging-go/common"
	"regexp"
	"strings"
)

var (
	Prefix  = "oms"
	PATTERN = fmt.Sprintf("^%s:.+://.+/.+$", Prefix)
)

/**
 * The standard OMS access point schema is:
 * <p>
 * {@literal oms:<driver_type>://[account_id@]host1[:port1][,host2[:port2],...[,hostN[:portN]]]/<region>}
 * <p>
 *
	oms:rocketmq://admin@127.0.0.1:9876,127.0.0.1:9875/bj01
 * More details please refer to:
 * <a href="https://github.com/openmessaging/specification/blob/master/oms_access_point_schema.md">Access Point Schema</a>
*/

type AccessPoint struct {
	Prefix    string
	Driver    string
	AccessKey string
	SecretKey string
	Address   string
	Region    string
}

// AccessPointURL oms:rocketmq://admin#dsds@127.0.0.1:9876,127.0.0.1:9875/bj01
func AccessPointURI(accessPointURL string) (*AccessPoint, error) {
	if err := validateAccessPoint(accessPointURL); err != nil {
		return nil, err
	}
	authHead := strings.Split(accessPointURL, ":")
	authData := strings.Split(accessPointURL, "//")[1]
	accessKey := strings.Split(authData, "#")[0]
	rightData := strings.SplitAfter(authData, "#")[1]
	url := strings.Split(rightData, "@")[1]
	return &AccessPoint{
		Prefix:    authHead[0],
		Driver:    authHead[1],
		AccessKey: accessKey,
		SecretKey: strings.Split(rightData, "@")[0],
		Address:   strings.Split(url, "/")[0],
		Region:    strings.Split(url, "/")[1],
	}, nil
}

func validateAccessPoint(accessPoint string) error {
	ok, err := regexp.MatchString(PATTERN, accessPoint)
	if err != nil || !ok {
		code, msg := OMSRuntimeStatusMsg(STATUS_1101)
		return fmt.Errorf("%d %s %s", code, msg, accessPoint)
	}
	return nil
}
