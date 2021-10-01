package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PersonInfo struct {
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Id       string `json:"id"` // 身份证号码
	Phone    string `json:"phone"`
	Comment  string `json:"comment"`
	Room     string `json:"room"`
	FamilyId string `json:"familyId"`
}

type Resp struct {
	ErrMsg  string      `json:"errMsg"`
	ErrCode int         `json:"errCode"`
	Data    interface{} `json:"data"`
}

type Info struct {
	PersonInfos []PersonInfo
}

func NewInfo(personInfos []PersonInfo) *Info {
	return &Info{PersonInfos: personInfos}
}


type PersonResp struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Id      string `json:"id"` // 身份证号码
	Phone   string `json:"phone"`
	Comment string `json:"comment"`
	Room    string `json:"room"`
	Family  bool   `json:"family"`
}

type FamilyInfoResp struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Id      string `json:"id"` // 身份证号码
	Phone   string `json:"phone"`
	Comment string `json:"comment"`
	Room    string `json:"room"`
}

func (info *Info) HandlerInfo(ctx echo.Context) error {
	code := ctx.QueryParam("code")
	name := ctx.QueryParam("name")
	if code == "" {
		return returnErr(ctx, "请确认输入正确")
	}
	if name == "" { // 查询家属信息
		data, err := info.checkFamily(code)
		if err != nil {
			return returnErr(ctx, err.Error())
		}
		return returnOk(ctx, data)
	} else { // 查询个人信息
		data, err := info.checkPerson(code, name)
		if err != nil {
			return returnErr(ctx, err.Error())
		}
		return returnOk(ctx, data)
	}
}

func (info *Info) checkPerson(code string, name string) (*PersonResp, error) {
	for i := 0; i < len(info.PersonInfos); i++ {
		single := info.PersonInfos[i]
		if len(single.Id) < 4 {
			return nil, fmt.Errorf("身份证号码错误")
		}
		if single.Name == name && single.Id[len(single.Id)-4:] == code {
			personResp := PersonResp{
				Name:    single.Name,
				Gender:  single.Gender,
				Id:      single.Id,
				Phone:   single.Phone,
				Comment: single.Comment,
				Room:    single.Room,
			}
			if single.FamilyId != "" { // todo: list
				personResp.Family = true
			}
			return &personResp, nil
		}
	}

	return nil, fmt.Errorf("查询不到对应的信息，请确认输入正确")
}

func (info *Info) checkFamily(code string) ([]FamilyInfoResp, error) {
	res := make([]FamilyInfoResp, 0)
	for i := 0; i < len(info.PersonInfos); i++ {
		single := info.PersonInfos[i]
		if single.Id == code {
			if single.FamilyId != "" { // 查询对应的 id 出来
				data := info.getPersonInfo(single.FamilyId)
				if data != nil {
					res = append(res, FamilyInfoResp{
						Name:    data.Name,
						Gender:  data.Gender,
						Id:      data.Id,
						Phone:   data.Phone,
						Comment: data.Comment,
						Room:    data.Room,
					})
				}
			}
		}
	}

	return res, nil
}

func (info *Info) getPersonInfo(code string) *PersonInfo {
	for i := 0; i < len(info.PersonInfos); i++ {
		single := info.PersonInfos[i]
		if single.Id == code {
			return &single
		}
	}

	return nil
}

func returnErr(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusOK, &Resp{
		ErrMsg:  msg,
		ErrCode: -1,
	})
}

func returnOk(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, &Resp{
		Data: data,
	})
}
