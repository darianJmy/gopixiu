/*
Copyright 2021 The Pixiu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cicd

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/caoyingjunz/gopixiu/api/server/httputils"
	"github.com/caoyingjunz/gopixiu/pkg/pixiu"
)

func (s *cicdRouter) runJob(c *gin.Context) {
	r := httputils.NewResponse()

	jobName := c.Param("jobName")
	if jobName == "" {
		httputils.SetFailed(c, r, fmt.Errorf("jobName should not be empty"))
		return
	}
	err := pixiu.CoreV1.Cicd().RunJob(context.TODO(), jobName)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (s *cicdRouter) createJob(c *gin.Context) {
	r := httputils.NewResponse()
	err := pixiu.CoreV1.Cicd().CreateJob(context.TODO())
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (s *cicdRouter) deleteJob(c *gin.Context) {
	r := httputils.NewResponse()

	jobName := c.Param("jobName")
	if jobName == "" {
		httputils.SetFailed(c, r, fmt.Errorf("jobName should not be empty"))
		return
	}
	err := pixiu.CoreV1.Cicd().DeleteJob(context.TODO(), jobName)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (s *cicdRouter) addViewJob(c *gin.Context) {
	r := httputils.NewResponse()

	viewName := c.Param("viewName")
	if viewName == "" {
		httputils.SetFailed(c, r, fmt.Errorf("viewName should not be empty"))
		return
	}
	err := pixiu.CoreV1.Cicd().AddViewJob(context.TODO(), viewName)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}
