(function (factory) {
    var activity = window.activity || (window.activity = {});
    var scope = activity.config || (activity.config = {});
    factory(scope);
})(function (scope) {
    /*
     // field
     fieldKey: { // fieldKey如果是带层级的，要用引号引起来
         name: '设置项', // 中文描述
         type: 'text', // 类型, text|midtext|longtext|textarea|select|datetime|number|map|array
         min: 0, // number下限
         max: 0, // number上线
         options: { val: 'text' }, // select选项，value=>说明
         printer: fn, // fn(object), 把json对象转换为配置文本项
         parser: fn, // 解析函数fn(text)，把文本解析成json对象，比如JSON.parse，parseInt
         checkValid: fn, // 校验配置是否合法函数fn(object)
         onChange: fn, // 字段被设置为指定值, fn(object)
         separate: true, // 是否在这个字段前面加上换行
         subFieldOption: {} // 子字段配置，可以是一个object，或者一个返回object的fn，参数为fieldKey，

         useWholeValue: false,  // type=map|array时有效，是否使用整个字段，如果为false则根据fieldKey取值，否则取整个字段本身
         vertical: false, // type=map|array时有效，是否竖向显示编辑框
         groupFieldOption: {  // type=map|array时有效，分组设置。里面不能再次嵌套map|array
            groupFieldKey: {
                 // 字段描述，同上，比如name, type, parser, ...
                useWholeValue: false, // 是否使用整个字段
                isMapKey: true, // type=map时必须设置一个，且只能一个，与map的key关联
                arrayIndex: 0, // type=array时必须设置，和数组的位置关联。
                                // useWholeValue, isMapKey, arrayIndex在groupFieldOption中只能包含一种
                isOptional: false, // 是否可选，会增加一个checkbox按钮
            }
         }
     }
     */

    // 通用选项
    var options = {
        optYesNo: {
            0: '否',
            1: '是'
        },
        ActType2:{
            1: '道具消耗',
            2: '英雄战力提升',
            3: '赏金招募次数',
            4: '王者竞技场胜利次数',
            5: '天赋升级次数',
            6: '英雄试练胜利次数',
            7: '巅峰竞技场胜利次数',
            8: '关卡进度',
            9: '商店购买次数',
            10: '王者竞技场获得积分',
            11: '宠物积分',
        },
        ActType3:{
            1: '道具消耗',
            2: '英雄战力提升',
            3: '赏金招募次数',
            4: '王者竞技场胜利次数',
            5: '天赋升级次数',
            6: '英雄试练胜利次数',
            7: '巅峰竞技场胜利次数',
            8: '关卡进度',
            12: '龙域探险加速次数',
            13: '天赋点购买次数',
            14: '宝石购买次数',
            15: '符文碎片购买次数',
        },
        ActType4:{
            1: '王者竞技场',
            3: '巅峰竞技场',
            4: '金币消耗',
            5: '战力',
            6: '钻石消耗',
            8: '关卡进度',
        },
        ActTypeAllianceTop100:{
            1: '体力消耗',
            2: '木材获得',
            3: '石料获得',
            4: '银币获得',
            5: '击杀野怪',
            6: '击杀Boss',
            7: '获得能力之石',
            8: '获得旗帜宝箱',
        },
        ActTypeAllianceConsume:{
            1: '购买体力次数',
            2: '木材消耗',
            3: '石料消耗',
            4: '银币消耗',
            5: '击杀野怪',
            6: '击杀Boss',
            7: '使用能力之石',
            8: '攻城成功',
            9: '埋地雷',
        },
    };

    // 工具函数
    var utils = {
        printItemNumOddsList: function(array) {
            if (!array) return undefined;

            var result = '';
            for (var i = 0; i < array.length; ++i) {
                if (i != 0) result += ';';
                result += array[i][0] + ',' + array[i][1] + ',' + array[i][2];
            }
            return result;
        },
        parseItemNumOddsList: function (text) {
            var result = [];
            var arr1, arr2, i, j, idnum, id, num, odds;
            arr1 = text.split(';');
            for (i = 0; i < arr1.length; ++i) {
                idnum = $.trim(arr1[i]);
                if (idnum == '') continue;

                arr2 = idnum.split(',');
                if (arr2.length != 3 || !$.isNumeric(arr2[0]) || !$.isNumeric(arr2[1]) || !$.isNumeric(arr2[2])) {
                    throw '格式错误(id,num,odds;id,num,odds)';
                }

                id = parseInt(arr2[0]);
                num = parseInt(arr2[1]);
                               odds = parseInt(arr2[2]);
                if (id < 0 || num < 0 || odds < 0) {
                    throw '三个数字均为正整数例如1,100,50;2,200,30;';
                }

                result.push([id, num, odds]);
            }
            return result;
        },
        printItemNumList: function(array) {
            if (!array) return undefined;

            var result = '';
            for (var i = 0; i < array.length; ++i) {
                if (i != 0) result += ';';
                result += array[i][0] + ',' + array[i][1];
            }
            return result;
        },
        parseItemNumList: function (text) {
            var result = [];
            var arr1, arr2, i, j, idnum, id, num;
            arr1 = text.split(';');
            for (i = 0; i < arr1.length; ++i) {
                idnum = $.trim(arr1[i]);
                if (idnum == '') continue;

                arr2 = idnum.split(',');
                if (arr2.length != 2 || !$.isNumeric(arr2[0]) || !$.isNumeric(arr2[1])) {
                    throw '格式错误(id,num;id,num)';
                }

                id = parseInt(arr2[0]);
                num = parseInt(arr2[1]);
                if (id < 0 || num < 0) {
                    throw '道具ID或数量错误';
                }

                result.push([id, num]);
            }
            return result;
        },
     	printIdList: function(array) {
			if (!array) return '';
			
			var result = '';
			for (var i = 0; i < array.length; ++i) {
				if (i != 0) result += ';';
				result += array[i];
			}
			return result;
		},
		parseIdList: function (text) {
			var result = [];
			var arr, i, iId, sId;
			arr = text.split(';');
			for (i = 0; i < arr.length; ++i) {
				sId = $.trim(arr[i]);
				if (sId == '') continue;
				
				if (!$.isNumeric(sId)) {
					throw '格式错误(id1;id2;id3)';
				}
				iId = parseInt(sId);
				
				result.push(iId);
			}
			return result;
		},
		printSuffixList: function(array) {
            if (!array) return '';

            var result = '';
            for (var i = 0; i < array.length; ++i) {
                if (i != 0) result += ';';
                result += array[i];
            }
            return result;
        },
        parseSuffixList: function (text) {
            var result = [];
            var arr, i, iId, sId;
            arr = text.split(';');
            for (i = 0; i < arr.length; ++i) {
                sId = $.trim(arr[i]);
                if (sId == '') continue;

                if (!$.isNumeric(sId)) {
                    throw '格式错误(id1;id2;id3)';
                }
				iId = parseInt(sId);
				if (iId > 9)
				{
					throw '数值错误(id取值[0,1,2,3,4,5,6,7,8,9])';
				}

                result.push(iId);
            }
            return result;
        },
        changeActivityType: function (type) {
            currentActivityType = type;
        },
        getFieldOptionByActivity: function (fieldKey) {
            var activityType = activityTypeDefine[currentActivityType];
            if (activityType && activityType.fieldOption) {
                return activityType.fieldOption[fieldKey];
            }
            return undefined;
        },
        stringToDate: function(dateStr) {
            var converted = Date.parse(dateStr);
            var myDate = new Date(converted);
            if (isNaN(myDate))
            {
                var arys= dateStr.split('-');
                myDate = new Date(arys[0],--arys[1],arys[2]);
            }
            return myDate;
        },
        formatDate: function(fmt,d) {
            var o = {
                "M+": d.getMonth() + 1, //月份 
                "d+": d.getDate(), //日 
                "h+": d.getHours(), //小时 
                "m+": d.getMinutes(), //分 
                "s+": d.getSeconds(), //秒 
                "q+": Math.floor((d.getMonth() + 3) / 3), //季度 
                "S": d.getMilliseconds() //毫秒 
            };
            if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (d.getFullYear() + "").substr(4 - RegExp.$1.length));
            for (var k in o)
                if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
            return fmt;
        },
        addHour: function(dateStr, hour) {
            var dt = this.stringToDate(dateStr);
            var ndt = new Date(dt.getTime() + hour*3600*1000);
            return this.formatDate('yyyy-MM-dd hh:mm:ss', ndt);
        }
    };

    // 活动类型定义
    var activityTypeDefine = {};
    // 1
    activityTypeDefine[1] = {
        name: '模块控制',
        fieldOption: {
            comm_param: {
                close_cmd: {
                    name: '关闭协议(,分隔)',
                    type: 'longtext'
                },
            }
		}
    };
    // 2
    activityTypeDefine[2] = {
        name: '个人消耗',
        fieldOption: {
            comm_param: {
                type: {
                    name: '子类型',
                    type: 'select',
                    options: options.ActType3,
                    parser: parseInt,
                },
                itemid: {
                    name: '道具id',
                    type: 'text',
                },
                shopid: {
                    name: '商店id',
                    type: 'text',
                },
                rank: {
                     name: '完成奖励',
                     type: 'map',
                     groupFieldOption: {
                         _: {
                            name: '完成条件',
                            type: 'text',
                            isMapKey: true,
                            parser: parseInt
                         },
                         reward: {
                            name: '奖励:id,num',
                            type: 'longtext',
                            printer: utils.printItemNumList,
                            parser: utils.parseItemNumList
                         }
                     }
                 }
            },
            client_param: {
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                },
                show_step: {
                    name: '显示档位',
                    type: "text"
                },
            }
        }
    };
    // 3
    activityTypeDefine[3] = {
        name: '冲榜',
        fieldOption: {
            comm_param: {
                type: {
                    name: '子类型',
                    type: 'select',
                    options: options.ActType2,
                    parser: parseInt,
                },
                itemid: {
                    name: '道具id',
                    type: 'text',
                },
                shopid: {
                    name: '商店id',
                    type: 'text',
                },
                rank: {
                     name: '排名奖励',
                     type: 'map',
                     groupFieldOption: {
                         _: {
                             name: '排名:from_to',
                             type: 'text',
                             isMapKey: true
                         },
                         reward: {
                             name: '奖励:id,num',
                             type: 'longtext',
                             printer: utils.printItemNumList,
                             parser: utils.parseItemNumList
                         }
                     }
                 }
            },
            client_param: {
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                },
                help_id : {
                    name: '帮助提示',
                    type: 'text',
                }

            }
        }
    };
    // 4
    activityTypeDefine[4] = {
        name: '累计充值',
        fieldOption: {
            comm_param: {
                everyday: {
                    name: '每日充值奖励',
                    type: 'map',
                    groupFieldOption: {
                         _: {
                            name: '完成条件',
                            type: 'text',
                            isMapKey: true,
                            parser: parseInt
                         },
                         reward: {
                           name: '奖励:id,num',
                           type: 'longtext',
                           printer: utils.printItemNumList,
                           parser: utils.parseItemNumList
                        }
                    }
                },
                totalMoney: {
                    name: '累计充值奖励',
                    type: 'map',
                    groupFieldOption: {
                         _: {
                            name: '完成条件',
                            type: 'text',
                            isMapKey: true,
                            parser: parseInt
                         },
                         reward: {
                           name: '奖励:id,num',
                           type: 'longtext',
                           printer: utils.printItemNumList,
                           parser: utils.parseItemNumList
                        }
                    }
                },
                totalDays: {
                    name: '累天充值奖励',
                    type: 'map',
                    groupFieldOption: {
                         _: {
                            name: '完成条件',
                            type: 'text',
                            isMapKey: true,
                            parser: parseInt
                         },
                         reward: {
                           name: '奖励:id,num',
                           type: 'longtext',
                           printer: utils.printItemNumList,
                           parser: utils.parseItemNumList
                        }
                    }
                },
                independent_step: {
                    name: '需要额外展示的挡位(挡位类型, 挡位)',
                    type: 'text',
                }
            },
            client_param: {
                every_day_pos: {
                    name: '每日充值金额隐藏档位',
                    type: 'text'
                },
                total_money_pos: {
                    name: '累计充值金额隐藏档位',
                    type: 'text'
                },
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                },
                every_day_pos: {
                    name: '每日充值档位',
                    type: 'text',
                    parser: parseInt,
                },
                total_money_pos: {
                    name: '累计充值档位',
                    type: 'text',
                    parser: parseInt,
                },
                hero_id: {
                    name: '额外展示档位展示的英雄id:单位表id',
                    type: 'text',
                }
            }
        }
    };
    // 5
    activityTypeDefine[5] = {
        name: '首冲送英雄',
        fieldOption: {
            comm_param: {
                item: {
                    name: '道具:id,num',
                    type: 'midtext',
                    printer: utils.printItemNumList,
                    parser: utils.parseItemNumList
                }
            },
            client_param: {
                show_count_down: {
                    name: '是否显示倒计时',
                    type: 'select',
                    options: {
                        '0' : '不显示',
                        '1' : '显示',
                    }
                },
                background_img: {
                    name: '背景',
                    type: 'longtext',
                },
                itemid: {
                    name: '道具id(英雄)',
                    type: 'text',
                },
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            }
        }
    };
    // 6
    activityTypeDefine[6] = {
        name: '一元购',
        fieldOption: {
            comm_param: {
                reward: {
                    name: '奖励:id,num',
                    type: 'longtext',
                    printer: utils.printItemNumList,
                    parser: utils.parseItemNumList
                },
                productid: {
                    name: '商品id',
                    type: 'text',
                    parser: parseInt
                }
            },
            client_param: {
                price: {
                    name: '原价',
                    type: 'text',
                },
                background_img: {
                    name: '背景',
                    type: 'longtext',
                },
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            }
        }
    };
    // 7
    activityTypeDefine[7] = {
        name: '七日登录',
        fieldOption: {
            comm_param: {
                days: {
                    name: '天数',
                    type: 'map',
                    groupFieldOption: {
                         _: {
                            name: '第几天',
                            type: 'text',
                            isMapKey: true,
                            parser: parseInt
                         },
                         reward: {
                           name: '奖励:活跃度;id,num',
                           type: 'longtext',
                        }
                    }
                }
            },
            client_param: {
                background_img: {
                    name: '背景',
                    type: 'longtext',
                },
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            }
        }
    };
    // 8
    activityTypeDefine[8] = {
        name: '免费福利',
        fieldOption: {
            comm_param: {
                reward: {
                   name: '奖励:id,num',
                   type: 'longtext',
                   printer: utils.printItemNumList,
                   parser: utils.parseItemNumList
                },
                condition: {
                    name: '达成条件',
                    type: 'select',
                    options: {
                        '1': '购买月卡',
                        '2': '通关关卡'
                    },
                    parser: parseInt
                }
            },
            server_param: {
                cond_param: {
                    name: '条件参数',
                    type: 'text',
                    parser: parseInt
                }
            },
            client_param: {
                one_time: {
                    name: '是否显示one_time offer',
                    type: 'select',
                    options: {
                        '0' : '不显示',
                        '1' : '显示',
                    }
                },
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            }
        }
    };
    // 9
    activityTypeDefine[9] = {
        name: '成长基金',
        fieldOption: {
            comm_param: {
                step: {
                    name: '奖励',
                    type: 'map',
                    groupFieldOption: {
                         _: {
                            name: '档位',
                            type: 'text',
                            isMapKey: true,
                            parser: parseInt
                         },
                         reward: {
                           name: '奖励:id,num',
                           type: 'longtext',
                           printer: utils.printItemNumList,
                           parser: utils.parseItemNumList
                         }
                    }
                },
                productid: {
                    name: '商品id',
                    type: 'text',
                    parser: parseInt
                },
                condition: {
                    name: '基金类型',
                    type: 'select',
                    options: {
                        '1': '登录',
                        '2': '通关关卡'
                    },
                    parser: parseInt
                }
            },
            client_param: {
                price: {
                    name: '原价',
                    type: 'text',
                    parser: parseInt
                },
                background_img: {
                    name: '背景',
                    type: 'longtext',
                },
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            }
        }
    };
    // 10
    activityTypeDefine[10] = {
        name: '个性化活动',
        fieldOption: {
            comm_param: {
                reward: {
                    name: '奖励:id,num',
                    type: 'longtext',
                    printer: utils.printItemNumList,
                    parser: utils.parseItemNumList
                },
                buy_times: {
                    name: '最大购买次数',
                    type: 'text',
                    parser: parseInt
                }
            },
            client_param: {
                discount: {
                    name: '折扣',
                    type: 'text',
                    parser: parseInt
                },
                icon: {
                    name: '角标',
                    type: 'text',
                },
                detail_bg: {
                    name: '活动背景',
                    type: 'text',
                },
                list_bg: {
                    name: '列表背景',
                    type: 'text',
                },
                title_icon: {
                    name: '标题图片',
                    type: 'text',
                },
                desc_detail_icon: {
                    name: '描述图片',
                    type: 'text',
                },
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                }
            },
            server_param: {
                trigger_point: {
                    name: '触发点',
                    type: 'select',
                    options: {
                        '1': '王者竞技场失败',
                        '2': '巅峰竞技场失败',
                        '3': '天赋等级提升',
                        '4': '钻石免费抽',
                        '5': '登录触发',
                        '6': '技能升级',
                        '7': '英雄升级',
                        '8': '购买符文',
                        '9': '英雄历练升级',
                        '10': '主线任务进度变化',
                        '11': '商店购买',
                        '12': '蓝钻商城购买',
                        '13': '天赋技能解锁',
                        '14': '钻石变更时',
                    },
                    parser: parseInt
                },
                trigger_interval: {
                    name: '触发间隔(秒)',
                    type: 'text',
                    parser: parseInt
                },
                trigger_times: {
                    name: '触发次数',
                    type: 'text',
                    parser: parseInt
                },
                last_time: {
                    name: '持续时间(秒)',
                    type: 'text',
                    parser: parseInt
                },
                productid: {
                    name: '商品id列表(,区分)',
                    type: 'text',
                },
                buy_month_card: {
                    name: '是否买过月卡',
                    type: 'select',
                    options: options.optYesNo,
                    parser: parseInt
                },
                no_heroids: {
                    name: '没有英雄(;区分)',
                    type: 'midtext',
                },
                talent_level_range: {
                    name: '天赋等级范围(-区分)',
                    type: 'midtext',
                },
                talent_range: {
                    name: '天赋点范围(-区分)',
                    type: 'midtext',
                },
                coin_range: {
                    name: '金币范围(-区分)',
                    type: 'midtext',
                },
                diamond_range: {
                    name: '钻石范围(-区分)',
                    type: 'midtext',
                },
                exp_range: {
                    name: '经验药水数量,单位K(-区分)',
                    type: 'midtext'
                },
                gem_range: {
                    name: '宝石范围(-区分)',
                    type: 'midtext',
                },
                temper_level_range: {
                    name: '历练等级范围(-区分)',
                    type: 'midtext',
                },
                rune_range: {
                    name: '符文碎片范围(-区分)',
                    type: 'midtext',
                },
                recharge_money_range: {
                    name: '充值美金范围(-区分)',
                    type: 'midtext',
                },
                recharge_last_time: {
                    name: 'N天前充值',
                    type: 'midtext',
                    parser: parseInt,
                },
                vippoint_range: {
                    name: 'vip点范围(-区分)',
                    type: 'midtext',
                },
                recharge_times_range: {
                    name: '付费次数范围(-区分)',
                    type: 'midtext',
                },
                hero_star_min: {
                    name: '所有英雄星级小于',
                    type: 'midtext',
                    parser: parseInt,
                },
                has_free_hero: {
                    name: '有剩余英雄未组队',
                    type: 'select',
                    options: options.optYesNo,
                    parser: parseInt
                },
                hero_skilllv_range: {
                    name: '英雄技能等级范围(-区分)',
                    type: 'midtext'
                },
                hero_level_range: {
                    name: '英雄等级范围(-区分)',
                    type: 'midtext'
                },
                hero_runelv_range: {
                    name: '英雄符文等级范围(-区分)',
                    type: 'midtext'
                },
                month_card_left: {
                    name: '月卡剩余天数',
                    type: 'midtext',
                    parser: parseInt,
                },
                month_card_expire: {
                    name: '月卡过期天数',
                    type: 'midtext',
                    parser: parseInt,
                },
                stage_range: {
                    name: '挂机关卡范围(-区分)',
                    type: 'midtext',
                },
                main_quest_range: {
                    name: '主线任务进度范围(-区分)',
                    type: 'midtext'
                },
                once_recharge_money: {
                    name: '单次充值超过数额',
                    type: 'midtext',
                    parser: parseInt,
                },
                coin_rate: {
                    name: '当前金币量小于(点金卷所得金币的倍率)',
                    type: 'midtext',
                    parser: parseInt,
                },
                exp_rate: {
                    name: '当前经验药水量小于(药水礼盒所得药水的倍率)',
                    type: 'midtext',
                    parser: parseInt,
                },
                coin_less_skillcost: {
                    name: '当前金币量小于上次技能升级所消耗的金币量',
                    type: 'select',
                    options: options.optYesNo,
                    parser: parseInt
                },
                exp_less_herolevelcost: {
                    name: '当前药水量小于上次英雄升级所消耗的药水量',
                    type: 'select',
                    options: options.optYesNo,
                    parser: parseInt
                },
                rune_level: {
                    name: '购买符文的等级(-区分)',
                    type: 'midtext'
                },
                growth_gift_step: {
                    name: '成长礼包活动阶段',
                    type: 'select',
                    options: {
                        '1': '野蛮成长礼包',
                        '2': '急速成长礼包',
                        '3': '快速成长礼包',
                        '4': '加速成长礼包',
                    },
                    parser: parseInt
                }
            }
        }
    };
    // 11
    activityTypeDefine[11] = {
        name: '在线时长奖励',
        fieldOption: {
            comm_param: {
                days: {
                     name: '第几天',
                     type: 'map',
                     vertical: true,
                     groupFieldOption: {
                         _: {
                             name: '天数',
                             type: 'text',
                             isMapKey: true
                         },
                         step: {
                             name: '时长奖励:时长:id,num;id,num|时长:id,num;id,num',
                             type: 'textarea',
                         }
                     }
                 }
            },
            client_param: {
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            }
        }
    };
    // 12
    activityTypeDefine[12] = {
        name: '商品团购',
        fieldOption: {
            comm_param: {
                heroes: {
                     name: '英雄id',
                     type: 'map',
                     vertical: true,
                     groupFieldOption: {
                         _: {
                             name: '英雄整卡id',
                             type: 'text',
                             isMapKey: true
                         },
                         step: {
                             name: '个数:人数:商品id:vip:diamond:price:id,num;id,num',
                             type: 'textarea',
                         },
                     }
                 }
            },
            client_param: {
                heroes: {
                     name: '英雄id',
                     type: 'map',
                     vertical: true,
                     groupFieldOption: {
                         _: {
                             name: '英雄整卡id',
                             type: 'text',
                             isMapKey: true
                         },
                         icon: {
                             name: '显示图片路径',
                             type: 'text',
                         },
                     }
                 },
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            },
        }
    };
    // 13
    activityTypeDefine[13] = {
        name: '许愿',
        fieldOption: {
            comm_param: {
                type: {
                    name: '子类型',
                    type: 'select',
                    options: options.ActType4,
                    parser: parseInt,
                },
                rank: {
                     name: '档位',
                     type: 'map',
                     vertical: true,
                     groupFieldOption: {
                         _: {
                             name: '排名',
                             type: 'text',
                             isMapKey: true
                         },
                         step: {
                             name: '档位名字:商品id:id,num;id,num',
                             type: 'textarea',
                         },
                     }
                 }
            },
            client_param: {
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            }
        }
    };
    // 14
    activityTypeDefine[14] = {
         name: '英雄兑换',
         fieldOption: {
             comm_param: {
                 heros: {
                     name: '英雄',
                     type: 'map',
                     groupFieldOption: {
                         _: {
                             name: '英雄整卡id',
                             type: 'text',
                             isMapKey: true,
                             parser: parseInt
                         },
                         items: {
                             name: '道具:id,num',
                             type: 'text',
                             printer: utils.printItemNumList,
                             parser: utils.parseItemNumList
                         }
                     }
                 }
            },
            client_param: {
                heroes: {
                     name: '英雄id',
                     type: 'map',
                     vertical: true,
                     groupFieldOption: {
                         _: {
                             name: '英雄id',
                             type: 'text',
                             isMapKey: true
                         },
                         icon: {
                             name: '显示图片路径',
                             type: 'text',
                         },
                     }
                 },
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            },
         }
    };
    // 15
    activityTypeDefine[15] = {
         name: '幸运转盘',
         fieldOption: {
             comm_param: {
                 play_diamond: {
                     name: '转动一次需要钻石',
                     type: 'text',
                     parser: parseInt,
                 },
                 wheels: {
                     name: '格子',
                     type: 'map',
                     groupFieldOption: {
                         _: {
                             name: '格子id',
                             type: 'text',
                             isMapKey: true,
                             parser: parseInt
                         },
                         reward: {
                             name: '道具，格式:id,num;权重(万分比);展示类型(1上,2下);走马灯显示(0不显示,1显示)',
                             type: 'longtext',
                         }
                     }
                 },
                 times: {
                     name: '每日转动奖励',
                     type: 'map',
                     groupFieldOption: {
                         _: {
                             name: '次数',
                             type: 'text',
                             isMapKey: true,
                             parser: parseInt
                         },
                         reward: {
                             name: '道具:id,num',
                             type: 'midtext',
                             printer: utils.printItemNumList,
                             parser: utils.parseItemNumList
                         }
                     }
                 },
                 rank: {
                     name: '排行奖励',
                     type: 'map',
                     groupFieldOption: {
                         _: {
                             name: '排名',
                             type: 'text',
                             isMapKey: true,
                             parser: parseInt
                         },
                         reward: {
                             name: '道具:id,num',
                             type: 'midtext',
                             printer: utils.printItemNumList,
                             parser: utils.parseItemNumList
                         }
                     }
                 }
            },
            client_param: {
                hot_hero_bg: {
                    name: '英雄立绘',
                    type: 'longtext',
                },
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                },
                hot_hero: {
                    name: '热点英雄',
                    type: 'text'
                }
            }
        }
    };
    // 16
    activityTypeDefine[16] = {
         name: '邀请有礼',
         fieldOption: {
             comm_param: {
                 bind_diamond: {
                     name: '绑定成功获得的钻石',
                     type: 'text',
                     parser: parseInt,
                 },
                 bind_reward: {
                     name: '绑定人数奖励',
                     type: 'map',
                     groupFieldOption: {
                         _: {
                             name: '人数',
                             type: 'text',
                             isMapKey: true,
                             parser: parseInt
                         },
                         reward: {
                             name: '奖励，格式 道具id,道具数量',
                             type: 'text',
                         }
                     }
                 },
                 bind_main_quest: {
                     name: '绑定主线任务',
                     type: 'map',
                     groupFieldOption: {
                         _: {
                             name: '主线任务id',
                             type: 'text',
                             isMapKey: true,
                             parser: parseInt
                         },
                         reward: {
                             name: '奖励，格式 人数:道具id,道具数量;道具id,道具数量|人数:道具id,道具数量;道具id,道具数量',
                             type: 'longtext',
                         }
                     }
                 },
                 bind_vip: {
                     name: '绑定vip等级',
                     type: 'map',
                     groupFieldOption: {
                         _: {
                             name: 'vip等级',
                             type: 'text',
                             isMapKey: true,
                             parser: parseInt
                         },
                         reward: {
                             name: '奖励，格式 人数:道具id,道具数量;道具id,道具数量|人数:道具id,道具数量;道具id,道具数量',
                             type: 'longtext',
                         }
                     }
                 },
             },
            client_param: {
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            }
        }
    };
    // 17
    activityTypeDefine[17] = {
         name: 'vip福利',
         fieldOption: {
             comm_param: {
                 vip_reward: {
                     name: 'vip等级',
                     type: 'map',
                     groupFieldOption: {
                         _: {
                             name: 'vip等级',
                             type: 'text',
                             isMapKey: true,
                             parser: parseInt
                         },
                         reward: {
                             name: '奖励，格式 道具id,道具数量',
                             type: 'text',
                         }
                     }
                 }
             },
            client_param: {
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            }
        }
    };
    // 18
    activityTypeDefine[18] = {
        name: '个性化赠礼',
        fieldOption: {
            comm_param: {
                reward: {
                    name: '奖励:id,num;id,num',
                    type: 'longtext',
                    printer: utils.printItemNumList,
                    parser: utils.parseItemNumList
                }
            },
            client_param: {
                detail_bg: {
                    name: '活动背景',
                    type: 'text'
                },
                list_bg: {
                    name: '列表背景',
                    type: 'text'
                },
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                }
            },
            server_param: {
                trigger_point: {
                    name: '触发点',
                    type: 'select',
                    options: {
                        '1': '穿戴装备',
                        '2': '打开英雄列表',
                        '3': '打开技能界面和点击技能升级'
                    },
                    parser: parseInt
                },
                trigger_interval: {
                    name: '触发间隔(秒)',
                    type: 'text',
                    parser: parseInt
                },
                trigger_times: {
                    name: '触发次数',
                    type: 'text',
                    parser: parseInt
                },
                last_time: {
                    name: '持续时间',
                    type: 'text',
                    parser: parseInt
                },
                stage_range: {
                    name: '挂机关卡范围(-区分)',
                    type: 'midtext'
                },
                hero_equip: {
                    name: '缺少装备',
                    type: 'select',
                    options: options.optYesNo,
                    parser: parseInt
                },
                hero_quality: {
                    name: 'X时间没有合成英雄(秒)',
                    type: 'text',
                    parser: parseInt
                },
                hero_skill: {
                    name: 'x个满技能等级英雄且升级技能缺少金币',
                    type: 'text',
                    parser: parseInt
                }
            },
        }
    };
    // 19
    activityTypeDefine[19] = {
        name: '联盟冲榜',
        fieldOption: {
            comm_param: {
                type: {
                    name: '子类型',
                    type: 'select',
                    options: options.ActTypeAllianceTop100,
                    parser: parseInt,
                },
                rank: {
                    name: '排名奖励',
                    type: 'map',
                    vertical: true,
                    groupFieldOption: {
                        _: {
                            name: '排名:from_to',
                            type: 'text',
                            isMapKey: true
                        },
                        masterreward: {
                            name: '盟主奖励:id,num',
                            type: 'longtext',
                            printer: utils.printItemNumList,
                            parser: utils.parseItemNumList
                        },
                        memberreward: {
                            name: '成员奖励:id,num',
                            type: 'longtext',
                            printer: utils.printItemNumList,
                            parser: utils.parseItemNumList
                        }
                    }
                }
            },
            client_param: {
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                },
                help_id : {
                    name: '帮助提示',
                    type: 'text',
                }
            }
        }
    };
    // 20
    activityTypeDefine[20] = {
        name: '联盟消耗',
        fieldOption: {
            comm_param: {
                type: {
                    name: '子类型',
                    type: 'select',
                    options: options.ActTypeAllianceConsume,
                    parser: parseInt,
                },
                rank: {
                    name: '排名奖励',
                    type: 'map',
                    vertical: true,
                    groupFieldOption: {
                        _: {
                            name: '排名:from_to',
                            type: 'text',
                            isMapKey: true
                        },
                        reward: {
                            name: '成员奖励:id,num',
                            type: 'longtext',
                            printer: utils.printItemNumList,
                            parser: utils.parseItemNumList
                        }
                    }
                }
            },
            client_param: {
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            }
        }
    };
    // 21
    activityTypeDefine[21] = {
        name: '白嫖活动',
        fieldOption: {
            comm_param: {
                reward: {
                    name: '可选的奖励: id(;分隔)',
                    type: 'text',
                },
                quest: {
                    name: '任务',
                    type: 'map',
                    vertical: true,
                    groupFieldOption: {
                        _: {
                            name: '任务id',
                            isMapKey: true,
                            parser: parseInt
                        },
                        cond_type: {
                            name: '完成条件类型',
                            type: 'select',
                            options: {
                                '1': '1. 完成成就(加入联盟)',
                                '2': '2. 多人对决借用英雄',
                                '3': '3. 在篝火派对中获得一次金币',
                                '4': '4. 联盟参加圆桌会议',
                                '5': '5. 给盟友赠送礼物',
                                '6': '6. 联盟商店兑换炼金石'
                            },
                            parser: parseInt
                        },
                        step: {
                            name: '完成次数',
                            type: 'text',
                            parser: parseInt
                        }
                    }
                }
            },
            client_param: {
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            }
        }
    };
    // 22
    activityTypeDefine[22] = {
        name: '联盟充值',
        fieldOption: {
            comm_param: {
                totalMoney: {
                    name: '累计金额',
                    type: 'map',
                    vertical: true,
                    groupFieldOption: {
                        _: {
                            name: '金额:money',
                            type: 'text',
                            isMapKey: true
                        },
                        masterreward: {
                            name: '盟主奖励:id,num',
                            type: 'longtext',
                            printer: utils.printItemNumList,
                            parser: utils.parseItemNumList
                        },
                        memberreward: {
                            name: '成员奖励:id,num',
                            type: 'longtext',
                            printer: utils.printItemNumList,
                            parser: utils.parseItemNumList
                        }
                    }
                },
                totalPeople: {
                    name: '累计人数',
                    type: 'map',
                    vertical: true,
                    groupFieldOption: {
                        _: {
                            name: '人数:people',
                            type: 'text',
                            isMapKey: true
                        },
                        masterreward: {
                            name: '盟主奖励:id,num',
                            type: 'longtext',
                            printer: utils.printItemNumList,
                            parser: utils.parseItemNumList
                        },
                        memberreward: {
                            name: '成员奖励:id,num',
                            type: 'longtext',
                            printer: utils.printItemNumList,
                            parser: utils.parseItemNumList
                        }
                    }
                }
            },
            client_param: {
                total_people_pos: {
                    name: '充值人数显示档位',
                    type: 'text'
                },
                total_money_pos: {
                    name: '充值金额显示档位',
                    type: 'text'
                },
                view_sort: {
                    name: '排序参数(01234，不能重复)',
                    type: 'text',
                },
                recommond: {
                    name: '焦点参数(秒)',
                    type: 'text',
                }
            }
        }
    };
    // 23
    activityTypeDefine[23] = {
        name: '开服活动',
        fieldOption: {
            comm_param: {
                activity_order: {
                    name: '加入开服活动',
                    type: 'map',
                    vertical: true,
                    groupFieldOption: {
                        _: {
                            name: '显示顺序',
                            type: 'text',
                            isMapKey: true
                        },
                        activityId: {
                            name: '活动id,活动id,活动id',
                            type: 'text',
                        },
                        extrareward: {
                            name: '额外奖励:id,num;id,num',
                            type: 'longtext',
                            printer: utils.printItemNumList,
                            parser: utils.parseItemNumList
                        },
                        extrarewardtime: {
                            name: '额外奖励的时间(秒)',
                            type: 'text',
                            parser: parseInt,
                        }
                    }
                }
            }
        }
    };

    activityTypeDefine[24] = {
        name: '通行证活动',
        fieldOption: {
            comm_param: {
                step: {
                    name: '通行证奖励',
                    type: 'map',
                    groupFieldOption: {
                        _: {
                            name: '档位',
                            type: 'text',
                            isMapKey: true,
                            parser: parseInt
                        },
                        freereward: {
                            name: '免费奖励:id,num',
                            type: 'longtext',
                            printer: utils.printItemNumList,
                            parser: utils.parseItemNumList
                        },
                        buyreward: {
                            name: '进阶奖励:id,num',
                            type: 'longtext',
                            printer: utils.printItemNumList,
                            parser: utils.parseItemNumList
                        }
                    }
                },
                productid: {
                    name: '商品id',
                    type: 'text',
                    parser: parseInt
                },
                coeffproductid: {
                    name: '折扣商品id',
                    type: 'text',
                    parser: parseInt
                },
                coeffday: {
                    name: '折扣天数',
                    type: 'text',
                    parser: parseInt
                },
                condition: {
                    name: '进度类型',
                    type: 'select',
                    options: {
                        '1': '圣山积分',
                        '2': '主线关卡',
                        '3': '迷宫关卡',
                    },
                    parser: parseInt
                }
            }
        }
    };

    // 当前活动类型
    var currentActivityType;

    // 语言翻译配置项
    var langFieldOption = {
        'info.name': {
            name: '活动名字',
            type: 'longtext'
        },
        'info.icon': {
            name: '活动图标',
            type: 'longtext'
        },
        'info.title': {
            name: '活动标题',
            type: 'longtext'
        },
        'info.brief_desc': {
            name: '简要描述',
            type: 'longtext'
        },
        'info.detail_desc': {
            name: '详细描述',
            type: 'textarea'
        },
        'info.jump_url': {
            name: '跳转链接',
            type: 'longtext'
        },
        client_param: {
            name: '客户端私有参数',
            type: 'textarea',
            parser: JSON.parse,
            subFieldOption: utils.getFieldOptionByActivity
        }
    };

    // 活动主配置项
    var mainFieldOption = {
        type: {
            name: '活动类型',
            type: 'select',
            options: (function () {
                var optionsActivityType = {};
                optionsActivityType[0] = '-- 请选择 --';
                for (var k in activityTypeDefine) {
                    optionsActivityType[k] = k + ': ' + activityTypeDefine[k].name;
                }
                return optionsActivityType;
            })(),
            checkValid: function (type) {
                if (type !== undefined && !activityTypeDefine[type]) {
                    return '未知的活动类型: ' + type;
                }
                return true;
            },
            parser: parseInt,
            onChange: utils.changeActivityType
        },
		'info.close': {
            name: '---紧急关闭活动---',
            type: 'select',
            options: options.optYesNo,
            parser: parseInt
        },
        'info.name': {
            name: '活动名字',
            type: 'longtext'
        },
        'info.icon': {
            name: '活动图标',
            type: 'longtext'
        },
        'info.title': {
            name: '活动标题',
            type: 'longtext'
        },
        'info.desc_brief': {
            name: '简要描述',
            type: 'longtext'
        },
        'info.desc_detail': {
            name: '详细描述',
            type: 'textarea'
        },
        'info.jump_url': {
            name: '跳转链接',
            type: 'longtext'
        },
        'info.jump_ui': {
            name: '跳转界面',
            type: 'text'
        },
        'info.hide_from_list': {
            name: '活动列表隐藏',
            type: 'select',
            options: options.optYesNo,
            parser: parseInt
        },
        'info.top': {
            name: '是否置顶:显示排序-优先级仅次于主推',
            type: 'select',
            options: options.optYesNo,
            parser: parseInt
        },
        'info.major': {
            name: '是否主推:开启游戏弹出活动',
            type: 'select',
            options: options.optYesNo,
            parser: parseInt
        },
        'info.hide_on_finish': {
            name: '完成后从列表隐藏',
            type: 'select',
            options: options.optYesNo,
            parser: parseInt
        },
        'info.hide_jump_button': {
            name: '隐藏跳转按钮',
            type: 'select',
            options: options.optYesNo,
            parser: parseInt
        },
        'constrain.begin_time': {
            name: '活动开始时间',
            type: 'datetime',
            separate: true,
        },
        'constrain.end_time': {
            name: '活动结束时间',
            type: 'datetime',
        },
        'constrain.show_begin_time': {
            name: '展示开始时间',
            type: 'datetime'
        },
        'constrain.show_end_time': {
            name: '展示结束时间',
            type: 'datetime'
        },
        'constrain.server_publish_day_begin': {
            name: '开服天数开始',
            type: 'number',
            parser: parseInt
        },
        'constrain.server_publish_day_end': {
            name: '开服天数结束',
            type: 'number',
            parser: parseInt
        },
        'constrain.server_publish_day_show_begin': {
            name: '开服天数展示开始',
            type: 'number',
            parser: parseInt
        },
        'constrain.server_publish_day_show_end': {
            name: '开服天数展示结束',
            type: 'number',
            parser: parseInt
        },
        'constrain.server_publish_day_last_time': {
            name: '持续时间',
            type: 'number',
            parser: parseInt
        },
        'constrain.reg_day_begin': {
            name: '创角天数开始',
            type: 'number',
            parser: parseInt
        },
        'constrain.reg_day_end': {
            name: '创角天数结束',
            type: 'number',
            parser: parseInt
        },
        'constrain.reg_day_show_begin': {
            name: '创角展示天数开始',
            type: 'number',
            parser: parseInt
        },
        'constrain.reg_day_show_end': {
            name: '创角展示天数结束',
            type: 'number',
            parser: parseInt
        },
        comm_param: {
            name: '公有参数',
            type: 'textarea',
            parser: JSON.parse,
            subFieldOption: utils.getFieldOptionByActivity,
            separate: true
        },
        server_param: {
            name: '服务端私有参数',
            type: 'textarea',
            parser: JSON.parse,
            subFieldOption: utils.getFieldOptionByActivity
        },
        client_param: {
            name: '客户端私有参数',
            type: 'textarea',
            parser: JSON.parse,
            subFieldOption: utils.getFieldOptionByActivity
        },
        'language.en': {
            name: '英文翻译',
            type: 'textarea',
            parser: JSON.parse,
            subFieldOption: langFieldOption
        }
    };

    // 导出配置
    scope.mainFieldOption = mainFieldOption;
    scope.activityTypeDefine = activityTypeDefine;
});
