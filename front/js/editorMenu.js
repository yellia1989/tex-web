
const E = window.wangEditor

const { $, BtnMenu, DropListMenu, PanelMenu, DropList, Panel, Tooltip } = E

class FonSizeMenu extends PanelMenu{

    constructor(editor) {
        // data-title属性表示当鼠标悬停在该按钮上时提示该按钮的功能简述
        const $elem = E.$(
            `<div class="w-e-menu" data-title="FontSize">
                <input id="fontSizeInput" style="width: 30px;text-align: center;" type="text" value="22"/>
            </div>`
        )
        super($elem, editor)
    }

    /**
     * 菜单点击事件
     */
    clickHandler() {
        const panel = new Panel(this,{width:100,height:100,tabs:[{
                title:"字体大小",
                tpl:'<div class="w-e-button-container"><div id="FontConfirm">确定</div></div>',
                events:[{
                    selector: '#FontConfirm',
                    type : 'click',
                    fn:()=>{
                        // console.log(editor.selection.getSelectionText());
                        // console.log(editor.selection.getSelectionContainerElem().elems[0]);
                        // console.log(editor.selection.getSelectionStartElem().elems[0]);
                        // console.log(editor.selection.getSelectionEndElem().elems[0]);
                        // console.log(editor.selection.isSelectionEmpty());
                        // var rang = editor.selection.getRange();
                        if(editor.selection.isSelectionEmpty()){
                            return true;
                        }
                        var rang = this.selectRange;
                        var selectedText = rang.extractContents();
                        var ele = document.createElement("span");
                        console.log($("#fontSizeInput").val());
                        ele.style = "font-size:"+$("#fontSizeInput").val().trim()+"px;";
                        ele.appendChild(selectedText);
                        rang.insertNode(ele);
                        console.log(this.selectRange);
                        //rang.surroundContents(ele);
                        return true;
                    },
                }],
            }]});
        panel.create();
        this.selectRange = editor.selection.getRange();
        // $("#FontConfirm").on("click",function(){
        //   var selectedText = rang.extractContents();
        //   var ele = document.createElement("span");
        //   var input = $("#fontSizeInput");
        //   var size = $("#fontSizeInput").val();
        //   console.log(size);
        //   ele.style = "font-size:30px;";
        //   ele.appendChild(selectedText);
        //   rang.insertNode(ele);
        // });
    }

    tryChangeActive() {

    }

}

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
        console.log("cCLICK");
        // const panel = new Panel(this,{width:150,height:100,tabs:[{
        //     title:"颜色",
        //     tpl:'<div >选择颜色<div id="colorPicker" style="margin-left: 10px"></div></div>',
        //     events:[{
        //       selector: '#TestColor',
        //       type : 'click',
        //       fn:()=>{
        //         console.log(editor.selection.getSelectionText());
        //         console.log(editor.selection.getSelectionContainerElem().elems[0]);
        //         console.log(editor.selection.getSelectionStartElem().elems[0]);
        //         console.log(editor.selection.getSelectionEndElem().elems[0]);
        //         console.log(editor.selection.isSelectionEmpty());
        //         editor.cmd.do('foreColor', '#FFFFFF');
        //         editor.cmd.do('cut');
        //         return true;
        //       },
        //     }],
        //   }]});
        // panel.create();

    }

    tryChangeActive() {

    }

}


// 注册菜单
E.registerMenu('myColorKey', MyColorMenu);
E.registerMenu('FonSizeKey',FonSizeMenu);


function searchData(jsonData,outputObject) {
    let endTag = "";
    switch (typeof jsonData){
        case "string":
            outputObject.result+=jsonData;
            break;
        case "object":
            switch (jsonData.tag){
                case "br":
                    outputObject.result+="\n ";
                    break;
                case "p":
                    endTag += "\n";
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
                            endTag = "[/color]";
                        }
                    }
                    break;
            }
            for(let i in jsonData.children){
                searchData(jsonData.children[i],outputObject);
            }
            outputObject.result+=endTag;
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

