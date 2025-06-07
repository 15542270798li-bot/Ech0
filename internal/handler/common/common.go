package handler

import (
	"github.com/gin-gonic/gin"
	commonModel "github.com/lin-snow/ech0/internal/model/common"
	service "github.com/lin-snow/ech0/internal/service/common"
	errorUtil "github.com/lin-snow/ech0/internal/util/err"
	"net/http"
)

type CommonHandler struct {
	commonService service.CommonServiceInterface
}

func NewCommonHandler(commonService service.CommonServiceInterface) *CommonHandler {
	return &CommonHandler{
		commonService: commonService,
	}
}

func (commonHandler *CommonHandler) UploadImage(ctx *gin.Context) {
	// 提取上传的 File数据
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, commonModel.Fail[string](errorUtil.HandleError(&commonModel.ServerError{
			Msg: commonModel.INVALID_REQUEST_BODY,
			Err: err,
		})))
		return
	}

	// 提取userid
	userId := ctx.MustGet("userid").(uint)

	// 调用 CommonService 上传文件
	imageUrl, err := commonHandler.commonService.UploadImage(userId, file)
	if err != nil {
		ctx.JSON(http.StatusOK, commonModel.Fail[string](errorUtil.HandleError(&commonModel.ServerError{
			Msg: "",
			Err: err,
		})))
		return
	}

	ctx.JSON(http.StatusOK, commonModel.OK[string](imageUrl, commonModel.UPLOAD_SUCCESS))
}

func (commonHandler *CommonHandler) DeleteImage(ctx *gin.Context) {
	userId := ctx.MustGet("userid").(uint)

	var imageDto commonModel.ImageDto
	if err := ctx.ShouldBindQuery(&imageDto); err != nil {
		ctx.JSON(http.StatusOK, commonModel.Fail[string](errorUtil.HandleError(&commonModel.ServerError{
			Msg: commonModel.INVALID_REQUEST_BODY,
			Err: err,
		})))
		return
	}

	if err := commonHandler.commonService.DeleteImage(userId, imageDto.URL, imageDto.SOURCE); err != nil {
		ctx.JSON(http.StatusOK, commonModel.Fail[string](errorUtil.HandleError(&commonModel.ServerError{
			Msg: "",
			Err: err,
		})))
		return
	}

	ctx.JSON(http.StatusOK, commonModel.OK[any](nil, commonModel.DELETE_SUCCESS))
}

func (commonHandler *CommonHandler) GetStatus(ctx *gin.Context) {
	_, err := commonHandler.commonService.GetSysAdmin()
	if err != nil {
		ctx.JSON(http.StatusOK, commonModel.OKWithCode[any](nil, commonModel.InitInstallCode, commonModel.SIGNUP_FIRST))
		return
	}

	status, err := commonHandler.commonService.GetStatus()
	if err != nil {
		ctx.JSON(http.StatusOK, commonModel.Fail[string](errorUtil.HandleError(&commonModel.ServerError{
			Msg: "",
			Err: err,
		})))
		return
	}

	ctx.JSON(http.StatusOK, commonModel.OK[commonModel.Status](status, commonModel.GET_STATUS_SUCCESS))
}

func (commonHandler *CommonHandler) GetHeatMap(ctx *gin.Context) {
	// 调用 Service 层获取热力图数据
	heatMap, err := commonHandler.commonService.GetHeatMap()
	if err != nil {
		ctx.JSON(http.StatusOK, commonModel.Fail[string](errorUtil.HandleError(&commonModel.ServerError{
			Msg: "",
			Err: err,
		})))
		return
	}

	ctx.JSON(http.StatusOK, commonModel.OK[[]commonModel.Heapmap](heatMap, commonModel.GET_HEATMAP_SUCCESS))
}

func (commonHandler *CommonHandler) GetRss(ctx *gin.Context) {
	atom, err := commonHandler.commonService.GenerateRSS(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, commonModel.Fail[string](errorUtil.HandleError(&commonModel.ServerError{
			Msg: "",
			Err: err,
		})))
		return
	}

	ctx.Data(http.StatusOK, "application/rss+xml; charset=utf-8", []byte(atom))
}

func (commonHandler *CommonHandler) UploadAudio(ctx *gin.Context) {
	// 提取userid
	userId := ctx.MustGet("userid").(uint)

	// 提取上传的 File数据
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, commonModel.Fail[string](errorUtil.HandleError(&commonModel.ServerError{
			Msg: commonModel.INVALID_REQUEST_BODY,
			Err: err,
		})))
		return
	}

	audioUrl, err := commonHandler.commonService.UploadMusic(userId, file)
	if err != nil {
		ctx.JSON(http.StatusOK, commonModel.Fail[string](errorUtil.HandleError(&commonModel.ServerError{
			Msg: "",
			Err: err,
		})))
		return
	}

	ctx.JSON(http.StatusOK, commonModel.OK[string](audioUrl, commonModel.UPLOAD_SUCCESS))
}

func (commonHandler *CommonHandler) DeleteAudio(ctx *gin.Context) {
	// 提取userid
	userId := ctx.MustGet("userid").(uint)

	if err := commonHandler.commonService.DeleteMusic(userId); err != nil {
		ctx.JSON(http.StatusOK, commonModel.Fail[string](errorUtil.HandleError(&commonModel.ServerError{
			Msg: "",
			Err: err,
		})))
		return
	}

	ctx.JSON(http.StatusOK, commonModel.OK[any](nil, commonModel.DELETE_SUCCESS))
}

func (commonHandler *CommonHandler) GetPlayMusic(ctx *gin.Context) {
	musicUrl := commonHandler.commonService.GetPlayMusicUrl()

	ctx.JSON(http.StatusOK, commonModel.OK[string](musicUrl, commonModel.GET_MUSIC_URL_SUCCESS))
}

func (commonHandler *CommonHandler) PlayMusic(ctx *gin.Context) {
	commonHandler.commonService.PlayMusic(ctx)
}
