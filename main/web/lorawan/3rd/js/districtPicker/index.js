$(function (){
		initComplexArea('seachprov', 'seachcity', 'seachdistrict', area_array, sub_array, '0', '0', '0');
});
		
function getAreaID(){//获取地区码封装函数
	var area = 0;
	if($("#seachdistrict").val() != "0"){
		area = $("#seachdistrict").val();
	}else if ($("#seachcity").val() != "0"){
		area = $("#seachcity").val();
	}else{
		area = $("#seachprov").val();
	}
	return area;
}

function showAreaID() {
	var areaID = getAreaID();//获取地区码
	var areaName = getAreaNamebyID(areaID) ;//获取地区名

	//如果是直辖市，则将其地区码增加到6位
	var areaIDindex1 = areaID.substring(0, 2);
	if(areaIDindex1 == "11" || areaIDindex1 == "12" || areaIDindex1 == "31" 
		|| areaIDindex1 == "50" || areaIDindex1 == "71" || areaIDindex1 == "81" || areaIDindex1 ==   "82"){
		var areaIDindex2 = areaID.slice(-2);
		areaID = areaIDindex1 + "01" +areaIDindex2;
	}
	
	$("#showAreaInfo").text("您选择的地区名："+ areaName + "，其地区码为："+ areaID);//点击获取提示
}


function getAreaNamebyID(areaID){//根据地区码查询地区名
	var areaName = "";
	if(areaID.length == 2){
		areaName = area_array[areaID];
	}else if(areaID.length == 4){
		var index1 = areaID.substring(0, 2);
		areaName = area_array[index1] + " " + sub_array[index1][areaID];
	}else if(areaID.length == 6){
		var index1 = areaID.substring(0, 2);
		var index2 = areaID.substring(0, 4);
		areaName = area_array[index1] + " " + sub_array[index1][index2] + " " + sub_arr[index2][areaID];
	}
	return areaName;
}


