package routers

import (
	"yinji/controllers"
	"github.com/astaxie/beego"
	"yinji/service"
)

func init() {

	var app = service.AppService{"yinji"};
	var test_controller = &controllers.TestController{}

	var audio = &controllers.AudioController{};
	var user = &controllers.UserController{}

	var login = &controllers.LoginController{}
	var audioComment = &controllers.AudioCommentController{}

	var audioCommentReply = &controllers.AudioCommentReplyController{}

	var collection = &controllers.CollectionController{}

	//var dashborad = &controllers.DashboardContrlller{}

	var collectionFolder = &controllers.AudioCollectionFolderController{}

	var love = &controllers.AudioUserLoveController{}

	var userTempDashboard = &controllers.UserTempDashoardController{}

	var userDateDashboard = &controllers.UserDateDashboardController{}

	var userDashboard = &controllers.UserDashboardController{}

	var resource = &controllers.ResourceController{}

	var page = &controllers.PageController{}

	var userDetails = &controllers.UserDetailsController{}

	var audioUserHistory = &controllers.AudioUserHistoryController{}

	var audioFast = &controllers.AudioFastController{}
	var dashboard = &controllers.DashboardContrlller{}
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test/json", test_controller, "*:JsonResult")
	beego.Router("/test/xml", test_controller, "*:XmlResult");
	beego.Router("/test/page/:page", test_controller, "*:FindHtml")
	beego.Router(app.GetUrl("/music/player"), audio, "*:Player")
	beego.Router(app.GetUrl("/user/new"), user, "*:NewToDb")
	beego.Router(app.GetUrl("/user/new/test"), user, "*:UploadTest")
	/**
		这下面便是输出页面的方法
	 */
	//主要的操作便是主页面操作
	beego.Router(app.GetUrl("/page/index/main"), audio, "*:IndexPage")
	//输出瀑布流视图
	beego.Router(app.GetUrl("/page/pbl/main"), audio, "*:AudioPblPage")
	//登录页面
	beego.Router(app.GetUrl("/page/mit/login"), login, "*:LoginPage")
	//注册页面
	beego.Router(app.GetUrl("/page/mit/register"), login, "*:RegisterPage")
	//对应的登录注册协议的页面
	beego.Router(app.GetUrl("/page/mit/policy"), login, "*:PolicePage")
	//对应的 index 页面
	beego.Router(app.GetUrl("/page/mit/index"),page , "*:MitIndexPage")
	//上传信息的页面
	beego.Router(app.GetUrl("/page/upload/audio"), audio, "*:AudioUploadPage")
	//修改音频信息的页面
	beego.Router(app.GetUrl("/page/upload/update"), audio, "*:AudioUpdatePage")
	beego.Router(app.GetUrl("/page/upload/audio1"),audio,"*:Audio1")
	//详细页面
	beego.Router(app.GetUrl("/page/details/main"),audioComment , "*:PageDetails")
	beego.Router(app.GetUrl("/page/upload/test"),test_controller,"*:TestUploadPage")
	//政策页面
	beego.Router(app.GetUrl("/page/upload/police"),audio,"*:AudioPolice")
	/**
		对应的 api的接口
	 */
	//利用账号密码进行注册
	/*
		我认为相对应的注册方法与登录方法应该一样
		用户应该可以支持account ， wx ， 支付宝等等方式 ， 来进行创建
		虽然，account 的方法 ， 可以说应该是应该是 其他注册方法的基石 ， 但是如何设计那个是后话了
	*/
	beego.Router(app.GetUrl("/api/register/by/account"), login, "*:RegisterByAccount")
	//进行登录的api
	beego.Router(app.GetUrl("/api/login/by/account"), login, "*:LoginByAccount")
	//获取当前session中的user
	beego.Router(app.GetUrl("/api/login/current/user"), login, "*:CurrentUser")
	//修改密码
	beego.Router(app.GetUrl("/api/login/update/pwd"),login , "*:UpdatePwd")
	//根据对应的id来获取user信息
	beego.Router(app.GetUrl("/api/user/find/id"),user , "*:FindUserById")
	//根据对应的user信息来进行更新
	beego.Router(app.GetUrl("/api/user/update"),user , "*:UpdateUser")
	//更新用户名与用户头像
	beego.Router(app.GetUrl("/api/user/update/image/name"),user,"*:UpdateUserNameAndImage")
	//上传对应的 amr 文件
	beego.Router(app.GetUrl("/api/audio/upload/amr"),audio , "*:AudioUploadAmr")
	//进行对应的上传页面的 api
	beego.Router(app.GetUrl("/api/audio/upload"), audio, "*:AudioUpload")
	//最新版本的 audio 新建方法
	beego.Router(app.GetUrl("/api/audio/upload/new"),audio,"*:NewAudioUpload")
	//最新版本的上传方法
	beego.Router(app.GetUrl("/api/audio/update/new"),audio,"*:NewAudioUpdate")
	//对应的修改页面的api
	beego.Router(app.GetUrl("/api/audio/update"), audio, "*:AudioUpdate")
	//进行目标用户下的 所有音频文件的信息
	beego.Router(app.GetUrl("/api/audio/user"), audio, "*:SearchAudioByUserId")
	//删除目标的audio 的 信息
	beego.Router(app.GetUrl("/api/audio/delete"), audio, "*:Delete")
	//下面 进行对应的搜索
	beego.Router(app.GetUrl("/api/audio/search"), audio, "*:SearchByString")
	beego.Router(app.GetUrl("/api/audio/favorite"), audio, "*:Favorites")
	beego.Router(app.GetUrl("/api/audio/find/id"), audio, "*:FindAudioById")
	//根据对应的userId 来查询目标发布了多少的作品
	beego.Router(app.GetUrl("/api/audio/count"),audio,"*:AudioLen")
	//根据对应的id 来输出对应的 信息
	beego.Router(app.GetUrl("/api/audio/search/dashboard") , audio , "*:AudioByDashboard")
	//这里输出当前点击量最多的信息
	beego.Router(app.GetUrl("/api/audio/search/browse/most"),audio,"*:SearchMostBrowseAudio")
	//输出最近点击最火热的信息
	beego.Router(app.GetUrl("/api/audio/search/browse/date"),audio,"*:SearchMostDateBrowseAudio")
	//获取某个audio的 评论
	beego.Router(app.GetUrl("/api/comment/find/audio"),audioComment,"*:ByAudioId")
	//插入评论
	beego.Router(app.GetUrl("/api/comment/insert"),audioComment,"*:InsertComment")
	//根据对应的 comment 的 id 来删除评论
	beego.Router(app.GetUrl("/api/comment/delete/id"),audioComment,"*:DeleteCommentById")
	//书写对应的评论的回复
	beego.Router(app.GetUrl("/api/reply/comment/insert"), audioCommentReply,"*:InsertCommentReply")
	beego.Router(app.GetUrl("/api/reply/comment/find/comment") , audioCommentReply , "*:ByCommentId")
	//收藏的接口
	beego.Router(app.GetUrl("/api/collection/insert") ,collection ,"*:InsertCollection")
	//取消收藏的接口
	beego.Router(app.GetUrl("/api/collection/delete"),collection,"*:DeleteCollection")
	//取消收藏 根据对应的 userId + audioId + folderId
	beego.Router(app.GetUrl("/api/collection/delete/userandaudio"),collection,"*:DeleteColl")
	//根据对应的信息， 来获取对应的信息
	beego.Router(app.GetUrl("/api/collection/find/relationship"), collection , "*:FindByUserAndAudio")
	//根据对应的 收藏夹的 id 来搜索出对应的 收藏信息与其绑定的 audio 信息
	beego.Router( app.GetUrl("/api/collection/all/and"),collection,"*:SearchCollectionAndAudio")
	//转移收藏信息
	beego.Router( app.GetUrl("/api/collection/update"),collection 	, "*:UpdateCollection")
	//仪表板输出格式相关的接口
	//beego.Router(app.GetUrl("/api/dashborad/audio") , dashborad , "*:FindByAudioId")
	//对应的文件夹
	beego.Router( app.GetUrl("/api/collection/folder/all"), collectionFolder , "*:AllByUserId")
	//新建收藏夹
	beego.Router( app.GetUrl("/api/collection/folder/insert"),collectionFolder,"*:AddCollectionFolder")
	//更改收藏夹的对应的信息
	beego.Router( app.GetUrl( "/api/collection/folder/update") , collectionFolder ,"*:UpdateCollectionFolder")
	//删除收藏夹
	beego.Router( app.GetUrl("/api/collection/folder/delete") , collectionFolder , "*:DeleteCollectionFolder");
	//点赞
	beego.Router( app.GetUrl("/api/love/insert"),love,"*:InsertLove")
	//取消点赞
	beego.Router( app.GetUrl( "/api/love/delete"),love,"*:DeleteLove")
	//查询点赞关系
	beego.Router( app.GetUrl("/api/love/find"),love,"*:FindLove")
	//获取一个周期内的信息
	beego.Router( app.GetUrl("/api/dashboard/user/temp/find") , userTempDashboard ,"*:FindById")
	beego.Router( app.GetUrl("/api/dashboard/user/find") , userDashboard, "*:FindById")
	beego.Router( app.GetUrl("/api/dashboard/user/date/search"),userDateDashboard , "*:SearchByAudioId")
	beego.Router( app.GetUrl("/api/dashboard/user/date/search/time"),userDateDashboard ,"*:SearchByAudioIdByTime")
	//输出一个 目标用户在一个周期内 收到 的信息的总和
	beego.Router( app.GetUrl("/api/dashboard/total/temp") , dashboard , "*:TempFindUserId")
	//输出一个 目标用户在全部 收到 的信息的总和
	beego.Router( app.GetUrl("/api/dashboard/total/all") , dashboard , "*:AudioDashboardFindUserId")
	//输出资源文件
	//输出 音频的封面资源
	beego.Router( app.GetUrl( "/api/resource/image/audio") , resource , "*:ImageAudio")
	//输出用户的头像资源
	beego.Router( app.GetUrl("/api/resource/image/user"),resource , "*:ImageUser")
	//根据对应的路径 ， 来输出对应的地址
	beego.Router(app.GetUrl("/api/resource/image/path"),resource,"*:ResourcePath")
	//输出音频资源文件
	beego.Router( app.GetUrl("/api/resource/audio") , resource , "*:Audio")
	//根据hostResource 资源文件
	beego.Router( app.GetUrl( "/api/resource/host"),resource,"*:HostResource")
	//根据对应的 resource 来进行搜索
	beego.Router( app.GetUrl("/api/resource/id"),resource,"*:ResourceById")
	//获取其对应的信息
	beego.Router( app.GetUrl("/api/resource/image"),resource,"*:ResourceImage")
	//根据对应的表来进行获取对应的信息
	beego.Router( app.GetUrl( "/api/resource/audio/t"),resource,"*:ResourceAudio")
	//api测试输出
	beego.Router(app.GetUrl("/api/test"),test_controller,"*:ApiTest")
	//根据id来进行查询对应的信息
	beego.Router( app.GetUrl("/api/user/details/find"),userDetails ,"*:FindById")
	//更新对应的details信息
	beego.Router(app.GetUrl("/api/user/details/update"),userDetails , "*:UpdateUserDetails")
	//根据对应的userid ， 来搜索对应的信息
	beego.Router(app.GetUrl("/api/audio/history/search"),audioUserHistory,"*:SearchByUserId")
	// 登记
	beego.Router(app.GetUrl("/api/audio/history/add"),audioUserHistory,"*:AddCount")
	//新建对应的 audioFast
	beego.Router(app.GetUrl("/api/audio/fast/new"),audioFast,"*:NewAudioFast")
	/*
		对应的 web 的接口， 该接口的主要目的
		只是服务与 web 前端服务 交互的前端
	*/
	beego.Router(app.GetUrl("/web/test"),test_controller,"*:WebTest")
	//下面是测试api
	//测试对应的上传图片的功能是否有效
	beego.Router(app.GetUrl("/test/upload/image"), test_controller, "*:UploadImage")
	beego.Router(app.GetUrl("/test/session/set"), test_controller, "*:SetSession")
	beego.Router(app.GetUrl("/test/session/get"), test_controller, "*:GetSession")
	beego.Router(app.GetUrl("/test/audio/comments"),test_controller,"*:ApiAudioComment")
	beego.Router(app.GetUrl("/test/resource/redirec"),test_controller,"*:TestResourceRedirect")
	beego.Router(app.GetUrl("/test/dashboard/add"),test_controller,"*:TestDashboard")
	beego.Router(app.GetUrl("/test/audio/change"),test_controller,"*:AudioFormatChange")
	beego.Router(app.GetUrl("/test/resource/image/change"),test_controller,"*:ChangeResourceAudioFormatChange")

}
