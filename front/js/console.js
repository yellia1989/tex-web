"use strict";
layui.use(["okUtils", "table", "okCountUp", "okMock"], function () {
  var countUp = layui.okCountUp;
  var table = layui.table;
  var okUtils = layui.okUtils;
  var okMock = layui.okMock;
  var $ = layui.jquery;

  /**
   * 收入、商品、博客、用户
   */
  function statText() {
    var elem_nums = $(".stat-text");
    elem_nums.each(function (i, j) {
      var ran = parseInt(Math.random() * 99 + 1);
      !new countUp({
        target: j,
        endVal: 20 * ran
      }).start();
    });
  }

  /**
   * 所有用户
   */
  function userList() {
    table.render({
      method: "get",
      url: okMock.api.user.list,
      elem: '#userData',
      height: 340,
      page: true,
      limit: 7,
      cols: [[
        { field: "id", title: "ID", width: 180 },
        { field: "username", title: "账号", width: 100 },
        { field: "password", title: "密码", width: 80 },
        { field: "email", title: "邮箱", width: 200 },
        { field: "createTime", title: "创建时间", width: 180 },
        { field: "logins", title: "登录次数", width: 100 }
      ]],
      //             parseData: function (res) {
      //                 res.data.list.forEach(function (i, j) {
      //                     var dateTime = new Date(i.u_endtime);
      //                     i.u_endtime = dateTime.getFullYear() + "-" + dateTime.getMonth() + "-" + dateTime.getDay();
      //                 });
      //                 return {
      //                     "code": res.code,
      //                     "count": res.data.count,
      //                     "data": res.data.list
      //                 }
      //             }
    });
  }

  statText();
  userList();
});
