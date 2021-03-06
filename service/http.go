package service

import (
	"os"
	"io"
	"yinji/models/bean"
	"mime/multipart"
	"path"
	"strconv"
)

type HttpFileService struct {
	fileService *FileService
}

const RESOURCE_BAST_PATH = "static"

const URL_AUDIO_PATH = "resources/audio"

const URL_IMAGE_AUDIO_PATH = "resources/image/audio"

const URL_IMAGE_USER_PATH = "resources/image/user"

func ( self *HttpFileService ) GetServerPath() string {
	var serverPath , _ = os.Getwd()
	//返回对应的 基础时间信息
	return serverPath + "/" + RESOURCE_BAST_PATH
}

func ( self *HttpFileService ) BuildServerPath( path string ) string {
	var serverPath = self.GetServerPath()
	return serverPath + "/" + path
}

/**
	获取对应的amr 文件
*/
func ( self *HttpFileService ) GetAmrPath( file string ) string {
	return self.GetAudioPath( file )
}

//得到的是返回的路径 ， 可成功 ， 可不成功
func ( self * HttpFileService ) GetImagUserePath( path string ) string {
	return  URL_IMAGE_USER_PATH + "/" + path
}

//得到的是返回的路径 ， 可成功 ， 可不成功
func ( self * HttpFileService ) GetImagePath( path string ) string {
	var serverPath = self.GetServerPath()
	return serverPath + "/" + URL_IMAGE_AUDIO_PATH + "/" + path
}

/**
	获取对应的 auido 的 相对应的 资源地址
 */
func ( self *HttpFileService ) GetAudioRelaPath( path  string ) string {
	return URL_AUDIO_PATH +"/" + path
}

func ( self * HttpFileService ) GetAudioPath( path string ) string{
	var serverPath = self.GetServerPath()

	var absPath = serverPath + "/" + URL_AUDIO_PATH +"/" + path

	return absPath
}

func ( self *HttpFileService) GetAudioFileName( audio *bean.Audio ) string {
	return audio.Url
}

 func ( self *HttpFileService ) BuildFileName( id int64 , header *multipart.FileHeader ) string {
	var imageSuffix = path.Ext( header.Filename )
	var imageName = strconv.FormatInt( id ,10) + imageSuffix
	return imageName
}

func ( self *HttpFileService ) BuildAudioFileName( audio *bean.Audio , header *multipart.FileHeader ) string {
	var imageSuffix = path.Ext( header.Filename )
	var imageName = strconv.FormatInt( audio.Id ,10) + imageSuffix
	return imageName
}

/**
	利用绝对路径上传
 */
func ( self *HttpFileService ) Upload( path string , reader io.Reader) ( string , error ){

	var serverPath = self.BuildServerPath(path)
	var file , getFileErr = self.fileService.GetFile( serverPath )

	if getFileErr != nil {
		//倘若查找文件出现错误 ， 那么我们便开始返回相对应的结果
		return "" , getFileErr
	}

	var fileAssembly = FileAssembly{ file }

	self.fileService.Write( fileAssembly , reader )

	defer fileAssembly.Release()

	return path , nil
}

func ( self *HttpFileService ) UploadAudio( path string  , reader io.Reader)  ( string , error )   {

	var absPath = self.GetAudioPath( path )

	var file , err = self.fileService.GetFile(absPath)

	if err != nil {
		return "" ,  err
	}

	var fileAssembly  = FileAssembly{ file  }

	var writeFileErr = self.fileService.Write( fileAssembly , reader )

	if writeFileErr != nil {
		return "" , writeFileErr
	}

	defer fileAssembly.Release()

	var fileName = URL_AUDIO_PATH + "/" + path
	return  fileName , nil
}

//上传图片的
func ( self *HttpFileService) UploadImage( path string , reader io.Reader ) ( string , error ){
	//使用该方法，上传图片，返回的是 对应的路径还有 error
	var absPath = self.GetImagePath( path )
	var file , getFileErr = self.fileService.GetFile( absPath )

	if getFileErr != nil {
		//倘若查找文件出现错误 ， 那么我们便开始返回相对应的结果
		return "" , getFileErr
	}

	var fileAssembly = FileAssembly{ file }

	self.fileService.Write( fileAssembly , reader )

	var fileName = URL_IMAGE_AUDIO_PATH + "/" + path

	defer fileAssembly.Release()

	return fileName , nil
}

func ( self *HttpFileService ) Delete( path string) error {
	var absPath = self.BuildServerPath( path )
	return self.fileService.Delete( absPath );
}

func ( self *HttpFileService ) CloseMultipart(  file multipart.File  ) error {
	if file != nil {
		return file.Close()
	}

	return nil

}

var HTTP_SERVICE_INSTANCE = &HttpFileService{GetFileServiceInstance()}

func GetHttpFileServiceInstance() *HttpFileService{
	return HTTP_SERVICE_INSTANCE
}

