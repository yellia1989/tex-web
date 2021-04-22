
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

const fontSizeMap = {'1': '16', '2': '18', '3': '20', '4': '22', '5': '24', '6': '26', '7': '28',};

function getDefaultEditor(id){
    let editor = new E(id)
    editor.config.menus = [
        'undo',
        'redo',
        'fontSize',
        'foreColor',
        'myColorKey',
    ]
    editor.config.colors = [
        '#955619',
        '#ff0000',
        '#cd7540',
    ]
    editor.config.fontSizes = {
        'x-small': { name: '16px', value: '1' },
        'small': { name: '18px', value: '2' },
        'normal': { name: '20px', value: '3' },
        'large': { name: '22px', value: '4' },
        'x-large': { name: '24px', value: '5' },
        'xx-large': { name: '26px', value: '6' },
        'xxx-large': { name: '28px', value: '7' },
    }
    editor.create();
    return editor
}

const returnSymbol = "\n";

function searchData(jsonData,outputObject) {
    let endTag = "";
    let tagStack = [];
    let top = 0;
    let bReturn = false;
    switch (typeof jsonData){
        case "string":
            outputObject.result+=jsonData;
            break;
        case "object":
            switch (jsonData.tag){
                case "br":
                    tagStack[top++] = " ";
                    bReturn = true;
                    break;
                case "p":
                    tagStack[top++] = returnSymbol;
                    bReturn = true;
                    break;
                case "span":
                    for(let i in jsonData.attrs){
                        if(jsonData.attrs[i].name==="style"){
                            let tempRegArr;
                            tempRegArr = /font-size: ([0-9]*)px/.exec(jsonData.attrs[i].value);
                            if(tempRegArr!=null){
                                outputObject.result += "[SIZE="+tempRegArr[1]+"]";
                                tagStack[top++] = "[/SIZE]";
                            }
                        }
                    }
                    break;
                case "font":
                    for(let i in jsonData.attrs){
                        if(jsonData.attrs[i].name==="color"){
                            outputObject.result += "[color="+jsonData.attrs[i].value+"]";
                            tagStack[top++] = "[/color]";
                        }else if(jsonData.attrs[i].name==="size"){
                            outputObject.result += "[SIZE="+fontSizeMap[jsonData.attrs[i].value]+"]";
                            tagStack[top++] = "[/SIZE]";
                        }
                    }
                    break;
            }
            for(let i in jsonData.children){
                searchData(jsonData.children[i],outputObject);
            }
            while (top>0){
                endTag+= tagStack[--top];
            }
            if(bReturn){
                var regex = new RegExp(/(\[\/[a-zA-z]*\])+$/);
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

