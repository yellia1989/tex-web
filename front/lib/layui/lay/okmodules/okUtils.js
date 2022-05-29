"use strict";
layui.define(["layer"], function (exprots) {
    var $ = layui.jquery;
    var okUtils = {
        stringToJson: function(text) {
          if ($.trim(text) == '') {
            return {};
          }
          return JSON.parse(text);
        },
        jsonToString: function(json) {
            return JSON.stringify(json, null, '\t');
        },
        getJsonValue: function(json, fieldKey) {
            if (json === undefined) {
                return undefined;
            }
            if (typeof (fieldKey) == 'number') {
                return json[fieldKey];
            }
            let idx = fieldKey.indexOf('.');
            if (idx === -1) {
                return json[fieldKey];
            }
            return this.getJsonValue(json[fieldKey.substr(0, idx)], fieldKey.substr(idx + 1));
        },
        setJsonValue: function(json, fieldKey, fieldValue) {
            if (typeof(json) != 'object') {
                return;
            }
            var idx = fieldKey.indexOf('.');
            if (idx === -1) {
                json[fieldKey] = fieldValue;
                return;
            }

            var k1 = fieldKey.substr(0, idx), k2 = fieldKey.substr(idx + 1);
            if (json[k1] == undefined) {
                json[k1] = {};
            }
            this.setJsonValue(json[k1], k2, fieldValue);
        },
        isEmpty: function(obj) {
            for (var k in obj) {
                if (obj[k] !== undefined) {
                    return false;
                }
            }
            return true;
        },
        /**
         * 获取body的总宽度
         */
        getBodyWidth: function () {
            return document.body.scrollWidth;
        },
        /**
         * 主要用于对ECharts视图自动适应宽度
         */
        echartsResize: function (element) {
            var element = element || [];
            window.addEventListener("resize", function () {
                var isResize = localStorage.getItem("isResize");
                // if (isResize == "false") {
                for (let i = 0; i < element.length; i++) {
                    element[i].resize();
                }
                // }
            });
        },
        /**
         * ajax()函数二次封装
         * @param url
         * @param type
         * @param params
         * @param load
         * @returns {*|never|{always, promise, state, then}}
         */
        ajax: function (url, type, params, load) {
            var deferred = $.Deferred();
            var loadIndex;
            $.ajax({
                url: url,
                type: type || "get",
                data: params || {},
                dataType: "json",
                timeout: 3000,
                beforeSend: function () {
                    if (load) {
                        loadIndex = layer.load(0, {shade: 0.3});
                    }
                },
                success: function (data) {
                    if (data.code == 0) {
                        // 业务正常
                        deferred.resolve(data)
                    } else if (data.code == 9999) {
                        // 需要登录
                        layer.msg(data.msg, {icon: 7, time: 2000}, function() {
                            parent.window.location.href = "/login.html";
                        });
                    } else {
                        // 业务异常
                        layer.msg(data.msg, {icon: 7, time: 2000});
                        deferred.reject(data);
                    }
                },
                complete: function () {
                    if (load) {
                        layer.close(loadIndex);
                    }
                },
                error: function () {
                    layer.close(loadIndex);
                    layer.msg("服务器错误", {icon: 2, time: 2000});
                    deferred.reject("okUtils.ajax error: 服务器错误");
                }
            });
            return deferred.promise();
        },
        /**
         * 主要用于针对表格批量操作操作之前的检查
         * @param table
         * @returns {string}
         */
        tableBatchCheck: function (table) {
            var checkStatus = table.checkStatus("tableId");
            var rows = checkStatus.data.length;
            if (rows > 0) {
                var idsStr = "";
                for (var i = 0; i < checkStatus.data.length; i++) {
                    idsStr += checkStatus.data[i].id;
                    if (i != checkStatus.data.length-1) {
                      idsStr += ",";
                    }
                }
                return idsStr;
            } else {
                layer.msg("未选择有效数据", {offset: "t", anim: 6});
            }
        },
        /**
         * 在表格页面操作成功后弹窗提示
         * @param content
         */
        tableSuccessMsg: function (content) {
            layer.msg(content, {icon: 1, time: 1000}, function () {
                // 刷新当前页table数据
                $(".layui-laypage-btn")[0].click();
            });
        },
        /**
         * 获取父窗体的okTab
         * @returns {string}
         */
        getOkTab: function () {
            return parent.objOkTab;
        },
        /**
         * 格式化当前日期
         * @param date
         * @param fmt
         * @returns {void | string}
         */
        dateFormat: function (date, fmt) {
            var o = {
                "M+": date.getMonth() + 1,
                "d+": date.getDate(),
                "h+": date.getHours(),
                "m+": date.getMinutes(),
                "s+": date.getSeconds(),
                "q+": Math.floor((date.getMonth() + 3) / 3),
                "S": date.getMilliseconds()
            };
            if (/(y+)/.test(fmt))
                fmt = fmt.replace(RegExp.$1, (date.getFullYear() + "").substr(4 - RegExp.$1.length));
            for (var k in o)
                if (new RegExp("(" + k + ")").test(fmt))
                    fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
            return fmt;
        },
        number: {
            /**
             * 判断是否为一个正常的数字
             * @param num
             */
            isNumber: function (num) {
                if (num && !isNaN(num)) {
                    return true;
                }
                return false;
            },
            /**
             * 判断一个数字是否包括在某个范围
             * @param num
             * @param begin
             * @param end
             */
            isNumberWith: function (num, begin, end) {
                if (this.isNumber(num)) {
                    if (num >= begin && num <= end) {
                        return true;
                    }
                    return false;
                }
            },
            /**
             * 格式化百分比，默认带2位小数点
             * @param num1 分子
             * @param num2 分母
             * @param width 小数点位数,默认2位小数点
             */
            percent: function(num1, num2, width, suffix) {
                if (width == undefined) {
                    width = 2
                }
                if (suffix == undefined) {
                    suffix = "%"
                }
 
                if (num2 == 0) return "0."+"0".repeat(width) + suffix;
                return (num1/num2).toFixed(width) + suffix;
            }
        },
        displayInterval: function(interval) {
            if (interval < 0) return "-";
            let day = Math.floor(interval / 86400);
            interval -= day*86400;

            let hour = Math.floor(interval / 3600);
            interval -= hour*3600;

            let min = Math.floor(interval / 60);
            interval -= min*60;

            let sec = Math.floor(interval);

            let result = '';
            if (day > 0) result += day+'d';
            if (hour > 0) result += hour+'h';
            if (min > 0) result += min+'m';
            if (sec > 0) result += sec+'s';
            return result;
        }
    };
    exprots("okUtils", okUtils);
});
