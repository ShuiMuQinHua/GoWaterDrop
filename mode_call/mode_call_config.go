package mode_call

const (
	TOPIC  = "o_iaas_modulecall"
	APP_ID = "10032"
	UDP_IP = "127.0.0.1"

	API_ID_STR  = "api_id"
	MODULE_ID   = "module_id"
	MODULE      = "module"
	CALL_IP     = "call_ip"
	RECEIVE_IP  = "receive_ip"
	RESULT      = "result"
	RESULT_CODE = "result_code"
	COST        = "response_time"

	DEFAULT_MODULE_ID = "999"
	UDP_PORT          = 40001
	DEPLOYMENT_REDIS  = "deployment_redis"
)

var MODULES = map[string]interface{}{
	"test":           0,
	DEPLOYMENT_REDIS: 10000270,
}
