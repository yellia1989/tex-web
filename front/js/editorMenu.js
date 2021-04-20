
const E = window.wangEditor

const { $, BtnMenu, DropListMenu, PanelMenu, DropList, Panel, Tooltip } = E

class MyColorMenu extends BtnMenu {
    constructor(editor) {
        // data-title属性表示当鼠标悬停在该按钮上时提示该按钮的功能简述
        const $elem = E.$(
            `<div class="w-e-menu" data-title="MyColor">
                <div style="margin:0 auto" id="colorPicker"></div>
            </div>`
        )
        super($elem, editor)
        layui.use('colorpicker', function(){
            var colorpicker = layui.colorpicker;
            //渲染
            colorpicker.render({
                elem: '#colorPicker'  //绑定元素
                ,done: function(color){
                    if(editor.selection.isSelectionEmpty()){
                        return true;
                    }
                    editor.cmd.do('foreColor', color);
                },
                size:'sm',
            });
        });
    }

    /**
     * 菜单点击事件
     */
    clickHandler() {

    }

    tryChangeActive() {

    }

}


// 注册菜单
E.registerMenu('myColorKey', MyColorMenu);

const fontSizeMap = {'1': '10', '2': '13', '3': '16', '4': '18', '5': '24', '6': '32', '7': '48',};

function searchData(jsonData,outputObject) {
    let endTag = "";
    switch (typeof jsonData){
        case "string":
            outputObject.result+=jsonData;
            break;
        case "object":
            switch (jsonData.tag){
                case "br":
                    endTag += "\n ";
                    break;
                case "p":
                    endTag += "\n ";
                    break;
                case "span":
                    for(let i in jsonData.attrs){
                        if(jsonData.attrs[i].name==="style"){
                            let tempRegArr;
                            tempRegArr = /font-size: ([0-9]*)px/.exec(jsonData.attrs[i].value);
                            if(tempRegArr!=null){
                                outputObject.result += "[SIZE="+tempRegArr[1]+"]";
                                endTag = "[/SIZE]";
                            }
                        }
                    }
                    break;
                case "font":
                    for(let i in jsonData.attrs){
                        if(jsonData.attrs[i].name==="color"){
                            outputObject.result += "[color="+jsonData.attrs[i].value+"]";
                            endTag += "[/color]";
                        }else if(jsonData.attrs[i].name==="size"){
                            outputObject.result += "[SIZE="+fontSizeMap[jsonData.attrs[i].value]+"]";
                            endTag += "[/SIZE]";
                        }
                    }
                    break;
            }
            for(let i in jsonData.children){
                searchData(jsonData.children[i],outputObject);
            }
            if(endTag === "\n "){
                let regex = new RegExp("\\[/[a-zA-z]*\\]$");
                let arr = outputObject.result.match(regex);
                if(arr!=null){
                    outputObject.result = outputObject.result.replace(regex,endTag+arr[0]);
                }else {
                    outputObject.result+=endTag;
                }

            }else {
                outputObject.result+=endTag;
            }
            break;
    }
}

function ExecuteJsonData(jsonData) {
    let outputObject = { "result":""};
    if(jsonData.length>0){
        for(let i=0;i<jsonData.length;i++){
            searchData(jsonData[i],outputObject);
        }
    }
    return outputObject.result;
}

