package Implemention

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type SubjectForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentId    int    `json:"parentId"`
	Level       int    `json:"level"`
}

type Subject struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Parent      string    `json:"parent"`
	ParentId    int       `json:"parentId"`
}

var BIG = 1
var SMALL = 2

func (api *Api) CreateBigSubject(name string) (err error) {
	var form SubjectForm
	form.Name = name
	form.Description = ""
	form.Level = BIG
	resp, err := createSubject(form)
	if err != nil {
		return
	}
	api.resp = *resp
	body := api.resp.Body()
	storeSubject(body)
	return
}



func (api *Api) CreateSmallSubject(name, parentName string) (err error) {
	var form SubjectForm
	form.Name = name
	parentId, err2 := getParentId(parentName)
	if err2 != nil {
		return
	}
	form.ParentId = parentId
	form.Level = SMALL
	resp, err := createSubject(form)
	if err != nil {
		return
	}
	api.resp = *resp
	return
}

func getParentId(name string) (int, error) {
	subject, err := GetSubjectByName(name)
	if err != nil {
		return 0, err
	}
	return subject.Id, nil
}

func GetSubjectByName(name string) (Subject, error) {
	var key = subjectStoreKey(name)
	subjectJSON := StoreGet(key)
	println(subjectJSON)
	var subject Subject
	err := json.Unmarshal([]byte(subjectJSON), &subject)
	return subject, err
}

func subjectStoreKey(name string) string {
	return "subject_" + name
}

func createSubject(form SubjectForm) (*resty.Response, error) {
	jsonBytes, errOfFormat := json.Marshal(form)
	if errOfFormat != nil {
		return nil, errOfFormat
	}
	return CreateRequest().
		SetBody(string(jsonBytes)).
		Post("http://localhost/api/subjects")
}

func storeSubject(body []byte) {
	subjectRet := parse(body)
	StorePut(subjectStoreKey(subjectRet.Name), string(body))
}

func parse(body []byte) Subject {
	var subjectRet Subject
	err := json.Unmarshal(body, &subjectRet)
	if err != nil {
		fmt.Printf("%v", err)
		return Subject{}
	}
	return subjectRet
}