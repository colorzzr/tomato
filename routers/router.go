package routers

import (
	"github.com/astaxie/beego"
	"github.com/lfq7413/tomato/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/classes",
			beego.NSInclude(
				&controllers.ClassesController{},
			),
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/logout",
			beego.NSInclude(
				&controllers.LogoutController{},
			),
		),
		beego.NSNamespace("/requestPasswordReset",
			beego.NSInclude(
				&controllers.ResetController{},
			),
		),
		beego.NSNamespace("/sessions",
			beego.NSInclude(
				&controllers.SessionsController{},
			),
		),
		beego.NSNamespace("/roles",
			beego.NSInclude(
				&controllers.RolesController{},
			),
		),
		beego.NSNamespace("/files",
			beego.NSInclude(
				&controllers.FilesController{},
			),
		),
		beego.NSNamespace("/events",
			beego.NSInclude(
				&controllers.AnalyticsController{},
			),
		),
		beego.NSNamespace("/push",
			beego.NSInclude(
				&controllers.PushController{},
			),
		),
		beego.NSNamespace("/installations",
			beego.NSInclude(
				&controllers.InstallationsController{},
			),
		),
		beego.NSNamespace("/functions",
			beego.NSInclude(
				&controllers.FunctionsController{},
			),
		),
		beego.NSNamespace("/jobs",
			beego.NSInclude(
				&controllers.JobsController{},
			),
		),
		beego.NSNamespace("/schemas",
			beego.NSInclude(
				&controllers.SchemasController{},
			),
		),
		beego.NSNamespace("/apps",
			beego.NSInclude(
				&controllers.PublicController{},
			),
		),
		beego.NSNamespace("/purge",
			beego.NSInclude(
				&controllers.PurgeController{},
			),
		),
		beego.NSNamespace("/config",
			beego.NSInclude(
				&controllers.GlobalConfigController{},
			),
		),
		beego.NSNamespace("/scriptlog",
			beego.NSInclude(
				&controllers.LogsController{},
			),
		),
		beego.NSNamespace("/validate_purchase",
			beego.NSInclude(
				&controllers.IAPValidationController{},
			),
		),
		beego.NSNamespace("/serverInfo",
			beego.NSInclude(
				&controllers.FeaturesController{},
			),
		),
		beego.NSNamespace("/hooks",
			beego.NSInclude(
				&controllers.HooksController{},
			),
		),
		beego.NSNamespace("/cloud_code",
			beego.NSInclude(
				&controllers.CloudCodeController{},
			),
		),
		beego.NSNamespace("/batch",
			beego.NSInclude(
				&controllers.BatchController{},
			),
		),
		beego.NSNamespace("/upgradeToRevocableSession",
			beego.NSInclude(
				&controllers.UpgradeSessionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
