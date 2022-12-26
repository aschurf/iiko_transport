package iikotransport

import (
	"encoding/json"
	"fmt"
	"github.com/aschurf/iiko_transport/internal/dishes"
	"log"
)

type IikoLogin struct {
	Login string
	Token string
}

func (login *IikoLogin) Auth() error {

	postBody := map[string]interface{}{
		"apiLogin": login.Login,
	}

	iikoRequest := IikoRequest{
		EndPoint: "1/access_token",
		Method:   "POST",
		Payload:  postBody,
	}

	res, err := iikoRequest.Send()
	if err != nil {
		log.Println(err)
	}

	//Если в ответе есть ошибка
	_, isError := res["errorDescription"]
	if isError {
		str := fmt.Sprint(res["errorDescription"])
		return fmt.Errorf("IIKOerror: %s", str)
	}

	login.Token = fmt.Sprint(res["token"])
	return nil
}

// GetOrganizations получить доступные организации
func (login *IikoLogin) GetOrganizations() {

	iikoRequest := IikoRequest{
		EndPoint: "1/organizations",
		Method:   "POST",
		Bearer:   login.Token,
		Payload:  make(map[string]interface{}),
	}

	result, er := iikoRequest.Send()
	if er != nil {
		fmt.Println(er)
	}

	fmt.Println(result)

	//for _, item := range result["organizations"].([]interface{}) {
	//	fmt.Printf("%v", item.(map[string]interface{})["id"])
	//}

	//var re map[string]string = result["organizations"].(map[string]string)
	//fmt.Println(re)

	//res, ok := result["organizations"].(map[string]string)
	//if !ok {
	//	fmt.Println("ERROR!")
	//}
	//
	//for _, v := range res {
	//	fmt.Println(v)
	//}

}

// GetMenu получить все доступные меню
func (login *IikoLogin) GetMenu() {

	iikoRequest := IikoRequest{
		EndPoint: "2/menu",
		Method:   "POST",
		Bearer:   login.Token,
		Payload:  make(map[string]interface{}),
	}

	result, er := iikoRequest.Send()
	if er != nil {
		fmt.Println(er)
	}

	fmt.Println(result)
}

// GetDishesByMenuId Получить состав меню по его ID и ID организаций
func (login *IikoLogin) GetDishesByMenuId(organizationIds []string, externalMenuId int) (dishes.Menu, error) {
	iikoRequest := IikoRequest{
		EndPoint: "2/menu/by_id",
		Method:   "POST",
		Bearer:   login.Token,
		Payload: map[string]interface{}{
			"externalMenuId":  externalMenuId,
			"organizationIds": organizationIds,
		},
	}

	result, er := iikoRequest.Send()
	if er != nil {
		fmt.Println(er)
	}

	jsonbody, err := json.Marshal(result)
	if err != nil {
		return dishes.Menu{}, err
	}

	menu := dishes.Menu{}
	if err := json.Unmarshal(jsonbody, &menu); err != nil {
		return dishes.Menu{}, err
	}

	return menu, nil
}
