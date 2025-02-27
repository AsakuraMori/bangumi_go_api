package bgm_v0

import (
	"context"
	"github.com/AsakuraMori/bangumi_go_api/model/items"
	"github.com/AsakuraMori/bangumi_go_api/pkg/errno"
)

//GetSubject 根据id获得单个条目的详细信息
func (cli *Client) GetSubject(ctx context.Context, authToken string, subjectId string, body []string) (*items.SubjectBgm, error) {

	//var subjectList []*items.SubjectBgm
	//
	//for {
	var resp = &items.SubjectBgm{}
	param := map[string]string{
		//"pageIndex": 1,
		//"pageSize":  100,
	}
	if subjectId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject id is required!")
	}
	err := cli.GET(ctx, "/v0/subjects/"+subjectId, authToken, 0, param, body, resp)
	if err != nil {
		return nil, err
	}
	//subjectList = append(subjectList, resp)
	//}
	return resp, nil
}

//GetSubjectPersons 获取条目下persons
func (cli Client) GetSubjectPersons(ctx context.Context, authToken string, subjectId string) ([]*items.RelatedPerson, error) {
	var resp = make([]*items.RelatedPerson, 0)
	if subjectId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject id is required!")
	}
	param := map[string]string{}
	err := cli.GET(ctx, "/v0/subjects/"+subjectId+"/persons", authToken, 0, param, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetSubjectCharacters 获取条目下characters
func (cli Client) GetSubjectCharacters(ctx context.Context, authToken string, subjectId string) ([]*items.RelatedCharacter, error) {
	var resp = make([]*items.RelatedCharacter, 0)
	if subjectId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject id is required!")
	}
	param := map[string]string{}
	err := cli.GET(ctx, "/v0/subjects/"+subjectId+"/characters", authToken, 0, param, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetSubjectRelations 获取条目下关联的条目subjects
func (cli Client) GetSubjectRelations(ctx context.Context, authToken string, subjectId string) ([]*items.RelatedSubject, error) {
	var resp = make([]*items.RelatedSubject, 0)
	if subjectId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject id is required!")
	}
	param := map[string]string{}
	err := cli.GET(ctx, "/v0/subjects/"+subjectId+"/subjects", authToken, 0, param, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetCalendar 获取每日放送的条目
func (cli Client) GetCalendar(ctx context.Context) ([]*items.CalendarItem, error) {
	var resp = make([]*items.CalendarItem, 0)
	param := map[string]string{}
	err := cli.GET(ctx, "/calendar", "", 0, param, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
