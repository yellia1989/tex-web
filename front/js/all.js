"use strict";
layui.use(["okUtils", "okCountUp"], function () {
  var countUp = layui.okCountUp;
  var okUtils = layui.okUtils;
  var $ = layui.jquery;

  function realText() {
    okUtils.ajax("/api/game/real/stat", "get", null, false).done(function (response) {
      let data = response.data;
      $("#accountToday").text(data.accountToday);
      $("#accountTotal").text(data.accountTotal);
      $("#activeToday").text(data.activeToday);
      $("#newaddToday").text(data.newaddToday);
      $("#newaddTotal").text(data.newaddTotal);
      $("#rgeNewRoleNumToday").text(data.rgeNewRoleNumToday);
      $("#rgeRoleNumToday").text(data.rgeRoleNumToday);
      $("#rgeRoleNumTotal").text(data.rgeRoleNumTotal);
      $("#rgeToday").text(data.rgeToday/100);
      $("#rgeTotal").text(data.rgeTotal/100);

      let elem_nums = $(".stat-text");
      elem_nums.each(function (i, j) {
        !new countUp({
          target: j,
        }).start();
      });

    }).fail(function (error) {
        console.log(error);
    });
  }

  realText();
});
