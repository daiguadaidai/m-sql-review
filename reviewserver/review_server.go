package reviewserver

import (
	"sync"
	"net/http"
	"github.com/daiguadaidai/m-sql-review/reviewserver/handle"
	log "github.com/cihub/seelog"
	"github.com/daiguadaidai/m-sql-review/config"
	"fmt"
)

/* 启动Http服务
Params:
	_httpServerConfig: 启动http服务的配置
	_wg: 并发等待值
 */
func StartReviewServer(_httpServerConfig *config.HttpServerConfig, _wg *sync.WaitGroup) {
	defer _wg.Done()

	// 添加路由
	http.HandleFunc("/hello", handle.SqlReviewHandle)

	log.Infof("监听地址为: %v", _httpServerConfig.Address())
	fmt.Printf("监听地址为: %v\n", _httpServerConfig.Address())
	err := http.ListenAndServe(_httpServerConfig.Address(), nil)
	if err != nil {
		log.Errorf("启动服务出错: %v", err)
	}
}
