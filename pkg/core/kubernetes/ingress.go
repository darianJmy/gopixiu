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

package kubernetes

import (
	"context"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/caoyingjunz/gopixiu/api/types"
	"github.com/caoyingjunz/gopixiu/pkg/log"
)

type IngressGetter interface {
	Ingresses(cloud string) IngressInterface
}

type IngressInterface interface {
	Create(ctx context.Context, ingress *networkingv1.Ingress) error
	Update(ctx context.Context, ingress *networkingv1.Ingress) error
	Delete(ctx context.Context, deleteOptions types.GetOrDeleteOptions) error
	Get(ctx context.Context, getOptions types.GetOrDeleteOptions) (*networkingv1.Ingress, error)
	List(ctx context.Context, listOptions types.ListOptions) ([]networkingv1.Ingress, error)
}

type Ingress struct {
	client *kubernetes.Clientset
	cloud  string
}

func NewIngress(c *kubernetes.Clientset, cloud string) *Ingress {
	return &Ingress{
		client: c,
		cloud:  cloud,
	}
}

func (c *Ingress) Create(ctx context.Context, ingress *networkingv1.Ingress) error {
	if c.client == nil {
		return clientError
	}
	if _, err := c.client.NetworkingV1().
		Ingresses(ingress.Namespace).
		Create(ctx, ingress, metav1.CreateOptions{}); err != nil {
		log.Logger.Errorf("failed to create %s statefulSet: %v", c.cloud, err)
		return err
	}

	return nil
}

func (c *Ingress) Update(ctx context.Context, ingress *networkingv1.Ingress) error {
	if c.client == nil {
		return clientError
	}
	if _, err := c.client.NetworkingV1().
		Ingresses(ingress.Namespace).
		Update(ctx, ingress, metav1.UpdateOptions{}); err != nil {
		log.Logger.Errorf("failed to update %s statefulSet: %v", c.cloud, err)
		return err
	}

	return nil
}

func (c *Ingress) Delete(ctx context.Context, deleteOptions types.GetOrDeleteOptions) error {
	if c.client == nil {
		return clientError
	}
	if err := c.client.NetworkingV1().
		Ingresses(deleteOptions.Namespace).
		Delete(ctx, deleteOptions.ObjectName, metav1.DeleteOptions{}); err != nil {
		log.Logger.Errorf("failed to delete %s statefulSet: %v", deleteOptions.CloudName, err)
		return err
	}

	return nil
}

func (c *Ingress) Get(ctx context.Context, getOptions types.GetOrDeleteOptions) (*networkingv1.Ingress, error) {
	if c.client == nil {
		return nil, clientError
	}
	ingress, err := c.client.NetworkingV1().
		Ingresses(getOptions.Namespace).
		Get(ctx, getOptions.ObjectName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorf("failed to get %s statefulSets: %v", getOptions.CloudName, err)
		return nil, err
	}

	return ingress, err
}

func (c *Ingress) List(ctx context.Context, listOptions types.ListOptions) ([]networkingv1.Ingress, error) {
	if c.client == nil {
		return nil, clientError
	}
	ingress, err := c.client.NetworkingV1().
		Ingresses(listOptions.Namespace).
		List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.Errorf("failed to list statefulsets: %v", listOptions.Namespace, err)
		return nil, err
	}

	return ingress.Items, err
}
