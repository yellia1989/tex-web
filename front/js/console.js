"use strict";
layui.use(["okUtils", "okCountUp"], function () {
  var countUp = layui.okCountUp;
  var okUtils = layui.okUtils;
  var $ = layui.jquery;

  function realText() {
    okUtils.ajax("/api/game/real/stat", "get", null, false).done(function (response) {
      let data = response.data;
      $("#today-income").text(data.todayIncome);
      $("#total-income").text(data.totalIncome);
      $("#today-active").text(data.todayActive);
      $("#today-newadd").text(data.todayNewadd);
      $("#total-newadd").text(data.totalNewadd);
      $("#today-account").text(data.todayAccount);
      $("#total-account").text(data.totalAccount);

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
