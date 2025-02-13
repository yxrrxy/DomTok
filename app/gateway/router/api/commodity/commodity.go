/*
Copyright 2024 The west2-online Authors.

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

// Code generated by hertz generator. DO NOT EDIT.

package commodity

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	commodity "github.com/west2-online/DomTok/app/gateway/handler/api/commodity"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			{
				_commodity := _v1.Group("/commodity", _commodityMw()...)
				{
					_category := _commodity.Group("/category", _categoryMw()...)
					_category.POST("/create", append(_createcategoryMw(), commodity.CreateCategory)...)
					_category.DELETE("/delete", append(_deletecategoryMw(), commodity.DeleteCategory)...)
					_category.GET("/search", append(_viewcategoryMw(), commodity.ViewCategory)...)
					_category.POST("/update", append(_updatecategoryMw(), commodity.UpdateCategory)...)
				}
				{
					_coupon := _commodity.Group("/coupon", _couponMw()...)
					_coupon.GET("/all", append(_viewuserallcouponMw(), commodity.ViewUserAllCoupon)...)
					_coupon.POST("/create", append(_createcouponMw(), commodity.CreateCoupon)...)
					_coupon.DELETE("/delete", append(_deletecouponMw(), commodity.DeleteCoupon)...)
					_coupon.POST("/receive", append(_createusercouponMw(), commodity.CreateUserCoupon)...)
					_coupon.GET("/search", append(_viewcouponMw(), commodity.ViewCoupon)...)
					_coupon.POST("/use", append(_useusercouponMw(), commodity.UseUserCoupon)...)
				}
				{
					_price := _commodity.Group("/price", _priceMw()...)
					_price.GET("/history", append(_viewhistoryMw(), commodity.ViewHistory)...)
				}
				{
					_sku := _commodity.Group("/sku", _skuMw()...)
					_sku.POST("/attr", append(_uploadskuattrMw(), commodity.UploadSkuAttr)...)
					_sku.POST("/create", append(_createskuMw(), commodity.CreateSku)...)
					_sku.DELETE("/delete", append(_deleteskuMw(), commodity.DeleteSku)...)
					_sku.GET("/image", append(_viewskuimageMw(), commodity.ViewSkuImage)...)
					_sku.GET("/list", append(_listskuinfoMw(), commodity.ListSkuInfo)...)
					_sku.GET("/search", append(_viewskuMw(), commodity.ViewSku)...)
					_sku.POST("/upadte", append(_updateskuMw(), commodity.UpdateSku)...)
				}
				{
					_spu := _commodity.Group("/spu", _spuMw()...)
					_spu.POST("/create", append(_createspuMw(), commodity.CreateSpu)...)
					_spu.DELETE("/delete", append(_deletespuMw(), commodity.DeleteSpu)...)
					_spu.GET("/image", append(_viewspuimageMw(), commodity.ViewSpuImage)...)
					_spu.GET("/search", append(_viewspuMw(), commodity.ViewSpu)...)
					_spu.POST("/update", append(_updatespuMw(), commodity.UpdateSpu)...)
				}
			}
		}
	}
}
