function checkLang(lang) {
  for (k in lang) {
    for (k2 in lang[k]) {
      let v = lang[k][k2]
      if (v == "") {
        return false;
      }
    }
  }
  return true;
}

function parseTime(d, zoneoffset, fmt) {
    zoneoffset = typeof(zoneoffset) != "undefined" ? zoneoffset : 8;
    zoneoffset = parseInt(zoneoffset);
    fmt = typeof(fmt) != "undefined" ? fmt : "YYYY-MM-DD HH:mm:ssZ";
    if (!d.match(/[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) ([01][0-9]|[2][0-3]):[0-5][0-9]:[0-5][0-9](\+|-)(0[1-9]|1[1-9]):\d{2}/)) {
        d += zoneoffset < 0 ? "-" : "+";
        let z = Math.abs(zoneoffset);
        d += z < 10 ? "0" : "";
        d += z + ":00";
    }

    let day = dayjs(d, fmt, true);
    if (!day.isValid()) throw 'invalid time format';
    return day.unix();
}

function formatTime(d, zoneoffset, fmt) {
    d = parseInt(d);
    zoneoffset = parseInt(zoneoffset);
    if (!Number.isInteger(d)) throw 'not a unixtimestamp';
    if (!Number.isInteger(zoneoffset)) throw 'not a valid zoneoffset';
    let d2 = dayjs.unix(d);
    if (!d2.isValid()) throw  'not a dayjs';
    fmt = typeof(fmt) != "undefined" ? fmt : "YYYY-MM-DD HH:mm:ssZ";
    return d2.utcOffset(zoneoffset).format(fmt);
}
