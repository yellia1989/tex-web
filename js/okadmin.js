var tabData = [
  {
    "title": "控制台",
    "href": "pages/console.html",
    "fontFamily": "ok-icon",
    "icon": "&#xe654;"
  },
  {
    "title": "GM工具",
    "fontFamily": "ok-icon",
    "icon": "&#xe68a;",
    "spread": true,
    "children": [
      {
        "title": "GM",
        "href": "pages/gm/gm.html",
        "fontFamily": "ok-icon",
        "icon": "&#xe654;"
      },
      {
        "title": "分区管理",
        "href": "pages/gm/dir.html",
        "fontFamily": "ok-icon",
        "icon": "&#xe654;",
        "isCheck": true,
      },
    ],
  },
  {
    "title": "框架使用",
    "fontFamily": "ok-icon",
    "icon": "ok-icon-yooxi",
    "children": [
      {
        "title": "字体图标",
        "href": "pages/help/ok_font.html",
        "icon": "&#xe62e;",
        "spread": false,
      },
      {
        "title": "插件目录",
        "href": "pages/help/plug_directory.html",
        "icon": "&#xe62e;",
        "spread": false
      },
      {
        "title": "内部添加导航",
        "href": "pages/help/nav_operate.html",
        "icon": "&#xe62e;",
        "spread": false
      },
      {
        "title": "导航的图标",
        "href": "pages/help/nav_icon.html",
        "icon": "&#xe62e;",
        "spread": false
      },
      {
        "title": "导航的参数",
        "href": "pages/help/nav_parameter.html",
        "icon": "&#xe62e;",
        "spread": false
      },
      {
        "title": "okUtils",
        "href": "pages/help/okUtils.html",
        "icon": "&#xe62e;",
        "spread": false
      },
      {
        "title": "okLayer",
        "href": "pages/help/okLayer.html",
        "icon": "&#xe62e;",
        "spread": false
      },
      {
        "title": "okFly",
        "href": "pages/help/okFly.html",
        "icon": "&#xe62e;",
        "spread": false
      }
    ]
  },
  {
    "title": "会员管理",
    "href": "",
    "icon": "&#xe66f;",
    "spread": false,
    "children": [
      {
        "title": "用户列表",
        "href": "pages/member/user.html",
        "fontFamily": "layui-icon",
        "icon": "&#xe62e;",
        "spread": false
      },
      {
        "title": "角色列表",
        "href": "pages/member/role.html",
        "icon": "&#xe62e;",
        "spread": false
      },
      {
        "title": "权限列表",
        "href": "pages/member/permission.html",
        "icon": "&#xe62e;",
        "spread": false
      }
    ]
  },
  {
    "title": "图表管理",
    "href": "",
    "icon": "&#xe62c;",
    "children": [
      {
        "title": "数据图",
        "fontFamily": "ok-icon",
        "icon": "ok-icon-shuju1",
        "children": [
          {
            "title": "统计图",
            "href": "pages/chart/chart1.html",
            "icon": "&#xe62e;",
            "spread": false,
          },
          {
            "title": "折线图",
            "href": "pages/chart/chart2.html",
            "icon": "&#xe62e;",
            "spread": false
          },
          {
            "title": "馅饼图",
            "href": "pages/chart/chart4.html",
            "icon": "&#xe62e;",
            "spread": false
          },
          {
            "title": "圆形图",
            "href": "pages/chart/chart5.html",
            "icon": "&#xe62e;",
            "spread": false
          },
          {
            "title": "指数图",
            "href": "pages/chart/chart6.html",
            "icon": "&#xe62e;",
            "spread": false
          }
        ]
      }
    ]
  },
  {
    "title": "系统管理",
    "href": "",
    "fontFamily": "ok-icon",
    "icon": "&#xe68a;",
    "spread": false,
    "children": [
      {
        "title": "系统设置",
        "href": "pages/system/setup.html",
        "icon": "&#xe62e;",
        "spread": false
      },
      {
        "title": "登录页面",
        "href": "pages/login.html",
        "icon": "&#xe609;",
        "spread": false,
        "target": "_blank"
      },
      {
        "title": "403页面",
        "href": "pages/system/403.html",
        "icon": "&#xe61c;",
        "spread": false
      },
      {
        "title": "404页面",
        "href": "pages/system/404.html",
        "icon": "&#xe61c;",
        "spread": false
      },
      {
        "title": "500页面",
        "href": "pages/system/500.html",
        "icon": "&#xe61c;",
        "spread": false
      }
    ]
  },
  {
    "title": "第三方库",
    "icon": "&#xe674;",
    "children": [
      {
        "title": "jquery.okQrcode.js",
        "href": "pages/tripartite/qrcode.html"
      },
      {
        "title": "jquery.okCountup.js",
        "href": "pages/tripartite/countup.html"
      },
      {
        "title": "okCookie",
        "href": "pages/tripartite/okCookie.html"
      },
      {
        "title": "okToastr",
        "href": "pages/tripartite/okToastr.html"
      },
      {
        "title": "okMd5",
        "href": "pages/tripartite/okMd5.html"
      },
      {
        "title": "okBarcode",
        "href": "pages/tripartite/okBarcode.html"
      },
      {
        "title": "okNprogress",
        "href": "pages/tripartite/okNprogress.html"
      },
      {
        "title": "okSweetAlert2",
        "href": "pages/tripartite/okSweetAlert2.html"
      },
      {
        "title": "okAnimate",
        "href": "pages/tripartite/okAnimate.html"
      },
      {
        "title": "okLayx",
        "href": "pages/tripartite/okLayx.html"
      }
    ]
  }
];

