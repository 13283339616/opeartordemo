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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AppSpec defines the desired state of App
type AppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of App. Edit app_types.go to remove/update
	//Foo string `json:"foo,omitempty"`

	//Action to do some object 对对象做什么动作 例如给一个Hello
	//+optional
	Action string `json:"action,omitempty"`

	//Object 对什么操作 optional可选 例如给一个world
	//+optional
	Object string `json:"object,omitempty"`
}

// AppStatus defines the observed state of App
type AppStatus struct {

	//Result 显示action+object
	//+optional
	Result string `json:"result,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// App 定义我们CRD的对象
// App is the Schema for the apps API
type App struct {
	metav1.TypeMeta   `json:",inline"`            //元信息
	metav1.ObjectMeta `json:"metadata,omitempty"` //元信息

	Spec   AppSpec   `json:"spec,omitempty"`   //CRD的核心内容 crd的描述定义
	Status AppStatus `json:"status,omitempty"` //CRD的现阶段的状态的内容
}

//+kubebuilder:object:root=true

// AppList 定义CRD的列表
// AppList contains a list of App
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []App `json:"items"`
}

func init() {
	//注册gostruct 到scheme
	SchemeBuilder.Register(&App{}, &AppList{})
}
