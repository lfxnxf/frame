package lib

import (
	"context"
	"fmt"

	"github.com/lfxnxf/frame/BackendPlatform/go-tools/ext"
	"github.com/lfxnxf/frame/BackendPlatform/go-tools/http"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"golang.org/x/sync/errgroup"
)

func GetUserInfo(c context.Context, uids []int64) ([]RUserInfo, error) {
	serviceName := "user.profile.gateway"
	operationURL := "/user/infos?aggs=rank&id=%v"
	limit := 200
	taskLen := len(uids) / limit
	if len(uids)%limit > 0 {
		taskLen++
	}
	tasks := make([][]int64, taskLen)
	result := make(chan RUserInfo, 10000)
	var i int
	for i = 0; i < taskLen-1; i++ {
		tasks[i] = uids[i*limit : (i+1)*limit]
	}
	if len(uids)%limit > 0 {
		tasks[taskLen-1] = uids[i*limit:]
	}
	var g errgroup.Group
	fetch := func(idx int) error {
		var r UserInfoResp
		url := fmt.Sprintf(operationURL, ext.Int64Join(tasks[idx], ","))
		err := http.GetMustSucc(c, serviceName, url, &r)
		if err != nil {
			return err
		}
		for _, v := range r.Infos {
			result <- v
		}
		return nil
	}
	for k := 0; k < taskLen; k++ {
		v := k
		g.Go(func() error {
			return fetch(v)
		})
	}
	if err := g.Wait(); err != nil {
		logging.Warnf("get user info err,uids %v,err %v", uids, err)
	}
	close(result)
	m := make(map[int64]RUserInfo)
	for v := range result {
		m[v.User.ID] = v
	}
	ret := make([]RUserInfo, len(uids))
	for i, v := range uids {
		vv, ok := m[v]
		if !ok {
			// 默认资料
			vv = RUserInfo{
				User: UserInfo{
					ID:       v,
					Nick:     fmt.Sprintf("inke%d", v),
					Portrait: "http://img.ikstatic.cn/MTUyODQyMzA0NTk2NiM2MjcjanBn.jpg"},
			}
		}
		ret[i] = vv
	}
	return ret, nil
}