var objOkTab = "";
layui.use(["element", "layer", "okUtils", "okTab", "okLayer", "okContextMenu", "okHoliday"], function () {
	var okUtils = layui.okUtils;
	var $ = layui.jquery;
	var layer = layui.layer;
	var okLayer = layui.okLayer;
	var okHoliday = layui.okHoliday;

	var okTab = layui.okTab({
		// 菜单请求路径
        data: tabData,
		// 允许同时选项卡的个数
		openTabNum: 30,
		// 如果返回的结果和navs.json中的数据结构一致可省略这个方法
		parseData: function (data) {
			return data;
		}
	});
	objOkTab = okTab;
	okLoading.close();/**关闭加载动画*/
	/**
	 * 左侧导航渲染完成之后的操作
	 */
	okTab.render(function () {
		/**tab栏的鼠标右键事件**/
		$("body .ok-tab").okContextMenu({
			width: 'auto',
			itemHeight: 30,
			menu: [
				{
					text: "定位所在页",
					icon: "ok-icon ok-icon-location",
					callback: function () {
						okTab.positionTab();
					}
				},
				{
					text: "关闭当前页",
					icon: "ok-icon ok-icon-roundclose",
					callback: function () {
						okTab.tabClose(1);
					}
				},
				{
					text: "关闭其他页",
					icon: "ok-icon ok-icon-roundclose",
					callback: function () {
						okTab.tabClose(2);
					}
				},
				{
					text: "关闭所有页",
					icon: "ok-icon ok-icon-roundclose",
					callback: function () {
						okTab.tabClose(3);
					}
				}
			]
		});
	});

	/**
	 * 添加新窗口
	 */
	$("body").on("click", "#navBar .layui-nav-item a, #userInfo a", function () {
		// 如果不存在子级
		if ($(this).siblings().length == 0) {
			okTab.tabAdd($(this));
		}
		// 关闭其他展开的二级标签
		$(this).parent("li").siblings().removeClass("layui-nav-itemed");
		if (!$(this).attr("lay-id")) {
			var topLevelEle = $(this).parents("li.layui-nav-item");
			var childs = $("#navBar > li > dl.layui-nav-child").not(topLevelEle.children("dl.layui-nav-child"));
			childs.removeAttr("style");
		}
	});

	/**
	 * 左侧菜单展开动画
	 */
	$("#navBar").on("click", ".layui-nav-item a", function () {
		if (!$(this).attr("lay-id")) {
			var superEle = $(this).parent();
			var ele = $(this).next('.layui-nav-child');
			var height = ele.height();
			ele.css({"display": "block"});
			// 是否是展开状态
			if (superEle.is(".layui-nav-itemed")) {
				ele.height(0);
				ele.animate({height: height + "px"}, function () {
					ele.css({height: "auto"});
				});
			} else {
				ele.animate({height: 0}, function () {
					ele.removeAttr("style");
				});
			}
		}
	});

	/**
	 * 左边菜单显隐功能
	 */
	$(".ok-menu").click(function () {
		$(".layui-layout-admin").toggleClass("ok-left-hide");
		$(this).find("i").toggleClass("ok-menu-hide");
		localStorage.setItem("isResize", false);
		setTimeout(function () {
			localStorage.setItem("isResize", true);
		}, 1200);
	});

	/**
	 * 移动端的处理事件
	 */
	$("body").on("click", ".layui-layout-admin .ok-left a[data-url], .ok-make", function () {
		if ($(".layui-layout-admin").hasClass("ok-left-hide")) {
			$(".layui-layout-admin").removeClass("ok-left-hide");
			$(".ok-menu").find('i').removeClass("ok-menu-hide");
		}
	});

	/**
	 * tab左右移动
	 */
	$("body").on("click", ".okNavMove", function () {
		var moveId = $(this).attr("data-id");
		var that = this;
		okTab.navMove(moveId, that);
	});

	/**
	 * 刷新当前tab页
	 */
	$("body").on("click", ".ok-refresh", function () {
		okTab.refresh(this, function (okTab) {
			//刷新之后所处理的事件
		});
	});

	/**
	 * 关闭tab页
	 */
	$("body").on("click", "#tabAction a", function () {
		var num = $(this).attr("data-num");
		okTab.tabClose(num);
	});

	/**
	 * 退出操作
	 */
	$("#logout").click(function () {
		okLayer.confirm("确定要退出吗？", function (index) {
			okTab.removeTabStorage(function (res) {
				okTab.removeTabStorage();
				window.location = "pages/login.html";
			});
		});
	});
});
