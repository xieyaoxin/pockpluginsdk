package cqtt

import (
	"errors"
	"pock_plugins/backend/log"
	"pock_plugins/backend/model"
	util "pock_plugins/backend/utils"
	"strconv"
	"strings"
)

var PetRepositoryImpl4CQTT = &petRepositoryImpl4CQTT{}

type petRepositoryImpl4CQTT struct{}

func (instance *petRepositoryImpl4CQTT) GetCarriedPets() ([]*model.Pet, error) {
	Pets := []*model.Pet{}
	params := util.InitParam()
	result := CallServerGetInterface("function/Pets_Mod.php", params)
	lines := strings.Split(result, "\n")
	// 主站宠物
	for lineNumber := range lines {
		line := lines[lineNumber]
		if strings.Contains(line, "cursor:hand;opacity: 1") {
			id := strings.Replace(strings.Split(strings.Split(line, "Setbb")[1], ",")[0], "(", "", 1)
			pet, err := instance.GetPetDetail(id)
			if err != nil {
				log.Error("获取宠物信息失败 宠物ID： ", id)
				return nil, err
			}
			pet.IsBattle = true
			Pets = append(Pets, pet)
			break
		}
	}
	// 非主站
	for lineNumber := range lines {
		line := lines[lineNumber]
		if strings.Contains(line, "cursor:hand;opacity: 0.5") {
			id := strings.Replace(strings.Split(strings.Split(line, "Setbb")[1], ",")[0], "(", "", 1)
			pet, err := instance.GetPetDetail(id)
			if err != nil {
				log.Error("获取宠物信息失败 宠物ID： ", id)
				return nil, err
			}
			Pets = append(Pets, pet)
			break
		}
	}
	return Pets, nil
}

func (*petRepositoryImpl4CQTT) GetPetDetail(PetId string) (*model.Pet, error) {
	params := util.InitParam()
	params["id"] = PetId
	result := CallServerGetInterface("function/mcbbshow.php", params)
	lines := strings.Split(result, "\n")
	//petList := []*model.Pet{}
	pet := model.Pet{Id: PetId}
	for _, line := range lines {
		if strings.Contains(line, "font-family") {
			petNameString := strings.Split(strings.Replace(line, "<", ">", -1), ">")[2]
			pet.Name = petNameString
		}
		if strings.Contains(line, "成长：") {
			cc := strings.Split(strings.Replace(line, "<", "：", -1), "：")[1]
			pet.Cc, _ = strconv.ParseFloat(cc, 64)
		}
		if strings.Contains(line, "等级：") {
			level := strings.Replace(strings.Split(strings.Replace(line, "<", "：", -1), "：")[1], " ", "", -1)
			pet.Level, _ = strconv.ParseInt(level, 10, 64)
		}
	}
	return &pet, nil
}

func (*petRepositoryImpl4CQTT) GetPetSkillList(PetId string) ([]*model.Skill, error) {
	SkillList := []*model.Skill{}
	params := util.InitParam()
	params["pid"] = PetId
	result := CallServerGetInterface("function/Pets_Mod.php", params)
	lines := strings.Split(result, "\n")
	for _, line := range lines {
		// todo 获取技能方式需要确认下
		// 			<span class="skill" onmouseover="showBox(44,'龙啸九天')" onmouseout="hidBox(this)" onclick="copyWord('龙啸九天');"> 龙啸九天&nbsp;&nbsp;10 级 </span>
		if strings.Contains(line, "span class=\"skill\"") {
			if strings.Contains(line, "升级") {
				Skill := strings.Split(strings.Replace(line, "<", ">", -1), ">")[4]
				SkillName := strings.Split(strings.Replace(Skill, " ", "", -1), "&nbsp;")[0]
				SkillLevel := strings.Split(strings.Replace(Skill, " ", "", -1), "&nbsp;")[2]
				SkillId := strings.Split(strings.Replace(strings.Replace(line, ")", ",", -1), "(", ",", -1), ",")[1]
				SkillList = append(SkillList, &model.Skill{
					SkillName:  SkillName,
					SkillLevel: SkillLevel,
					SkillId:    SkillId,
				})
			} else {
				Skill := strings.Split(strings.Replace(line, "<", ">", -1), ">")[2]
				SkillName := strings.Split(strings.Replace(Skill, " ", "", -1), "&nbsp;")[0]
				SkillLevel := strings.Split(strings.Replace(Skill, " ", "", -1), "&nbsp;")[2]
				SkillId := strings.Split(strings.Replace(strings.Replace(line, ")", ",", -1), "(", ",", -1), ",")[1]
				SkillList = append(SkillList, &model.Skill{
					SkillName:  SkillName,
					SkillLevel: SkillLevel,
					SkillId:    SkillId,
				})
			}

		}

	}
	return SkillList, nil
}

func (*petRepositoryImpl4CQTT) GetFarmedPets() ([]*model.Pet, error) {
	//TODO implement me
	panic("implement me")
}

func (*petRepositoryImpl4CQTT) SetBattlePet(PetId string) error {
	params := util.InitParam()
	params["id"] = PetId
	params["op"] = "z"
	result := CallServerGetInterface("function/mcGate.php", params)
	log.Info("设置主站宠物 %s %s", PetId, result)
	if result != "更改主战宝宝成功!" && result != "已经是主战！" {
		return errors.New(result)
	}
	return nil
}

func (instance *petRepositoryImpl4CQTT) CarryPet(PetId string) error {
	//TODO implement me
	panic("implement me")
}

func (instance *petRepositoryImpl4CQTT) SavePet(PetId string) error {
	params := util.InitParam()
	params["id"] = PetId
	params["op"] = "s"
	result := CallServerGetInterface("function/mcGate.php", params)
	log.Info("寄存宠物 %s 操作结果 %s:", PetId, result)
	if result != "操作成功!" {
		return errors.New("寄样宠物失败")
	}
	return nil
}
