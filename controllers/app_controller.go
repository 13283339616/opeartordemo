/*
Copyright 2024 dgj.

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

package controllers

import (
	"context"
	"fmt"
	//appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demov1 "github.com/13283339616/opeartordemo/api/v1"
)

// AppReconciler 是提供调和函数（把现阶段状态和我们定义的状态进行统一 趋近的调和函数） 对象中的数据 可以在上层统一管理
// AppReconciler reconciles a App object
type AppReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.mashibing.com,resources=apps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.mashibing.com,resources=apps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.mashibing.com,resources=apps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the App object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.1/pkg/reconcile
// 写业务逻辑的地方 核心功能函数要去定义operator的行为
func (r *AppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//一般程序中都会记录这个logger
	logger := log.FromContext(ctx)
	//防止进入死循环
	logger.Info("start reconcile")

	//实现你的业务逻辑在这里
	//1.获取资源对象
	app := &demov1.App{}
	err := r.Client.Get(ctx, req.NamespacedName, app)
	if err != nil {
		logger.Error(err, "Failed to get app")
		//不存在的资源 不报错
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	//2.处理数据
	//2.1获取对应的字段
	action := app.Spec.Action
	object := app.Spec.Object

	//2.2 拼接数据
	result := fmt.Sprintf("%s,%s", action, object)

	//3. 创建结果
	//3.1 定义crd的status
	//3.2 完成reconcile

	//本地申请空间 app指向缓存区 apiserver 同步 有其他地方可能也引用它 为了避免这些问题 更新之前copy出来
	appCopy := app.DeepCopy()
	appCopy.Status.Result = result

	//更新状态 这个更新指的是更新appStatus
	err = r.Client.Status().Update(ctx, appCopy)

	//设置资源的关系
	//controllerutil.SetControllerReference(app, appCopy, r.Scheme)

	if err != nil {
		return ctrl.Result{}, err
	}

	logger.Info("end reconcile")
	//根据返回值 退出 再次进来
	return ctrl.Result{}, nil
}

// SetupWithManager 把Reconciler加入到controller 将controller 注册到manager中
// 将我们的crd加入”过滤器“ 可以到controller订阅到crd的变化
// SetupWithManager sets up the controller with the Manager.
func (r *AppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.App{}).
		//设置子资源的关系
		//Owns(&appsv1.Deployment{}).
		Complete(r)
}
