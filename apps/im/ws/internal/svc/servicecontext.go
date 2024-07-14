/**
 * @author: Yanko/xiaoxiaoyang-sheep
 * @doc:
 **/

package svc

import "easy-chat/apps/im/ws/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}