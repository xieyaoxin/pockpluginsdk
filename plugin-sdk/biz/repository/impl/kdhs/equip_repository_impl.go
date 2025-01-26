package kdhs

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	util "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
	"strconv"
	"strings"
)

var EquipRepositoryImplKDHSInstance = &equipRepositoryImplKDHS{}

type equipRepositoryImplKDHS struct{}

func (inst *equipRepositoryImplKDHS) GetEquipList(PetId string) []*model.Equip {
	params := util.InitParam()
	params["pid"] = PetId
	result := CallServerGetInterface("function/Pets_Mod.php", params)
	lines := strings.Split(result, "\n")
	equips := []*model.Equip{}
	for _, line := range lines {
		if strings.Contains(line, "takeoff") && strings.Contains(line, "ondblclick") {
			Messages := strings.Split(strings.Replace(strings.Replace(strings.Replace(strings.Replace(line, "<", ",", -1), ">", ",", -1), "(", ",", -1), ")", ",", -1), ",")
			equipId := Messages[14]
			equipPid := Messages[12]
			equipName := Messages[16]
			equip := &model.Equip{
				EquipId:  equipId,
				EquipPId: equipPid,
				Name:     equipName,
			}
			equips = append(equips, equip)
		}
	}
	return equips
}

func (equipRepositoryImplKDHS) StrengthEquip(EquipId string, EquipPid string, DragonBallId string) bool {
	//TODO implement me
	panic("implement me")
}

func (equipRepositoryImplKDHS) OffEquip(PetId string, EquipPid string) bool {
	///function/offprops.php?id=10103&bid=8178204
	params := util.InitParam()
	params["bid"] = PetId
	params["id"] = EquipPid
	result := CallServerGetInterface("/function/offprops.php", params)
	return result == "2"
}

func (equipRepositoryImplKDHS) GetEquip(id string) *model.Equip {
	params := util.InitParam()
	params["id"] = id
	params["sign"] = "1"
	params["type"] = "2"
	result := CallServerGetInterface("function/getPropsInfo.php", params)
	lines := strings.Split(result, "\n")
	equip := model.Equip{}
	for _, line := range lines {
		if !strings.Contains(line, "强化") {
			continue
		}
		// <td style="background:#1F1F30;filter:Alpha(opacity=90);"><font color="#0067CB"><b>圣灵靴&nbsp;+15</b></font><br/><font color=#A8A7A4>可交易</font><br/><font color=#FEFDFA>永久</font><br /><font color=#FEFDFA>脚部装备&nbsp(可强化)</font><br/><font color=#FEFDFA class="line">+12% 攻击力 <font color= red>+20%</font></font><br/><font color=#0067CB>+50000 魔法</font><br/><font color=#0067CB>伤害加深 25%</font><br/><font color=#14FD10>卡槽数：1/1</font><br/><font color="red">宝石效果：对敌人造成的伤害增加60%</font><br/><font color=#FED625>仙侠套装(10/10)</font><br/><font color=#14FD10>(2)套装：+60000 生命</font><br/><font color=#14FD10>(4)套装：+60% 防御</font><br/><font color=#14FD10>(5)套装：偷取伤害的60%转化为生命</font><br/><font color=#14FD10>(6)套装：+150% 攻击</font><br/><font color=#14FD10>(7)套装：伤害抵消 70%</font><br/><font color=#14FD10>(8)套装：+150% 命中</font><br/><font color=#14FD10>(9)套装：战斗等待时间减少3秒</font><br/><font color=#14FD10>(10)套装：伤害加深 150%</font><br/><font color=#FEFDFA>装备：圣灵飞侠的专属脚部装备，由魔力创造的靴子，充满神奇的魔法效果。只能使用幻灵冰晶才能进行强化。</font><br/></td><td width=5 background=../images/ui/tips/border4_r.gif></td>
		propertyList := strings.Split(strings.Replace(line, "</b>", "<b>", -1), "<b>")
		nameProperty := strings.Split(propertyList[1], "&nbsp;")
		name := nameProperty[0]
		var Strengthen int64 = 0
		if len(nameProperty) > 1 {
			Strengthen, _ = strconv.ParseInt(strings.Replace(nameProperty[1], "+", "", -1), 10, 64)
		}
		// <font color=#FEFDFA>脚部装备&nbsp(可强化)</font>
		positionPropertys := strings.Split(strings.Replace(strings.Replace(line, "<br/>", "<br>", -1), "<br />", "<br>", -1), "<br>")
		positionProperty := strings.Split(strings.Split(strings.Split(positionPropertys[3], ">")[1], "<")[0], "&nbsp")
		effect := model.Effect{}
		for _, sj := range positionPropertys {
			if strings.Contains(sj, "宝石效果") {
				EffectProperty := strings.Split(strings.Split(sj, "</font")[0], "宝石效果：")[1]

				if strings.Contains(EffectProperty, "对敌人造成的伤害增加") {
					effect.EffectType = "加深"
					fig, _ := strconv.ParseInt(strings.Replace(strings.Replace(EffectProperty, "对敌人造成的伤害增加", "", -1), "%", "", -1), 10, 64)
					effect.Figure = fig
				}
				if strings.Contains(EffectProperty, "会心一击率增加") {
					effect.EffectType = "暴击"
					fig, _ := strconv.ParseInt(strings.Replace(strings.Replace(EffectProperty, "会心一击率增加", "", -1), "%", "", -1), 10, 64)
					effect.Figure = fig
				}

				if strings.Contains(EffectProperty, "增加命中") {
					effect.EffectType = "命中"
					fig, _ := strconv.ParseInt(strings.Replace(strings.Replace(EffectProperty, "增加命中", "", -1), "%", "", -1), 10, 64)
					effect.Figure = fig
				}
			}

		}
		equip = model.Equip{
			EquipId:    id,
			Strengthen: Strengthen,
			Position:   positionProperty[0],
			Name:       name,
			Effect:     effect,
		}
	}
	return &equip
}
