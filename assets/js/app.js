(function (d) {
  var cleaned = d.querySelector('code.language-js');
  var copy = d.querySelector('pre > .copy');
  var msg = d.querySelector('pre > .copy > span');

  if (cleaned && copy) {
    copy.addEventListener('click', function () {
      if (msg) {
        msg.style.opacity = 1;

        setTimeout(function () {
          msg.style.opacity = 0;
        }, 1250);
      }

      navigator.clipboard.writeText(cleaned.innerText);
    });
  }
})(document)
