package main

import (
	_ "yinji/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
	_ "github.com/go-sql-driver/mysql"
	"yinji/service"
	"github.com/astaxie/beego/session"
	"yinji/thread"
)


func Init(){
	RegistOrm();
	SessionInit()

	thread.ParseUserTempToUserDate()
	thread.ParseAudioTempToUserDate()
}

func RegistOrm(){

	// 进行注册  相对应的 信息
	orm.RegisterDriver("mysql",orm.DRMySQL);

	/**
		var maxId =
		var maxConnection = 30;
	 */
	orm.RegisterDataBase("default","mysql","root:980621@/yinji?charset=utf8&loc=Local")
	//orm.RegisterDataBase("default","mysql","root:@Linxiang621!@tcp(cdb-nj4f7n3e.cd.tencentcdb.com:10025)/yinji?charset=utf8&loc=Local")
	//设置 相对应的 信息 内容
	orm.RegisterModel(
		new(bean.Login),
		new(bean.User),
		new(bean.Audio),
		new(bean.AudioComment),
		new(bean.AudioCommentReply),

		//对应的信息
		new(bean.Resource),
		new(bean.HostResource),
		new(bean.ResourceImage),
		new(bean.ResourceAudio),

		//收藏有关的bean类
		new(bean.AudioUserCollection),
		new(bean.AudioDashboard),
		new(bean.AudioCollectionFolder),

		//点赞 Bean
		new(bean.AudioUserLove),

		//用户角色这里的三个
		new(bean.UserDashboard),
		new(bean.UserDateDashboard),
		new(bean.UserTempDashboard),
		new(bean.UserDetails),

		//详细记录对应的信息
		new(bean.AudioUserHistroy),
		new(bean.AudioFast),
		new(bean.AudioTempDashboard),
		);

	orm.Debug = true;

}

func TemplateInit(){
}

//初始化对应的 session
func SessionInit(){
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}

	//这里设置不同对应的方式
	var globalSession , _ = session.NewManager("memory" , sessionConfig)
	go globalSession.GC()

}

func main() {
	Init();
	//设置相对应的 静态文件的 代码
	var app = service.AppService{"yinji"}
	beego.SetStaticPath(app.GetUrl("/music"),"static/music")

	beego.SetStaticPath(app.GetUrl("/resources") , "static/resources")
	beego.SetStaticPath(app.GetUrl("/framework"),"static/framework")
	beego.SetStaticPath(app.GetUrl("/audio"),"static/audio")
	beego.SetStaticPath(app.GetUrl("/image"),"static/image")
	beego.SetStaticPath(app.GetUrl("/js"),"static/js")

	beego.SetStaticPath(app.GetUrl("/page/"),"static/page")
	beego.SetStaticPath(app.GetUrl("/static"),"static/static")
	beego.Run();
}

