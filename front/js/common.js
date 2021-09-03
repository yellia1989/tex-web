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
