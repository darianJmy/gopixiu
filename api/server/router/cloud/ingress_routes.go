package cloud

import (
	"context"

	"github.com/gin-gonic/gin"
	networkingv1 "k8s.io/api/networking/v1"

	"github.com/caoyingjunz/gopixiu/api/server/httputils"
	"github.com/caoyingjunz/gopixiu/api/types"
	"github.com/caoyingjunz/gopixiu/pkg/pixiu"
)

func (s *cloudRouter) createIngress(c *gin.Context) {
	r := httputils.NewResponse()
	var (
		err           error
		createOptions types.GetOrCreateOptions
		ingress       networkingv1.Ingress
	)
	if err = c.ShouldBindUri(&createOptions); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}
	ingress.Name = createOptions.ObjectName
	ingress.Namespace = createOptions.Namespace
	err = pixiu.CoreV1.Cloud().Ingresses(createOptions.CloudName).Create(context.TODO(), &ingress)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (s *cloudRouter) updateIngress(c *gin.Context) {
	r := httputils.NewResponse()
	var (
		err           error
		createOptions types.GetOrCreateOptions
		ingress       networkingv1.Ingress
	)
	if err = c.ShouldBindUri(&createOptions); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}
	ingress.Name = createOptions.ObjectName
	ingress.Namespace = createOptions.Namespace
	err = pixiu.CoreV1.Cloud().Ingresses(createOptions.CloudName).Update(context.TODO(), &ingress)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (s *cloudRouter) deleteIngress(c *gin.Context) {
	r := httputils.NewResponse()
	var (
		err        error
		delOptions types.GetOrDeleteOptions
	)
	if err = c.ShouldBindUri(&delOptions); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}
	err = pixiu.CoreV1.Cloud().Ingresses(delOptions.CloudName).Delete(context.TODO(), delOptions)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (s *cloudRouter) getIngress(c *gin.Context) {
	r := httputils.NewResponse()
	var (
		err        error
		getOptions types.GetOrDeleteOptions
	)
	if err = c.ShouldBindUri(&getOptions); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}
	r.Result, err = pixiu.CoreV1.Cloud().Ingresses(getOptions.CloudName).Get(context.TODO(), getOptions)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (s *cloudRouter) listIngress(c *gin.Context) {
	r := httputils.NewResponse()
	var (
		err         error
		listOptions types.ListOptions
	)
	if err = c.ShouldBindUri(&listOptions); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}
	r.Result, err = pixiu.CoreV1.Cloud().Ingresses(listOptions.CloudName).List(context.TODO(), listOptions)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}
